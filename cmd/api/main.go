package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"trax/pkg/shutdown"

	ggrpc "github.com/elojah/trax/pkg/grpc"
	"github.com/elojah/trax/pkg/grpcweb"
	ghttp "github.com/elojah/trax/pkg/http"
	glog "github.com/elojah/trax/pkg/log"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/redis"
	"github.com/rs/zerolog/log"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

const (
	// Time allocated for init phase (connections + setup).
	initTO = 30 * time.Second
)

var version string

// run services.
func run(prog string, filename string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, initTO)
	// no need for defer cancel
	_ = cancel

	var cs shutdown.Closers

	logs := glog.Service{}
	if err := logs.Dial(ctx, glog.Config{}); err != nil {
		fmt.Println("failed to dial logger")

		return
	}

	cs = append(cs, &logs)

	log.Logger = logs.With().Str("version", version).Str("exe", prog).Logger()

	// read config
	cfg := config{}
	if err := cfg.Populate(ctx, filename); err != nil {
		log.Error().Err(err).Msg("failed to read config")

		return
	}

	// init Postgres storage
	postgress := postgres.Service{}
	if err := postgress.Dial(ctx, cfg.Postgres); err != nil {
		log.Error().Err(err).Msg("failed to dial postgres")

		return
	}

	cs = append(cs, &postgress)

	// init redis storage
	rediss := redis.Service{}
	if err := rediss.Dial(ctx, cfg.Redis); err != nil {
		log.Error().Err(err).Msg("failed to dial redis")

		return
	}

	cs = append(cs, &rediss)

	cookieCache := &cookieredis.Cache{Service: rediss}
	cookieApp := &cookieapp.A{
		CacheKeys: cookieCache,
	}

	// init http api server
	https := ghttp.Service{}

	if err := https.Dial(ctx, cfg.HTTP); err != nil {
		log.Error().Err(err).Msg("failed to dial http")

		return
	}

	cs = append(cs, &https)

	userStore := &userpostgres.Store{Service: postgress}
	userCache := &userredis.Cache{Service: rediss}
	userApp := userapp.App{
		Store:        userStore,
		StoreSession: userStore,
		CacheSession: userCache,
		Cookie:       cookieApp,
	}

	if err := userApp.Dial(ctx, cfg.Session); err != nil {
		log.Error().Err(err).Msg("failed to dial user application")

		return
	}

	// init handler
	h := handler{
		user: userApp,
	}

	// init grpc ONLY server
	grpcapi := ggrpc.Service{}
	grpcapi.Register = func() {
		reflection.Register(grpcapi.Server)
		apigrpc.RegisterAPIServer(grpcapi.Server, &h)
	}

	if err := grpcapi.Dial(ctx, cfg.GRPC); err != nil {
		log.Error().Err(err).Msg("failed to dial grpc")

		return
	}

	cs = append(cs, &grpcapi)

	go func() {
		if err := grpcapi.Serve(ctx); err != nil {
			log.Error().Err(err).Msg("failed to serve grpc")
		}
	}()

	// init grpc api server
	grpcwapi := grpcweb.Service{}
	grpcwapi.Register = func() {
		apigrpc.RegisterAPIServer(grpcwapi.Server, &h)
		// reflection.Register(grpcwapi.Server)
		https.Handler = grpcwapi.WrappedGrpcServer
	}

	if err := grpcwapi.Dial(ctx, cfg.GRPCWeb); err != nil {
		log.Error().Err(err).Msg("failed to dial grpcweb")

		return
	}

	cs = append(cs, &grpcwapi)

	// serve grpcweb api
	go func() {
		if err := https.Server.ListenAndServeTLS("", ""); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info().Msg("http server closed")
			} else {
				log.Error().Err(err).Msg("failed to serve http")
			}
		}
	}()

	log.Info().Msg("api up")

	cs.WaitSignal(ctx, initTO)
}

func main() {
	args := os.Args
	if len(args) != 2 { //nolint: gomnd
		fmt.Printf("Usage: ./%s configfile\n", args[0])

		return
	}

	run(args[0], args[1])
}

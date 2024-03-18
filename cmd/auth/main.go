package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	authgrpc "github.com/elojah/trax/cmd/auth/grpc"
	userapp "github.com/elojah/trax/internal/user/app"
	userpostgres "github.com/elojah/trax/internal/user/postgres"
	cookieapp "github.com/elojah/trax/pkg/cookie/app"
	cookieredis "github.com/elojah/trax/pkg/cookie/redis"
	googleapp "github.com/elojah/trax/pkg/google/app"
	ggrpc "github.com/elojah/trax/pkg/grpc"
	glog "github.com/elojah/trax/pkg/log"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/redis"
	"github.com/elojah/trax/pkg/shutdown"
	twitchapp "github.com/elojah/trax/pkg/twitch/app"
	twitchhttp "github.com/elojah/trax/pkg/twitch/http"
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

	// init ghttp web server
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

	// init domain
	twitchApp := twitchapp.App{
		Client: &twitchhttp.Client{
			Client: http.DefaultClient,
		},
	}
	if err := twitchApp.Dial(ctx, cfg.Twitch); err != nil {
		log.Error().Err(err).Msg("failed to dial twitch app")

		return
	}

	googleApp := googleapp.App{}
	if err := googleApp.Dial(ctx, cfg.Google); err != nil {
		log.Error().Err(err).Msg("failed to dial google app")

		return
	}

	cs = append(cs, &googleApp)

	cookieCache := &cookieredis.Cache{Service: rediss}
	cookieApp := &cookieapp.A{
		CacheKeys: cookieCache,
	}

	userStore := &userpostgres.Store{}
	// userStore := &userpostgres.StoreProfile{}
	userApp := userapp.App{
		Transactioner: postgress,
		Store:         userStore,
		Cookie:        cookieApp,
	}
	if err := userApp.Dial(ctx); err != nil {
		log.Error().Err(err).Msg("failed to dial user application")

		return
	}

	// init handler
	h := handler{
		twitch: twitchApp,
		google: googleApp,

		user: userApp,
	}

	// init grpc api server
	grpcauth := ggrpc.Service{}
	grpcauth.Register = func() {
		reflection.Register(grpcauth.Server)
		authgrpc.RegisterAuthServer(grpcauth.Server, &h)
	}

	if err := grpcauth.Dial(ctx, cfg.GRPC); err != nil {
		log.Error().Err(err).Msg("failed to dial grpc")

		return
	}

	cs = append(cs, &grpcauth)

	go func() {
		if err := grpcauth.Serve(ctx); err != nil {
			log.Error().Err(err).Msg("failed to serve grpc")
		}
	}()

	log.Info().Msg("auth up")

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

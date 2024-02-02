package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	admingrpc "github.com/elojah/trax/internal/admin/grpc"
	"github.com/elojah/trax/pkg/cookie"
	cookieapp "github.com/elojah/trax/pkg/cookie/app"
	cookieredis "github.com/elojah/trax/pkg/cookie/redis"
	ggrpc "github.com/elojah/trax/pkg/grpc"
	glog "github.com/elojah/trax/pkg/log"
	migrateapp "github.com/elojah/trax/pkg/migrate/app"
	"github.com/elojah/trax/pkg/postgres"
	"github.com/elojah/trax/pkg/redis"
	"github.com/elojah/trax/pkg/shutdown"
	"github.com/rs/zerolog/log"
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

	var cs []shutdown.Closer

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

	cookieCache := &cookieredis.Cache{Service: rediss}
	cookieApp := &cookieapp.A{
		CacheKeys: cookieCache,
	}

	// setup initial cookie keys
	if err := cookieApp.Setup(ctx, cookie.NKeys); err != nil {
		log.Error().Err(err).Msg("failed to setup cookie keys")

		return
	}

	migrateApp := &migrateapp.App{
		Service: &postgress,
	}

	h := handler{
		cookie:  cookieApp,
		migrate: migrateApp,
	}

	// init grpc api server
	grpcadmin := ggrpc.Service{}
	grpcadmin.Register = func() {
		reflection.Register(grpcadmin.Server)
		admingrpc.RegisterAdminServer(grpcadmin.Server, &h)
	}

	if err := grpcadmin.Dial(ctx, cfg.GRPC); err != nil {
		log.Error().Err(err).Msg("failed to dial grpc")

		return
	}

	cs = append(cs, &grpcadmin)

	go func() {
		if err := grpcadmin.Serve(ctx); err != nil {
			log.Error().Err(err).Msg("failed to serve grpc")
		}
	}()

	log.Info().Msg("admin up")

	// listen for signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for sig := range c {
		switch sig {
		case syscall.SIGHUP:
			fallthrough
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGTERM:
			if err := closers(cs).Close(ctx); err != nil {
				fmt.Printf("error closing service: %s\n", err.Error())

				return
			}

			cancel()

			fmt.Println("successfully closed service")

			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 { //nolint: gomnd
		fmt.Printf("Usage: ./%s configfile\n", args[0])

		return
	}

	run(args[0], args[1])
}

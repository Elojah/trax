package main

import (
	"context"
	"fmt"
	"os"
	"time"

	admingrpc "github.com/elojah/trax/cmd/admin/grpc"
	"github.com/elojah/trax/pkg/cookie"
	cookieredis "github.com/elojah/trax/pkg/cookie/redis"
	cookieagg "github.com/elojah/trax/pkg/cookie/service"
	ggrpc "github.com/elojah/trax/pkg/grpc"
	glog "github.com/elojah/trax/pkg/log"
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

	// init redis storage
	rediss := redis.Service{}
	if err := rediss.Dial(ctx, cfg.Redis); err != nil {
		log.Error().Err(err).Msg("failed to dial redis")

		return
	}

	cs = append(cs, &rediss)

	cookieCache := &cookieredis.Cache{Service: rediss}
	cookieService := &cookieagg.S{
		CacheKeys: cookieCache,
	}

	// setup initial cookie keys
	if err := cookieService.Setup(ctx, cookie.NKeys); err != nil {
		log.Error().Err(err).Msg("failed to setup cookie keys")

		return
	}

	h := handler{
		cookie: cookieService,
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

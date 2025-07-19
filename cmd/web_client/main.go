package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	authgrpc "github.com/elojah/trax/cmd/auth/grpc"
	cookieredis "github.com/elojah/trax/pkg/cookie/redis"
	cookieagg "github.com/elojah/trax/pkg/cookie/service"
	googleapp "github.com/elojah/trax/pkg/google/service"
	tgrpc "github.com/elojah/trax/pkg/grpc"
	thttp "github.com/elojah/trax/pkg/http"
	tlog "github.com/elojah/trax/pkg/log"
	"github.com/elojah/trax/pkg/redis"
	"github.com/elojah/trax/pkg/shutdown"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	_ "google.golang.org/grpc/encoding/gzip"
)

const (
	// Time allocated for init phase (connections + setup).
	initTO = 30 * time.Second
)

var version string

type closer interface {
	Close(context.Context) error
}

type closers []closer

func (cs closers) Close(ctx context.Context) error {
	var result *multierror.Error

	for _, c := range cs {
		if c != nil {
			result = multierror.Append(result, c.Close(ctx))
		}
	}

	return result.ErrorOrNil()
}

// run services.
func run(prog string, filename string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, initTO)
	// no need for defer cancel
	_ = cancel

	var cs shutdown.Closers

	logs := tlog.Service{}
	if err := logs.Dial(ctx, tlog.Config{}); err != nil {
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

	// init http web server
	https := thttp.Service{}

	if err := https.Dial(ctx, cfg.HTTP); err != nil {
		log.Error().Err(err).Msg("failed to dial http")

		return
	}

	cs = append(cs, &https)

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

	authclient := tgrpc.Client{}
	if err := authclient.Dial(ctx, cfg.AuthClient); err != nil {
		log.Error().Err(err).Msg("failed to dial auth")

		return
	}

	cs = append(cs, &authclient)

	googleService := googleapp.S{}
	if err := googleService.Dial(ctx, cfg.Google); err != nil {
		log.Error().Err(err).Msg("failed to dial google service")

		return
	}

	h := handler{
		cookie:     cookieService,
		google:     googleService,
		AuthClient: authgrpc.NewAuthClient(authclient.ClientConn),
	}

	// auth
	https.Router.Path("/signup").Methods(http.MethodPost).HandlerFunc(h.signup)
	https.Router.Path("/signin").Methods(http.MethodPost).HandlerFunc(h.signin)
	https.Router.Path("/signout").Methods(http.MethodGet).HandlerFunc(h.signout)
	https.Router.Path("/signin_google").Methods(http.MethodPost).HandlerFunc(h.signinGoogle)
	https.Router.Path("/redirect_google").Methods(http.MethodPost).HandlerFunc(h.redirectGoogle)
	https.Router.Path("/refresh_token").Methods(http.MethodPost).HandlerFunc(h.refreshToken)

	// serve static dir
	https.Router.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir(path.Join(cfg.Web.Static, "assets")))))
	favicon := path.Join(cfg.Web.Static, "favicon.ico")
	https.Router.Path("/favicon.ico").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, favicon)
	}))
	index := path.Join(cfg.Web.Static, "index.html")
	https.Router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, index)
	}))

	https.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	// serve http web
	go func() {
		if err := https.Server.ListenAndServeTLS(cfg.HTTP.Cert, cfg.HTTP.Key); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info().Msg("http server closed")
			} else {
				log.Error().Err(err).Msg("failed to serve http")
			}
		}
	}()
	log.Info().Msg("web client up")

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

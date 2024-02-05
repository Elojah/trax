package shutdown

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Closer interface {
	Close(context.Context) error
}

type Closers []Closer

func (cs Closers) Close(ctx context.Context) error {
	var err error

	for _, c := range cs {
		if c != nil {
			err = errors.Join(err, c.Close(ctx))
		}
	}

	return err
}

func (cs Closers) WaitSignal(ctx context.Context, to time.Duration) {
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
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, to)

			defer cancel()

			if err := cs.Close(ctx); err != nil {
				fmt.Printf("error closing service: %s\n", err.Error())

				return
			}

			fmt.Println("successfully closed service")

			return
		}
	}
}

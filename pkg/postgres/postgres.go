package postgres

import (
	"context"
	"fmt"

	"github.com/elojah/trax/pkg/errors"
	"github.com/elojah/trax/pkg/transaction"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var _ transaction.Transactioner = (*Service)(nil)

// Service wraps a postgres client.
type Service struct {
	*pgxpool.Pool
}

// Dial connects postgres server.
func (s *Service) Dial(ctx context.Context, cfg Config) error {

	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	var err error

	s.Pool, err = pgxpool.New(ctx, info)

	if err != nil {
		return err
	}

	return s.Pool.Ping(ctx)
}

func (s *Service) Close(ctx context.Context) error {
	if s.Pool != nil {
		s.Pool.Close()
	}

	return nil
}

func (s Service) Tx(ctx context.Context, access transaction.AccessMode, f func(context.Context) (transaction.Operation, error)) (rerr error) {
	tx, err := s.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable, // safest
		AccessMode: func(access transaction.AccessMode) pgx.TxAccessMode {
			switch access {
			case transaction.Read:
				return pgx.ReadOnly
			case transaction.Write:
				return pgx.ReadWrite
			default:
				return pgx.ReadOnly
			}
		}(access),
		DeferrableMode: pgx.NotDeferrable,
	})
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, transaction.Key{}, tx)

	// panic safeguard for commit operation
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(ctx); err != nil {
				rerr = err
			}

			panic(p)
		}
	}()

	op, err := f(ctx)

	switch op {
	case transaction.Rollback:
		if terr := tx.Rollback(ctx); terr != nil {
			return terr
		}

		return err
	case transaction.Commit:
		if cerr := tx.Commit(ctx); cerr != nil {
			if rerr := tx.Rollback(ctx); rerr != nil {
				return rerr
			}

			return cerr
		}
	}

	return err
}

func Tx(ctx context.Context) (pgx.Tx, error) {
	tx, ok := ctx.Value(transaction.Key{}).(*pgxpool.Tx)
	if !ok {
		return nil, errors.ErrMissingTransaction{}
	}
	return tx, nil
}

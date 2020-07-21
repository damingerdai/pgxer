package transaction

import (
	"context"

	"github.com/damingerdai/pgxer/pkg/connection"
	"github.com/jackc/pgx/v4"
)

type Transaction interface {
	connection.Connection
	DoTransaction(ctx context.Context, f func(ctx context.Context, tx *pgx.Tx) error)
}

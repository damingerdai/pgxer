package db

import (
	"context"
)

type Transaction interface {
	Connection
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

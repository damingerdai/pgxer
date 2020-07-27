package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DataSource struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *DataSource {
	return &DataSource {
		pool: pool,
	}
}

func (ds *DataSource) GetConnection() Connection {
	return ds.pool
}

func (ds *DataSource) GetTransaction(ctx context.Context) (Transaction, error) {
	return ds.pool.Begin(ctx)
}

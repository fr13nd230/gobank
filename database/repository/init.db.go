package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDb will initialize new DB pool and acquire new conn from that pool
// in order to create a new Queries for intercating with SQLc boiler plate code.
func NewDb(path string) (*Queries, error) {	
	pool, err := pgxpool.New(context.Background(), path)
	if err != nil {
		return nil, err
	}
	
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		pool.Close()
		return nil, err
	}
	
	q := New(conn)
	
	return q, nil
}
package repository

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/fr13nd230/gobank/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var queries *Queries

func TestMain(m *testing.M) {

	err := config.LoadConfig("../../.env")
	if err != nil {
		slog.Error("[MainTest]: Could not load env file.", "error", err)
	}
	
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_PATH"))
	if err != nil {
		slog.Error("[MainTest]: Could not create a new pool.", "error", err)
		os.Exit(1)
	}
	defer pool.Close()
	
	db, err := pool.Acquire(context.Background())
	if err != nil {
		slog.Error("[MainTest]: Could not aquire new connection from pool.", "error", err)
		os.Exit(1)
	}
	defer db.Release()
	
	queries = New(db)
	code := m.Run()
	
	os.Exit(code)
}

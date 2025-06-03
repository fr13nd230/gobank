package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/fr13nd230/gobank/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var queries *Queries

func TestMain(m *testing.M) {

	err := config.LoadConfig("../../.env")
	if err != nil {
		log.Fatalf("Error in main test while loading [.env]: %v", err)
	}
	
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		os.Exit(1)
	}
	defer pool.Close()
	
	db, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Error acquiring connection from pool: %v", err)
		os.Exit(1)
	}
	defer db.Release()
	
	queries = New(db)
	code := m.Run()
	
	os.Exit(code)
}

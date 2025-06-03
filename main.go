package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	cfg "github.com/fr13nd230/gobank/config"
)

// Represents the main entry file for now
// until we complete setting up all necessaries.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()
	defer func(){
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()
	
	err := cfg.LoadConfig()
	if err != nil {
		log.Fatalf("Error while loading .env file: %v", err)
	}
	
	<-ctx.Done() // Leave this in here for now
	os.Exit(0)
}
package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	cfg "github.com/fr13nd230/gobank/config"
	"github.com/fr13nd230/gobank/database/repository"
	"github.com/fr13nd230/gobank/src/domains/accounts"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	fRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

// main represent the entry point and mouting point
// of all configurations in the project.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	defer func(){
		if r := recover(); r != nil {
			log.Print("[Server]: Recovered from panic.", r)
		}
	}()
	
	err := cfg.LoadConfig()
	if err != nil {
		log.Printf("[Config]: Error while loading .env file: %v", err)
	}
	port := cfg.GetVar("PORT")
	dbPath := cfg.GetVar("POSTGRES_PATH")
	
	q, err := repository.NewDb(dbPath)
	if err != nil {
		log.Printf("[Database]: Error while creating new queries: %v", err)
	}
		
	app := fiber.New(fiber.Config{
		Prefork: false,
		CaseSensitive: true,
		BodyLimit: 25 * 1024 * 1024,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		// ErrorHandler: func(c *fiber.Ctx) error {}
	})
	
	app.Use(fRecover.New())
	app.Use(limiter.New(limiter.Config{
		Max: 200,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(helmet.New())
	app.Use(idempotency.New(idempotency.Config{
		Lifetime: 40 * time.Minute,
		Lock: nil,
		Storage: nil,
	}))
		
	v1 := app.Group("/v1")
	registerRoutesV1(v1, q)
	
	log.Print("[localhost]: Server is starting...", "port", port)
	
	serverErrs := make(chan error, 1)
	go func() {
		serverErrs <- app.Listen(port)
	}()
	
	select {
	case err := <-serverErrs:
		log.Print("[localhost]: Could not start the server.", "error", err)
	case <-ctx.Done():
		log.Print("[Server]: Shutdown signal received, shutting down gracefully...")
		
		shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		
		if err := app.ShutdownWithContext(shutCtx); err != nil {
			log.Printf("[Server]: Could not shutdown properly: %v", err)
		}
		
		log.Print("[Server]: Server shutdown complete")
	}
}

func registerRoutesV1(r fiber.Router, q *repository.Queries) {
	// Mount Accounts Routes
	accounts.RegisterRoutes(r, q)
}
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "github.com/fr13nd230/gobank/config"
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

// Represents the main entry file for now
// until we complete setting up all necessaries.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()
	defer func(){
		if r := recover(); r != nil {
			slog.Info("Recovered from panic: %s", r)
		}
	}()
	
	err := cfg.LoadConfig()
	if err != nil {
		slog.Error("Error while loading .env file: %v", err)
	}
	
	port := cfg.GetVar("PORT")
	
	app := fiber.New(fiber.Config{
		Prefork: false,
		CaseSensitive: true,
		BodyLimit: 25 * 1024 * 1024,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		// TODO: To be handled later 
		// ErrorHandler:,
	})
	
	app.Use(fRecover.New())
	app.Use(limiter.New(limiter.Config{
		Max: 60000,
		Expiration: 5 * time.Minute,
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
	registerRoutesV1(v1)
	
	if err := app.Listen(port); err != nil {
		slog.Error("Server error while listening on %s, with error: %v", port, err)
	} else {
		slog.Info("[http://localhost:%s]: Server is running.", port)
	}
	
	<-ctx.Done()
	os.Exit(0)
}

func registerRoutesV1(r fiber.Router) {
	// Mount Accounts Routes
	accounts.RegisterRoutes(r)
}
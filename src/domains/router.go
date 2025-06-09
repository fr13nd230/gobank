package domains

import (
	"github.com/fr13nd230/gobank/database/repository"
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	RegisterRoutes(r fiber.Router, q *repository.Queries)
}
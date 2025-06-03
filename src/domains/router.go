package domains

import "github.com/gofiber/fiber/v2"

type Router interface {
	RegisterRoutes(r fiber.Router)
}
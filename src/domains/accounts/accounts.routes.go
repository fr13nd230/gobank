package accounts

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router) {
	
	r.Post("/account", CreateAccountHandler)
}
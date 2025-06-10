package accounts

import (
	"github.com/fr13nd230/gobank/database/repository"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, q *repository.Queries) {
    r.Post("/account", CreateAccountHandler(q))
    r.Get("/accounts", ListAccountsHandler(q))
    r.Get("/account/:id", FindAccountByIdHandler(q))
}
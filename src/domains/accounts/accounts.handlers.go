package accounts

import (
	"context"
	"log/slog"

	"github.com/fr13nd230/gobank/database/repository"
	"github.com/fr13nd230/gobank/src/types"
	"github.com/gofiber/fiber/v2"
)

type CreateAccountBody struct {
	Owner    string `json:"owner" form:"owner"`
	Currency string `json:"currency" form:"currency"`
}

func CreateAccountHandler(q *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(CreateAccountBody)
		if err := c.BodyParser(body); err != nil {
			r := types.NewBase(
				false,
				fiber.StatusBadRequest,
				"Body is invalid JSON or encoded forms.",
			)
		
			return c.Status(fiber.StatusBadRequest).JSON(r)
		} 
		
		arg := repository.NewAccountParams{
			Owner: body.Owner,
			Currency: body.Currency,
		}
		acc, err := CreateAccountProvider(context.Background(), arg, q)
		if err != nil {
			slog.Error("[AccountsHandlers]: Could not call the provier", "error", err)
			r := types.NewBase(
				false,
				fiber.StatusBadRequest,
				"Something bad happened, try again later.",
			)
			return c.Status(fiber.StatusBadRequest).JSON(r)
		} 
		
		r := types.NewResponse[repository.Account](
			false, 
			fiber.StatusOK, 
			"Account created successfully", 
			acc,
		)
		
		return c.Status(fiber.StatusOK).JSON(r)
	}
}

func ListAccountsHandler(q *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func FindAccountByIdHandler(q *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func UpdateAccountByIdHandler(q *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func DeleteAccountByIdHandler(q *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
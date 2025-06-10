package accounts

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/emicklei/pgtalk/convert"
	"github.com/fr13nd230/gobank/database/repository"
	"github.com/fr13nd230/gobank/src/types"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
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
		
		return c.Status(fiber.StatusCreated).JSON(r)
	}
}

func ListAccountsHandler(q *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit := c.QueryInt("limit", 10)
		offset := c.QueryInt("offset", 0)
	    arg := repository.ListAccountsParams{
			Limit: int32(limit),
			Offset: int32(offset),
		}
		
		accs, err := ListAccountsProvider(context.Background(), arg, q)
		if err != nil {
		    slog.Error("[AccountsHandlers]: List accounts provider failed", "error", err)
			r := types.NewBase(
			    false, 
				fiber.StatusBadRequest, 
				"Something bad happened, try again later",
			)
			return c.Status(fiber.StatusBadRequest).JSON(r)
		}
		
		r := types.NewManyResponse[repository.Account](
		    true, 
			fiber.StatusOK, 
			"Accounts has been fetched succesfully.",
			accs,
		)
		return c.Status(fiber.StatusOK).JSON(r)
	}
}

func FindAccountByIdHandler(q *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
	    id := convert.StringToUUID(c.Params("id", ""))
		acc, err := FindAccountByIdProvider(context.Background(), id, q)
		if err == pgx.ErrNoRows {
            r := types.NewBase(
                false,
                fiber.StatusNotFound,
                fmt.Sprintf("Account with id: %s, is not found. Possibly deleted.", id),
            )
            return c.Status(fiber.StatusOK).JSON(r)			
		}
		if err != nil {
		    slog.Error("[AccountsHandlers]: Couldn not fetch the account with ", "id", id, "error", err)
		    r := types.NewBase(
				false,
				fiber.StatusBadRequest,
				"Something bad happened, try again later.",
			)
			return c.Status(fiber.StatusOK).JSON(r)			
		}
		
		r := types.NewResponse[repository.Account](
		    true, 
			fiber.StatusOK, 
			"Account has been fetched succesfully.",
			acc,
		)
		
		return c.Status(fiber.StatusOK).JSON(r)
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
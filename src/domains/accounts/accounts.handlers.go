package accounts

import (
	"github.com/fr13nd230/gobank/src/types"
	"github.com/gofiber/fiber/v2"
)

type CreateAccountBody struct {
	Owner    string `json:"owner" form:"owner"`
	Currency string `json:"currency" form:"currency"`
}

func CreateAccountHandler(c *fiber.Ctx) error {//(rp.Account, error) {
	
	// TODO: Get Body
	user := new(CreateAccountBody)
	
	if err := c.BodyParser(user); err != nil {
		return err
	}
	
	// TODO: Validate Body
	// Create New Arg
	// TODO: Call Database
	// Return Account, nil
	// 
	
	r := types.NewResponse[CreateAccountBody](
		false, 
		fiber.StatusOK, 
		"Account created successfully", 
		*user,
	)
	
	c.SendStatus(fiber.StatusOK)
	return c.JSON(r)
}

func ListAccountsHandler(c *fiber.Ctx) string {
	return ""
}

func FindAccountByIdHandler(c *fiber.Ctx) string {
	return ""
}

func UpdateAccountByIdHandler(c *fiber.Ctx) string {
	return ""
}

func DeleteAccountByIdHandler(c *fiber.Ctx) string {
	return ""
}
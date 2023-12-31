package handlers

import (
	"github.com/godofprodev/simple-pass/internal/auth"
	"github.com/godofprodev/simple-pass/internal/models"
	"github.com/godofprodev/simple-pass/internal/response"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (h *Handlers) HandleLogin(c *fiber.Ctx) error {
	params := new(models.LoginUserParams)

	if err := c.BodyParser(&params); err != nil {
		return response.ErrParsingParams()
	}

	if params.Username == "" {
		return response.ErrRequired("user")
	}
	if params.Password == "" {
		return response.ErrRequired("password")
	}

	user, err := h.store.GetUser(params.Username)
	if err != nil {
		return response.ErrNotFound(params.Username)
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(params.Password)); err != nil {
		return response.ErrIncorrectPassword()
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		return response.ErrGeneratingToken()
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return response.SuccessMessage("successfully logged in as " + user.Username)
}

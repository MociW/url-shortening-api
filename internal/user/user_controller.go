package user

import (
	"fmt"
	"time"
	"url-shortening-api/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	UserService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) Register(c *fiber.Ctx) error {
	request := new(RegisterUserRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := u.UserService.Register(c.UserContext(), request)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(response)
}

func (u *UserController) Login(c *fiber.Ctx) error {
	request := new(LoginUserRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := u.UserService.Login(c.UserContext(), request)
	if err != nil {
		return err
	}

	claims := jwt.MapClaims{
		"id":  response.UUID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Salt))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (u *UserController) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Logout Successfully",
	})
}

func (u *UserController) Update(c *fiber.Ctx) error {
	request := new(UpdateUserRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Salt), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	if (*claims)["id"].(string) != request.UUID {
		fmt.Println((*claims)["id"].(string))
		return fiber.ErrBadRequest
	}

	response, err := u.UserService.Update(c.UserContext(), request)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (u *UserController) Delete(c *fiber.Ctx) error {
	request := new(DeleteUserRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Salt), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse claims",
		})
	}

	if (*claims)["email"].(string) != request.Email {
		return fiber.ErrBadRequest
	}

	if err := u.UserService.Delete(c.UserContext(), request); err != nil {
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Delete User Successfully",
	})
}

package link

import (
	"fmt"
	"url-shortening-api/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LinkController struct {
	*LinkService
}

func NewLinkController(service *LinkService) *LinkController {
	return &LinkController{service}
}

func (l *LinkController) RedirectLink(c *fiber.Ctx) error {
	request := new(GetLinkRequest)
	id := c.Params("id")
	fmt.Println(id)

	request.Link = fmt.Sprintf("http://localhost:3000/short/%s", id)

	response, err := l.LinkService.Get(c.Context(), request)
	if err != nil {
		return err
	}
	return c.Redirect(response.Link, fiber.StatusMovedPermanently)
}

func (l *LinkController) Create(c *fiber.Ctx) error {
	request := new(CreateLinkRequest)
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
			"error": "Unauthorized: Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Token validation failed",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "InternalServerError: Failed to parse token claims",
		})
	}

	jwtId, ok := (*claims)["id"].(string)
	if !ok || jwtId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest: User ID not found in token",
		})
	}

	if jwtId != request.UserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: You do not have access to this data",
		})
	}

	response, err := l.LinkService.Create(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

func (l *LinkController) List(c *fiber.Ctx) error {
	request := new(ListLinkRequest)
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
			"error": "Unauthorized: Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Token validation failed",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "InternalServerError: Failed to parse token claims",
		})
	}

	jwtId, ok := (*claims)["id"].(string)
	if !ok || jwtId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest: User ID not found in token",
		})
	}

	if jwtId != request.UserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: You do not have access to this data",
		})
	}

	response, err := l.LinkService.List(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(response)
}

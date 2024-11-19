package link

import (
	"context"
	"fmt"
	"log"
	"url-shortening-api/internal/helper"

	"github.com/gofiber/fiber/v2"
)

type LinkService struct {
	LinkRepository *LinkRepository
}

func NewLinkService(linkRepository *LinkRepository) *LinkService {
	return &LinkService{LinkRepository: linkRepository}
}

func (l *LinkService) Create(ctx context.Context, request *CreateLinkRequest) (*LinkResponse, error) {

	key := helper.RandomString()
	shortLink := fmt.Sprintf("http://localhost:3000/short/%s", key)

	link := &Link{
		UserId:    request.UserId,
		Link:      request.Link,
		ShortLink: shortLink,
	}

	if err := l.LinkRepository.Create(ctx, request.UserId, link); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return ConvertLinkResponse(link), nil
}

func (l *LinkService) Get(ctx context.Context, request *GetLinkRequest) (*LinkResponse, error) {

	link := new(Link)
	err := l.LinkRepository.FindByLink(ctx, link, request.Link)
	if err != nil {
		return nil, fiber.ErrNotFound
	}

	responses := ConvertLinkResponse(link)

	return responses, nil
}

func (l *LinkService) List(ctx context.Context, request *ListLinkRequest) ([]LinkResponse, error) {

	links, err := l.LinkRepository.ListByUUID(ctx, request.UserId)
	log.Println(links)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]LinkResponse, len(links))
	for i, todo := range links {
		responses[i] = *ConvertLinkResponse(&todo)
	}

	return responses, nil
}

func (l *LinkService) Delete(ctx context.Context, request *DeleteLinkRequest) error {

	err := l.LinkRepository.Delete(ctx, request.UserId, request.Link)

	return err
}

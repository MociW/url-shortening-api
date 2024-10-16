package link

import "time"

type LinkResponse struct {
	ID        uint      `json:"id"`
	UserId    string    `json:"user_id"`
	Link      string    `json:"link"`
	ShortLink string    `json:"short_link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateLinkRequest struct {
	UserId string `json:"user_id"`
	Link   string `json:"link"`
}

type DeleteLinkRequest struct {
	UserId string `json:"user_id"`
	Link   string `json:"link"`
}

type GetLinkRequest struct {
	Link string `json:"-"`
}
type ListLinkRequest struct {
	UserId string `json:"-"`
}

func ConvertLinkResponse(link *Link) *LinkResponse {
	response := &LinkResponse{
		ID:        link.ID,
		UserId:    link.UserId,
		Link:      link.Link,
		ShortLink: link.ShortLink,
		CreatedAt: link.CreatedAt,
		UpdatedAt: link.UpdatedAt,
	}
	return response
}

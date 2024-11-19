package user

import (
	"time"
	"url-shortening-api/internal/link"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserProfileResponse struct {
	ID       uint        `json:"id"`
	UUID     string      `json:"uuid"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Links    []link.Link `json:"links"`
}

type GetUserRequest struct {
	UUID string `json:"-"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ConvertUserResponse(user *User) *UserResponse {
	response := &UserResponse{
		ID:        user.ID,
		UUID:      user.UUID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return response
}

func ConvertUserProfileResponse(user *User) *UserProfileResponse {
	return &UserProfileResponse{
		ID:       user.ID,
		UUID:     user.UUID,
		Username: user.Username,
		Email:    user.Email,
		Links:    user.Links,
	}
}

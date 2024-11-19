package user

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) Register(ctx context.Context, request *RegisterUserRequest) (*UserResponse, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}
	token := uuid.New().String()

	user := &User{
		Username: request.Username,
		UUID:     token,
		Email:    request.Email,
		Password: string(password),
	}

	if err := u.UserRepository.Create(ctx, user); err != nil {
		return nil, fiber.NewError(fiber.StatusForbidden, string(err.Error()))
	}

	return ConvertUserResponse(user), nil
}

func (u *UserService) Login(ctx context.Context, request *LoginUserRequest) (*UserResponse, error) {
	user := new(User)
	if err := u.UserRepository.FindByEmail(ctx, user, request.Email); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	return ConvertUserResponse(user), nil
}

func (u *UserService) Update(ctx context.Context, request *UpdateUserRequest) (*UserResponse, error) {

	user := new(User)
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = string(password)

	if err := u.UserRepository.Update(ctx, user, request.UUID); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return ConvertUserResponse(user), nil
}

func (u *UserService) Delete(ctx context.Context, request *DeleteUserRequest) error {

	err := u.UserRepository.Delete(ctx, request.Email, request.Password)

	return err
}

func (u *UserService) FindUser(ctx context.Context, request *GetUserRequest) (*UserProfileResponse, error) {

	user := new(User)
	if err := u.UserRepository.FindUser(ctx, user, request.UUID); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	return ConvertUserProfileResponse(user), nil
}

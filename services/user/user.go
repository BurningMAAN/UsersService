package user

import (
	"UsersService/models"
	"context"
)

type repository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
}

type service struct {
	repository repository
}

func New(userRepository repository) *service {
	return &service{
		repository: userRepository,
	}
}

func (s *service) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (s *service) GetUser(ctx context.Context, userID string) (models.User, error) {
	return models.User{}, nil
}

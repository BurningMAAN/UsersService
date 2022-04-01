package user

import (
	"UsersService/models"
	"context"
)

type repository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUser(ctx context.Context, userID string) (models.User, error)
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
	return s.repository.CreateUser(ctx, user)
}

func (s *service) GetUser(ctx context.Context, userID string) (models.User, error) {
	return s.repository.GetUser(ctx, userID)
}

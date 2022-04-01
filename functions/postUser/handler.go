package main

import (
	"UsersService/models"
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type response struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type service interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
}
type handler struct {
	Service service
}

func (h *handler) CreateUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil
}

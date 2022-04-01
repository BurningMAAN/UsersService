package main

import (
	"UsersService/models"
	"context"
	"encoding/json"
	"net/http"

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
	request := request{}
	err := json.Unmarshal([]byte(event.Body), &request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	createdUser, err := h.Service.CreateUser(ctx, models.User{
		UserName:      request.UserName,
		CreditBalance: 0,
		Email:         request.Email,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	_ = response{
		ID:       createdUser.ID,
		UserName: createdUser.UserName,
		Email:    createdUser.Email,
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

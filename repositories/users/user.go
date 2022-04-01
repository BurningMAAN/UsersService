package user

import (
	"UsersService/models"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type UsersDB struct {
	PK             string
	SK             string
	UserName       string
	CreditsBalance float32
	Email          string
}

type DB interface {
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.GetItem
	GetItem(ctx context.Context, input *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.Query
	Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.PutItem
	PutItem(ctx context.Context, input *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.UpdateItem
	UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
}

type repository struct {
	DB DB
}

func New(db DB) *repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (r *repository) GetUser(ctx context.Context, userID string) (models.User, error) {
	return models.User{}, nil
}

func (r *repository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

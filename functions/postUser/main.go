package main

import (
	"context"
	"log"

	userRepository "UsersService/repositories/users"
	userService "UsersService/services/user"

	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kelseyhightower/envconfig"
)

type lambdaConfig struct {
	UserServiceTableName string `envconfig:"DYNAMODB_USER_DATA_TABLE" required:"true"`
}

func main() {
	var cfg lambdaConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed to read environment variables: %v", err)
	}

	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db := dynamodb.NewFromConfig(awsCfg)

	userRepository := userRepository.New(db)
	h := handler{
		Service: userService.New(userRepository),
	}
	lambda.Start(h.CreateUser)
}

package main

import (
	"context"
	"log"
	"os"

	"github.com/maximilienandile/backend-go-tuto/internal/storage"

	"github.com/aws/aws-lambda-go/lambda"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/aws/aws-lambda-go/events"

	"github.com/maximilienandile/backend-go-tuto/internal/server"
)

var ginLambda *ginadapter.GinLambda

func init() {
	allowedOrigin, found := os.LookupEnv("ALLOWED_ORIGIN")
	if !found {
		log.Fatal("env variable ALLOWED_ORIGIN was not found")
	}
	dynamoStorage, err := storage.NewDynamo("ecommerce-dev")
	if err != nil {
		log.Fatalf("impossible to create storage interface: %s", err)
	}
	myServer, err := server.New(server.Config{
		Port:          9090,
		AllowedOrigin: allowedOrigin,
		Storage:       dynamoStorage,
	})
	if err != nil {
		log.Fatalf("impossible to create the server: %s", err)
	}
	ginLambda = ginadapter.New(myServer.Engine)
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, event)
}

func main() {
	lambda.Start(Handler)
}

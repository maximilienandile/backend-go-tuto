package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"

	"github.com/maximilienandile/backend-go-tuto/internal/secret"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/maximilienandile/backend-go-tuto/internal/uniqueid"

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
	parameterStoreName, found := os.LookupEnv("PARAMETER_STORE_NAME")
	if !found {
		log.Fatal("env variable PARAMETER_STORE_NAME was not found")
	}

	ssmClient := ssm.New(session.Must(session.NewSession()))
	outSSM, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(parameterStoreName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("impossible to access SSM: %s", err)
	}
	parameterRawValue := *outSSM.Parameter.Value
	var secretsFromSSM secret.Parameters
	err = json.Unmarshal([]byte(parameterRawValue), &secretsFromSSM)
	if err != nil {
		log.Fatalf("impossible to unmarshall secrets: %s", err)
	}
	jsonCreds, err := json.Marshal(secretsFromSSM.Google)
	if err != nil {
		log.Fatalf("impossible to marshall Google secrets: %s", err)
	}
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsJSON(jsonCreds))
	if err != nil {
		log.Fatalf("impossible to create firebase app: %s", err)
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("impossible to build auth client: %s", err)
	}

	myServer, err := server.New(server.Config{
		Port:               9090,
		AllowedOrigin:      allowedOrigin,
		Storage:            dynamoStorage,
		UniqueIDGenerator:  uniqueid.UUIDV4{},
		FirebaseAuthClient: authClient,
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

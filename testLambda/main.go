package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	MonthNumber uint
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	return "ABC", nil
}

func main() {
	lambda.Start(HandleRequest)
}

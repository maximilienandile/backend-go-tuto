package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	MonthNumber uint
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	m := time.Month(event.MonthNumber)
	return m.String(), nil
}

func main() {
	lambda.Start(HandleRequest)
}

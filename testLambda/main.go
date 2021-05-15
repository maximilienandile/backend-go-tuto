package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	MonthNumber uint
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	if event.MonthNumber > 12 {
		return "", fmt.Errorf("the provided month number is greater than 12. Got : %d", event.MonthNumber)
	}
	m := time.Month(event.MonthNumber)
	return m.String(), nil
}

func main() {
	lambda.Start(HandleRequest)
}

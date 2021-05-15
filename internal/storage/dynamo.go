package storage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maximilienandile/backend-go-tuto/internal/product"
)

type Dynamo struct {
	tableName  string
	awsSession *session.Session
	client     *dynamodb.DynamoDB
}

func (d Dynamo) String() string {
	return "Dynamo"
}

func NewDynamo(tableName string) (*Dynamo, error) {
	awsSession, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf("impossible to create aws session: %w", err)
	}
	dynamodbClient := dynamodb.New(awsSession)
	return &Dynamo{
		tableName:  tableName,
		awsSession: awsSession,
		client:     dynamodbClient,
	}, nil
}

func (d *Dynamo) CreateProduct(product product.Product) error {
	item, err := dynamodbattribute.MarshalMap(product)
	if err != nil {
		return fmt.Errorf("impossible to marshall product: %w", err)
	}
	item["PK"] = &dynamodb.AttributeValue{
		S: aws.String("product"),
	}
	item["SK"] = &dynamodb.AttributeValue{
		S: aws.String(product.ID),
	}
	_, err = d.client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: &d.tableName,
	})
	if err != nil {
		return fmt.Errorf("impossible to Put item in db: %w", err)
	}
	return nil
}

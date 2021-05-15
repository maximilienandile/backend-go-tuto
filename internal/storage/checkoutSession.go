package storage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maximilienandile/backend-go-tuto/internal/checkout"
)

func (d *Dynamo) CreateCheckoutSession(session checkout.Session) error {
	item, err := dynamodbattribute.MarshalMap(session)
	if err != nil {
		return fmt.Errorf("impossible to marshall product: %w", err)
	}
	item[partitionKeyAttributeName] = &dynamodb.AttributeValue{
		S: aws.String(pkCheckoutSession),
	}
	item[sortKeyAttributeName] = &dynamodb.AttributeValue{
		S: aws.String(session.ID),
	}
	_, err = d.client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: &d.tableName,
	})
	if err != nil {
		return fmt.Errorf("impossible to Put checkout session in db: %w", err)
	}
	return nil
}

func (d *Dynamo) GetCheckoutSession(ID string) (checkout.Session, error) {
	// query the db
	// PK = checkoutSession
	// SK = userID
	out, err := d.getElementsByPKAndSK(pkCheckoutSession, ID)
	if err != nil {
		return checkout.Session{}, fmt.Errorf("impossible to retrieve the checkout session in db: %w", err)
	}
	if len(out.Items) == 0 {
		return checkout.Session{}, fmt.Errorf("no checkout session found %w", ErrNotFound)
	}
	if len(out.Items) > 1 {
		return checkout.Session{}, fmt.Errorf("retrieve more than one checkout session")
	}
	var session checkout.Session
	err = dynamodbattribute.UnmarshalMap(out.Items[0], &session)
	if err != nil {
		return checkout.Session{}, fmt.Errorf("impossible to unmarshall checkout Session: %w", err)
	}
	return session, nil
}

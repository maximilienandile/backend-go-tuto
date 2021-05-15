package storage

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/maximilienandile/backend-go-tuto/internal/product"
)

var ErrNotFound = errors.New("item not found")

func (d *Dynamo) UpdateInventory(productID string, delta int) error {
	// get product in db
	productFound, err := d.getProductByID(productID)
	if err != nil {
		return fmt.Errorf("impossible to retrive product %w", err)
	}
	// compute the new stock
	newStockValue := int(productFound.Stock) + delta
	if newStockValue < 0 {
		return fmt.Errorf("we cannot have a stock that is less than 0")
	}
	keyCondition := make(map[string]*dynamodb.AttributeValue)
	// PK
	keyCondition[partitionKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(pkProduct)}
	// SK
	keyCondition[sortKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(productID)}

	// condition expression
	condition := expression.Name("version").Equal(expression.Value(productFound.Version))
	// update (what needs to be updated)
	update := expression.Set(expression.Name("stock"), expression.Value(newStockValue))
	update.Set(expression.Name("version"), expression.Value(productFound.Version+1))

	// build expressions with expression builders
	builder := expression.NewBuilder().WithCondition(condition).WithUpdate(update)
	expr, err := builder.Build()
	if err != nil {
		return fmt.Errorf("impossible to build expression: %w", err)
	}
	input := dynamodb.UpdateItemInput{
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Key:                       keyCondition,
		TableName:                 &d.tableName,
		UpdateExpression:          expr.Update(),
	}
	_, err = d.client.UpdateItem(&input)
	if err != nil {
		return fmt.Errorf("impossible to run UpdateItem request: %w", err)
	}
	return nil
}

func (d *Dynamo) getProductByID(productID string) (product.Product, error) {
	getItemInput := dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		Key: map[string]*dynamodb.AttributeValue{
			partitionKeyAttributeName: {
				S: aws.String(pkProduct),
			},
			sortKeyAttributeName: {
				S: aws.String(productID),
			},
		},
		TableName: &d.tableName,
	}
	out, err := d.client.GetItem(&getItemInput)
	if err != nil {
		return product.Product{}, fmt.Errorf("impossible to GetItem: %w", err)
	}
	if len(out.Item) == 0 {
		return product.Product{}, ErrNotFound
	}
	var productFound product.Product
	err = dynamodbattribute.UnmarshalMap(out.Item, &productFound)
	if err != nil {
		return productFound, fmt.Errorf("impossible to marshall product: %s", err)
	}
	return productFound, nil
}

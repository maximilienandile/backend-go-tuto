package storage

import (
	"errors"
	"fmt"

	"github.com/maximilienandile/backend-go-tuto/internal/category"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maximilienandile/backend-go-tuto/internal/product"
)

const partitionKeyAttributeName = "PK"
const sortKeyAttributeName = "SK"
const pkProduct = "product"
const pkCategory = "category"
const pkCart = "cart"

var ErrNotFound = errors.New("element not found")

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
	item[partitionKeyAttributeName] = &dynamodb.AttributeValue{
		S: aws.String(pkProduct),
	}
	item[sortKeyAttributeName] = &dynamodb.AttributeValue{
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

func (d *Dynamo) Products() ([]product.Product, error) {
	out, err := d.getElementsByPK(pkProduct)
	if err != nil {
		return nil, fmt.Errorf("impossible to retrieve elements by PK: %w", err)
	}
	products := make([]product.Product, 0)
	err = dynamodbattribute.UnmarshalListOfMaps(out.Items, &products)
	if err != nil {
		return nil, fmt.Errorf("impossible to unmarshall results: %s", err)
	}
	return products, nil
}

func (d *Dynamo) CreateCategory(category category.Category) error {
	item, err := dynamodbattribute.MarshalMap(category)
	if err != nil {
		return fmt.Errorf("impossible to marshall category: %w", err)
	}
	item[partitionKeyAttributeName] = &dynamodb.AttributeValue{
		S: aws.String(pkCategory),
	}
	item[sortKeyAttributeName] = &dynamodb.AttributeValue{
		S: aws.String(category.ID),
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

func (d *Dynamo) Categories() ([]category.Category, error) {
	out, err := d.getElementsByPK(pkCategory)
	if err != nil {
		return nil, fmt.Errorf("impossible to retrieve elements by PK: %w", err)
	}
	categories := make([]category.Category, 0)
	err = dynamodbattribute.UnmarshalListOfMaps(out.Items, &categories)
	if err != nil {
		return nil, fmt.Errorf("impossible to unmarshall results: %w", err)
	}
	return categories, nil
}

func (d *Dynamo) getElementsByPK(pkAttributeValue string) (*dynamodb.QueryOutput, error) {
	// PK = :myValue
	keyCondition := expression.Key(partitionKeyAttributeName).Equal(expression.Value(pkAttributeValue))
	builder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("impossible to build expression: %s", err)
	}
	input := dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 &d.tableName,
	}
	out, err := d.client.Query(&input)
	if err != nil {
		return nil, fmt.Errorf("impossible to query database: %s", err)
	}
	return out, nil
}

func (d *Dynamo) getElementsByPkandSk(pkAttributeValue string, skAttributeValue string) (map[string]*dynamodb.AttributeValue, error) {
	result, err := d.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(d.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			partitionKeyAttributeName: {
				S: aws.String(pkAttributeValue),
			},
			sortKeyAttributeName: {
				S: aws.String(skAttributeValue),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(result.Item) == 0 {
		return nil, ErrNotFound
	}
	return result.Item, nil
}

func (d *Dynamo) getProductByID(productID string) (product.Product, error) {
	p := product.Product{}
	item, err := d.getElementsByPkandSk(pkProduct, productID)
	if err != nil {
		return p, err
	}
	err = dynamodbattribute.UnmarshalMap(item, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (d *Dynamo) UpdateInventory(productID string, delta int) error {
	// first read product by id
	// to retrieve current stock
	productDB, err := d.getProductByID(productID)
	if err != nil {
		return fmt.Errorf("impossible to retrieve product with id '%s': %s", productID, err)
	}
	// new stock
	newStockValue := int(productDB.Stock) + delta
	if newStockValue < 0 {
		return fmt.Errorf("impossible to have a stock that is less than 0, actual in db %d, delta %d, newValue %d", productDB.Stock, delta, newStockValue)
	}
	// Build Update Item
	// KEY Condition
	// 		PK = product
	// 		SK = productID
	keyCondition := map[string]*dynamodb.AttributeValue{}
	keyCondition[partitionKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(pkProduct)}
	keyCondition[sortKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(productID)}
	// Condition
	condition := expression.Name("version").Equal(expression.Value(productDB.Version))
	// Update (what needs to be updated ?)
	// update the stock
	update := expression.Set(expression.Name("stock"), expression.Value(int(productDB.Stock)+delta))
	// increment the version
	update = update.Set(expression.Name("version"), expression.Value(productDB.Version+1))

	builder := expression.NewBuilder().WithCondition(condition).WithUpdate(update)
	expr, err := builder.Build()
	if err != nil {
		return fmt.Errorf("impossible to build expression: %s", err)
	}
	input := dynamodb.UpdateItemInput{
		Key:                       keyCondition,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		TableName:                 &d.tableName,
	}
	_, err = d.client.UpdateItem(&input)
	if err != nil {
		return fmt.Errorf("impossible to query database: %s", err)
	}
	return nil
}

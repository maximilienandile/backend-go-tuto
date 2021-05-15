package storage

import (
	"fmt"

	"github.com/maximilienandile/backend-go-tuto/internal/cart"

	"github.com/maximilienandile/backend-go-tuto/internal/category"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/maximilienandile/backend-go-tuto/internal/product"
)

const (
	partitionKeyAttributeName = "PK"
	sortKeyAttributeName      = "SK"
	pkProduct                 = "product"
	pkCategory                = "category"
	pkCart                    = "cart"
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

func (d *Dynamo) getElementsByPKAndSK(pkAttributeValue, skAttributeValue string) (*dynamodb.QueryOutput, error) {
	// PK = :myValue
	keyCondition := expression.Key(partitionKeyAttributeName).Equal(expression.Value(pkAttributeValue))
	sortKeyCondition := expression.Key(sortKeyAttributeName).Equal(expression.Value(skAttributeValue))
	keyCondition.And(sortKeyCondition)
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

func (d *Dynamo) UpdateProduct(input UpdateProductInput) error {
	// get product in db
	productFound, err := d.getProductByID(input.ProductID)
	if err != nil {
		return fmt.Errorf("impossible to retrive product %w", err)
	}
	// key condition
	keyCondition := make(map[string]*dynamodb.AttributeValue)
	// PK
	keyCondition[partitionKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(pkProduct)}
	// SK
	keyCondition[sortKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(input.ProductID)}

	// condition expression
	condition := expression.Name("version").Equal(expression.Value(productFound.Version))

	// update expression (what we need to update)
	update := expression.Set(expression.Name("name"), expression.Value(input.Name))
	update.Set(expression.Name("version"), expression.Value(productFound.Version+1))
	update.Set(expression.Name("image"), expression.Value(input.Image))
	update.Set(expression.Name("shortDescription"), expression.Value(input.ShortDescription))
	update.Set(expression.Name("priceVatExcluded"), expression.Value(input.PriceVATExcluded))
	update.Set(expression.Name("vat"), expression.Value(input.VAT))
	update.Set(expression.Name("totalPrice"), expression.Value(input.TotalPrice))

	// build expressions with expression builders
	builder := expression.NewBuilder().WithCondition(condition).WithUpdate(update)
	expr, err := builder.Build()
	if err != nil {
		return fmt.Errorf("impossible to build expression: %w", err)
	}
	// request UpdateItem
	inputDB := dynamodb.UpdateItemInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Key:                       keyCondition,
		TableName:                 &d.tableName,
		ConditionExpression:       expr.Condition(),
		UpdateExpression:          expr.Update(),
	}
	_, err = d.client.UpdateItem(&inputDB)
	if err != nil {
		return fmt.Errorf("impossible to run UpdateItem request: %w", err)
	}
	return nil
}

func (d *Dynamo) CreateCart(cart cart.Cart, userID string) error {
	item, err := dynamodbattribute.MarshalMap(cart)
	if err != nil {
		return fmt.Errorf("impossible to marshall product: %w", err)
	}
	item[partitionKeyAttributeName] = &dynamodb.AttributeValue{
		S: aws.String(pkCart),
	}
	item[sortKeyAttributeName] = &dynamodb.AttributeValue{
		S: aws.String(userID),
	}
	_, err = d.client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: &d.tableName,
	})
	if err != nil {
		return fmt.Errorf("impossible to Put cart in db: %w", err)
	}
	return nil
}

func (d *Dynamo) GetCart(userID string) (cart.Cart, error) {
	// query the db
	// PK = cart
	// SK = userID
	out, err := d.getElementsByPKAndSK(pkCart, userID)
	if err != nil {
		return cart.Cart{}, fmt.Errorf("impossible to retrieve the cart in db: %w", err)
	}
	if len(out.Items) == 0 {
		return cart.Cart{}, fmt.Errorf("no cart found %w", ErrNotFound)
	}
	if len(out.Items) > 1 {
		return cart.Cart{}, fmt.Errorf("retrieve more than one item")
	}
	var cartRetrieved cart.Cart
	err = dynamodbattribute.UnmarshalMap(out.Items[0], &cartRetrieved)
	if err != nil {
		return cart.Cart{}, fmt.Errorf("impossible to unmarshall cart: %w", err)
	}
	return cartRetrieved, nil
}

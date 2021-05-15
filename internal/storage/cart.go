package storage

import (
	"errors"
	"fmt"

	"github.com/maximilienandile/backend-go-tuto/internal/product"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	uuid "github.com/satori/go.uuid"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maximilienandile/backend-go-tuto/internal/cart"
)

func (d *Dynamo) CreateOrUpdateCart(userID string, productID string, delta int) (cart.Cart, error) {
	// read Cart
	cartFound, err := d.GetCart(userID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			cartFound = cart.Cart{Version: 1}
			// no cart found, let's create one
			err = d.createCart(cartFound, userID)
			if err != nil {
				return cart.Cart{}, fmt.Errorf("impossible to create new cart : %s", userID)
			}
		} else {
			return cart.Cart{}, fmt.Errorf("impossible to get cart for userID : %s", userID)
		}
	}
	err = cartFound.UpsertItem(productID, delta)
	if err != nil {
		return cart.Cart{}, fmt.Errorf("impossible to update items of cart: %s", err)
	}
	// read product
	productDB, err := d.getProductByID(productID)
	if err != nil {
		return cart.Cart{}, fmt.Errorf("impossible to get product with ID : %s", productID)
	}
	// transaction
	// prepare a slice of operations in the transaction
	items := make([]*dynamodb.TransactWriteItem, 0)
	// 1. Update the stock of the product
	updateStockReq, err := d.buildUpdateStockReq(productDB, delta)
	if err != nil {
		return cart.Cart{}, fmt.Errorf("impossible to update items of cart: %s", err)
	}
	items = append(items, updateStockReq)
	// 2. update the cart
	updateCartReq, err := d.buildUpdateCartReq(cartFound, userID)
	if err != nil {
		return cart.Cart{}, fmt.Errorf("impossible to update items of cart: %s", err)
	}
	items = append(items, updateCartReq)
	_, err = d.client.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		ClientRequestToken: aws.String(uuid.NewV4().String()),
		TransactItems:      items,
	})
	if err != nil {
		return cart.Cart{}, fmt.Errorf("imposible to run transaction: %s", err)
	}
	return cartFound, nil
}

func (d *Dynamo) buildUpdateStockReq(productDB product.Product, delta int) (*dynamodb.TransactWriteItem, error) {
	// if delta > 0 (add to cart) we remove from stock the delta
	// oldStock: 12; delta:2
	// ex: 12 - 2 = 10

	// if delta < 0 (remove from cart) we add to stock the delta
	// oldStock: 12; delta:-3
	// ex: 12 - (-3) = 12 + 3 = 15
	newStock := int(productDB.Stock) - delta

	// if delta > 0 (add to cart) we add the delta to the reserved item
	// oldReserved:5;  delta:2
	// ex: 5 + 2 = 7

	// if delta < 0 (remove from cart) we remove the delta from the reserved item
	// oldReserved:5;  delta:-3
	// ex: 5 + (-3) = 2
	newReserved := int(productDB.Reserved) + delta
	// check that we cannot have a negative reserved amount or stock
	if newReserved < 0 || newStock < 0 {
		return nil, fmt.Errorf("impossible to have a stock value or reserved values less than 0")
	}
	keyCondition := map[string]*dynamodb.AttributeValue{}
	keyCondition[partitionKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(pkProduct)}
	keyCondition[sortKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(productDB.ID)}
	// Condition
	condition := expression.Name("version").Equal(expression.Value(productDB.Version))
	// update the stock
	update := expression.Set(expression.Name("stock"), expression.Value(newStock))
	// update item reserved
	update = update.Set(expression.Name("reserved"), expression.Value(newReserved))
	// increment the version
	update = update.Set(expression.Name("version"), expression.Value(productDB.Version+1))

	builder := expression.NewBuilder().WithCondition(condition).WithUpdate(update)
	expr, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("impossible to build expression: %s", err)
	}
	updateStock := &dynamodb.TransactWriteItem{
		Update: &dynamodb.Update{
			Key:                       keyCondition,
			ConditionExpression:       expr.Condition(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			TableName:                 &d.tableName,
		},
	}
	return updateStock, nil
}

func (d *Dynamo) buildUpdateCartReq(cartUpdated cart.Cart, userID string) (*dynamodb.TransactWriteItem, error) {
	keyCondition2 := map[string]*dynamodb.AttributeValue{}
	keyCondition2[partitionKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(pkCart)}
	keyCondition2[sortKeyAttributeName] = &dynamodb.AttributeValue{S: aws.String(userID)}
	// Condition
	condition2 := expression.Name("version").Equal(expression.Value(cartUpdated.Version))
	// Update (what needs to be updated ?)
	// update the stock
	update2 := expression.Set(expression.Name("items"), expression.Value(cartUpdated.Items))
	// increment the version
	update2 = update2.Set(expression.Name("version"), expression.Value(cartUpdated.Version+1))

	builder2 := expression.NewBuilder().WithCondition(condition2).WithUpdate(update2)
	expr2, err := builder2.Build()
	if err != nil {
		return nil, fmt.Errorf("impossible to build expression: %s", err)
	}
	updateCart := &dynamodb.TransactWriteItem{
		Update: &dynamodb.Update{
			Key:                       keyCondition2,
			ConditionExpression:       expr2.Condition(),
			ExpressionAttributeNames:  expr2.Names(),
			ExpressionAttributeValues: expr2.Values(),
			UpdateExpression:          expr2.Update(),
			TableName:                 &d.tableName,
		},
	}
	return updateCart, nil
}

func (d *Dynamo) GetCart(userID string) (cart.Cart, error) {
	c := cart.Cart{}
	item, err := d.getElementsByPkandSk(pkCart, userID)
	if err != nil {
		return c, err
	}
	err = dynamodbattribute.UnmarshalMap(item, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (d *Dynamo) createCart(cart cart.Cart, userID string) error {
	item, err := dynamodbattribute.MarshalMap(cart)
	if err != nil {
		return fmt.Errorf("impossible to marshall cart: %w", err)
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
		return fmt.Errorf("impossible to Put item in db: %w", err)
	}
	return nil
}

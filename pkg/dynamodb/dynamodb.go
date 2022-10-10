package db

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoConfig struct {
	DBService  *dynamodb.Client
	PrimaryKey string
	SortKey    string
	TableName  string
}

const (
	maxBatch      int    = 25
	status        string = "Status"
	errorMarshall string = "Couldn't unmarshal response. Here's why: %v"
)

func NewDynamoDB(tn, pk, sk string) *DynamoConfig {
	dbCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(config.DefaultSharedConfigProfile))
	if err != nil {
		log.Fatalf("Unable to LoadDB")
	}

	return &DynamoConfig{
		DBService:  dynamodb.NewFromConfig(dbCfg),
		PrimaryKey: pk,
		SortKey:    sk,
		TableName:  tn,
	}
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		log.Fatalf("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func Chunk(array []interface{}, chunkSize int) [][]interface{} {
	var divided [][]interface{}

	for i := 0; i < len(array); i += chunkSize {
		end := i + chunkSize
		if end > len(array) {
			end = len(array)
		}

		divided = append(divided, array[i:end])
	}

	return divided
}

func (dbc *DynamoConfig) Save(prop interface{}) (interface{}, error) {
	av, err := attributevalue.MarshalMap(prop)
	if err != nil {
		log.Printf(errorMarshall, err)

		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(dbc.TableName),
	}

	_, err = dbc.DBService.PutItem(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("error saving item - put - %w", err)
	}

	return prop, err
}

func (dbc *DynamoConfig) SaveMany(data interface{}) error {
	batches := Chunk(InterfaceSlice(data), maxBatch)

	for _, dataArray := range batches {
		items := []types.WriteRequest{}

		for _, item := range dataArray {
			av, err := attributevalue.MarshalMap(item)
			if err != nil {
				log.Printf(errorMarshall, err)

				return err
			}

			items = append(items, types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: av,
				},
			})
		}

		bwii := &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				dbc.TableName: items,
			},
		}

		_, err := dbc.DBService.BatchWriteItem(context.TODO(), bwii)
		if err != nil {
			return fmt.Errorf("error savemany - batchwrite - %w", err)
		}
	}

	return nil
}

func (dbc *DynamoConfig) Delete(prop interface{}) (interface{}, error) {
	av, err := attributevalue.MarshalMap(prop)
	if err != nil {
		log.Printf(errorMarshall, err)

		return nil, err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       av,
		TableName: aws.String(dbc.TableName),
	}

	_, err = dbc.DBService.DeleteItem(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("error delete - deleteitem - %w", err)
	}

	return prop, err
}

func (dbc *DynamoConfig) Get(pk, sk string, data interface{}) error {
	primaryKey := map[string]string{
		dbc.PrimaryKey: pk,
	}

	if sk != "" {
		primaryKey[dbc.SortKey] = sk
	}

	key, err := attributevalue.MarshalMap(primaryKey)
	if err != nil {
		return err
	}

	response, err := dbc.DBService.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(dbc.TableName),
	})
	if err != nil {
		log.Printf("Couldn't get info about. Here's why: %v\n", err)

		return err
	}

	err = attributevalue.UnmarshalMap(response.Item, &data)
	if err != nil {
		log.Printf(errorMarshall, err)

		return err
	}

	return err
}

func (dbc *DynamoConfig) Update(item, pk, sk string) error {
	primaryKey := map[string]string{
		dbc.PrimaryKey: pk,
		dbc.SortKey:    sk,
	}

	key, err := attributevalue.MarshalMap(primaryKey)
	if err != nil {
		return err
	}

	upd := expression.
		Set(expression.Name(status), expression.Value(item))

	expr, err := expression.NewBuilder().WithUpdate(upd).Build()
	if err != nil {
		return err
	}

	input := &dynamodb.UpdateItemInput{
		TableName:                 &dbc.TableName,
		Key:                       key,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err = dbc.DBService.UpdateItem(context.TODO(), input)

	return err
}

func (dbc *DynamoConfig) UpdateMany(pk, value string, sks []string) error {
	upd := expression.
		Set(expression.Name(status), expression.Value(value))

	expr, errBuilder := expression.NewBuilder().WithUpdate(upd).Build()
	if errBuilder != nil {
		return errBuilder
	}

	transactItems := make([]types.TransactWriteItem, len(sks))

	for idx, itemID := range sks {
		primaryKey := map[string]string{
			dbc.PrimaryKey: pk,
			dbc.SortKey:    itemID,
		}

		key, err := attributevalue.MarshalMap(primaryKey)
		if err != nil {
			return err
		}

		update := types.Update{
			TableName:                 aws.String(dbc.TableName),
			Key:                       key,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
		}
		transactItems[idx] = types.TransactWriteItem{Update: &update}
	}

	_, err := dbc.DBService.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{
		TransactItems: transactItems,
	})
	if err != nil {
		return err
	}

	return err
}

func (dbc *DynamoConfig) Query(pk string, data interface{}) error {
	var err error

	var response *dynamodb.QueryOutput

	keyEx := expression.Key(dbc.PrimaryKey).Equal(expression.Value(pk))

	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Printf("Couldn't build epxression for query. Here's why: %v\n", err)

		return err
	}

	response, err = dbc.DBService.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(dbc.TableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		log.Printf("Couldn't query for pk = %v. Here's why: %v\n", keyEx, err)

		return err
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &data)
	if err != nil {
		log.Printf(errorMarshall, err)

		return err
	}

	return err
}

func (dbc *DynamoConfig) QueryByGSI(value, indexName, indexPk string, data interface{}) error {
	var err error

	var response *dynamodb.QueryOutput

	keyEx := expression.Key(indexPk).Equal(expression.Value(value))

	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Printf("Couldn't build epxression for query. Here's why: %v\n", err)

		return err
	}

	response, err = dbc.DBService.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(dbc.TableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		IndexName:                 &indexName,
	})
	if err != nil {
		log.Printf("Couldn't query for GSI = %v. Here's why: %v\n", indexName, err)

		return err
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &data)
	if err != nil {
		log.Printf(errorMarshall, err)

		return err
	}

	return err
}

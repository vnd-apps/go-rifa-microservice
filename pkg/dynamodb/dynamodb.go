// Package mongodb implements mongodb connection.
package db

import (
	"fmt"
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoConfig struct {
	DBService  *dynamodb.DynamoDB
	PrimaryKey string
	SortKey    string
	TableName  string
}

const (
	maxBatch = 25
)

func NewDynamoDB(tn, pk, sk string) *DynamoConfig {
	dbSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &DynamoConfig{
		DBService:  dynamodb.New(dbSession),
		PrimaryKey: pk,
		SortKey:    sk,
		TableName:  tn,
	}
}

func (dbc *DynamoConfig) Save(prop interface{}) (interface{}, error) {
	av, err := dynamodbattribute.MarshalMap(prop)
	if err != nil {
		return nil, fmt.Errorf("error saving item - marshal - %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(dbc.TableName),
	}

	_, err = dbc.DBService.PutItem(input)
	if err != nil {
		return nil, fmt.Errorf("error saving item - put - %w", err)
	}

	return prop, err
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

func (dbc *DynamoConfig) SaveMany(data interface{}) error {
	batches := Chunk(InterfaceSlice(data), maxBatch)

	for _, dataArray := range batches {
		items := make([]*dynamodb.WriteRequest, len(dataArray))

		for i, item := range dataArray {
			av, err := dynamodbattribute.MarshalMap(item)
			if err != nil {
				return fmt.Errorf("error savemany - marshal - %w", err)
			}

			items[i] = &dynamodb.WriteRequest{
				PutRequest: &dynamodb.PutRequest{
					Item: av,
				},
			}
		}

		bwii := &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				dbc.TableName: items,
			},
		}

		_, err := dbc.DBService.BatchWriteItem(bwii)
		if err != nil {
			return fmt.Errorf("error savemany - batchwrite - %w", err)
		}
	}

	return nil
}

func (dbc *DynamoConfig) Delete(prop interface{}) (interface{}, error) {
	av, err := dynamodbattribute.MarshalMap(prop)
	if err != nil {
		return nil, fmt.Errorf("error delete - marshal - %w", err)
	}

	input := &dynamodb.DeleteItemInput{
		Key:       av,
		TableName: aws.String(dbc.TableName),
	}

	_, err = dbc.DBService.DeleteItem(input)
	if err != nil {
		return nil, fmt.Errorf("error delete - deleteitem - %w", err)
	}

	return prop, err
}

func (dbc *DynamoConfig) Get(pk, sk string, data interface{}) error {
	av := map[string]*dynamodb.AttributeValue{
		dbc.PrimaryKey: {
			S: aws.String(pk),
		},
	}

	if sk != "" {
		av[dbc.SortKey] = &dynamodb.AttributeValue{
			S: aws.String(sk),
		}
	}

	result, err := dbc.DBService.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(dbc.TableName),
		Key:       av,
	})
	if err != nil {
		return fmt.Errorf("NOT FOUND, %w", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, data)
	if err != nil {
		return fmt.Errorf("unmarshalMap, %w", err)
	}

	return err
}

func (dbc *DynamoConfig) FindStartingWith(pk, value string, data interface{}) error {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(dbc.TableName),
		KeyConditions: map[string]*dynamodb.Condition{
			dbc.PrimaryKey: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(pk),
					},
				},
			},
			dbc.SortKey: {
				ComparisonOperator: aws.String("BEGINS_WITH"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(value),
					},
				},
			},
		},
	}

	result, err := dbc.DBService.Query(queryInput)
	if err != nil {
		return fmt.Errorf("DB:FindStartingWith> NOT FOUND, %w", err)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal record, %w", err)
	}

	return err
}

func (dbc *DynamoConfig) FindByGsi(value, indexName, indexPk string, data interface{}) error {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(dbc.TableName),
		IndexName: aws.String(indexName),
		KeyConditions: map[string]*dynamodb.Condition{
			indexPk: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(value),
					},
				},
			},
		},
	}

	result, err := dbc.DBService.Query(queryInput)
	if err != nil {
		return fmt.Errorf("DB:QUERY NOT FOUND, %w", err)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal record, %w", err)
	}

	return err
}

func (dbc *DynamoConfig) GetAll(pk string, data interface{}) error {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(dbc.TableName),
		KeyConditions: map[string]*dynamodb.Condition{
			dbc.PrimaryKey: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(pk),
					},
				},
			},
		},
	}

	result, err := dbc.DBService.Query(queryInput)
	if err != nil {
		return fmt.Errorf("NOT FOUND, %w", err)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal record, %w", err)
	}

	return err
}

func (dbc *DynamoConfig) GetProduct(pk string, data, dataItem interface{}) error {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(dbc.TableName),
		KeyConditions: map[string]*dynamodb.Condition{
			dbc.PrimaryKey: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(pk),
					},
				},
			},
		},
	}

	result, err := dbc.DBService.Query(queryInput)
	if err != nil {
		return fmt.Errorf("NOT FOUND, %w", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Items[aws.Int64Value(result.Count)-1], data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal record, %w", err)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items[0:aws.Int64Value(result.Count)-1], dataItem)
	if err != nil {
		return fmt.Errorf("failed to unmarshal record, %w", err)
	}

	return err
}

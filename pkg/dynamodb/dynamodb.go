// Package mongodb implements mongodb connection.
package db

import (
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
	marshalError = "Failed to unmarshal Record, %v"
	maxBatch     = 25
)

// init setup the session and define table name, primary key and sort key.
func NewDynamoDB(tn, pk, sk string) *DynamoConfig {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	dbSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create DynamoDB client
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
		log.Fatalf(err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(dbc.TableName),
	}

	_, err = dbc.DBService.PutItem(input)
	if err != nil {
		log.Fatalf(err.Error())
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
				log.Fatalf("Got error calling DeleteItem, %v", err)
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
			return err
		}
	}

	return nil
}

func (dbc *DynamoConfig) Delete(prop interface{}) (interface{}, error) {
	av, err := dynamodbattribute.MarshalMap(prop)
	if err != nil {
		log.Fatalf(err.Error())
	}

	input := &dynamodb.DeleteItemInput{
		Key:       av,
		TableName: aws.String(dbc.TableName),
	}

	_, err = dbc.DBService.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem, %v", err)
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
		log.Fatalf("NOT FOUND, %v", err)

		return err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, data)
	if err != nil {
		log.Fatalf(marshalError, err)
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
		log.Fatalf("DB:FindStartingWith> NOT FOUND, %v", err)

		return err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
		log.Fatalf("Failed to unmarshal Record, %v", err)
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
		log.Fatalf("DB:QUERY NOT FOUND, %v", err)

		return err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
		log.Fatalf(marshalError, err)
	}

	return err
}

func (dbc *DynamoConfig) FindAll(data interface{}) error {
	params := &dynamodb.ScanInput{
		TableName: aws.String(dbc.TableName),
	}

	// Make the DynamoDB Query API call
	result, err := dbc.DBService.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
		log.Fatalf(marshalError, err)
	}

	return err
}

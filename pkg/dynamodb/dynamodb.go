// Package mongodb implements mongodb connection.
package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DBConfig struct {
	DBService  *dynamodb.DynamoDB
	PrimaryKey string
	SortKey    string
	TableName  string
}

// init setup the session and define table name, primary key and sort key.
func NewDynamoDB(tn, pk, sk string) *DBConfig {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	dbSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create DynamoDB client
	return &DBConfig{
		DBService:  dynamodb.New(dbSession),
		PrimaryKey: pk,
		SortKey:    sk,
		TableName:  tn,
	}
}

func (dbc *DBConfig) Save(prop interface{}) (interface{}, error) {
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

func (dbc *DBConfig) Delete(prop interface{}) (interface{}, error) {
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
		log.Fatalf("Got error calling DeetItem:")
		log.Fatalf(err.Error())
	}

	return prop, err
}

func (dbc *DBConfig) Get(pk, sk string, data interface{}) error {
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
		log.Fatalf("NOT FOUND")
		log.Fatalf(err.Error())

		return err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, data)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return err
}

func (dbc *DBConfig) FindStartingWith(pk, value string, data interface{}) error {
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
		log.Fatalf("DB:FindStartingWith> NOT FOUND")
		log.Fatalf(err.Error())

		return err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return err
}

func (dbc *DBConfig) FindByGsi(value, indexName, indexPk string, data interface{}) error {
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
		log.Fatalf("NOT FOUND")
		log.Fatalf(err.Error())

		return err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, data)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return err
}

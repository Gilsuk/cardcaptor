package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewDDB returns DynamoDB connection
func NewDDB(region, akid, secret string) *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(akid, secret, ""),
	}))
	return dynamodb.New(sess)
}

// CreateDDBTable is
func CreateDDBTable(db *dynamodb.DynamoDB, tableName string, rcu, wcu int64) error {

	input := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String("PROVISIONED"),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(rcu),
			WriteCapacityUnits: aws.Int64(wcu),
		},
	}

	_, err := db.CreateTable(input)

	return err
}

// PutItem store item to DynamoDB
func (c *CardItem) PutItem(db *dynamodb.DynamoDB, tableName string) (err error) {
	av, err := dynamodbattribute.MarshalMap(c)
	if err != nil {
		return
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = db.PutItem(input)
	return
}

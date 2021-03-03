package repository

import (
	"fmt"
	"os"

	"code.qburst.com/navaneeth.k/DynamoDB-example/config"
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetRecord(movieName string, movieYear string) {
	svc := config.Connect()

	tableName := "Movies"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(movieYear),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		
		os.Exit(1)
	}
	if result.Item == nil {
		fmt.Println("Not found")
		os.Exit(1)
	}

	movie := models.Movie{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &movie)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println(movie)
}

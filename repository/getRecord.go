package repository

import (
	"errors"
	"fmt"

	"code.qburst.com/navaneeth.k/DynamoDB-example/config"
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetRecord(movieName string, movieYear string) (models.Movie, error) {
	svc := config.Connect()

	tableName := "Movies"
	movie := models.Movie{}

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

		return movie, err
	}
	if result.Item == nil {
		fmt.Println("Not found")
		err = errors.New("Not found")
		return movie, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &movie)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		return movie, err
	}

	//fmt.Println(movie)
	return movie, err
}

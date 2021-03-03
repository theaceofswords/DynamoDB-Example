package repository

import (
	"fmt"
	"os"
	"strconv"

	"code.qburst.com/navaneeth.k/DynamoDB-example/config"
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AddRecord(movie models.Movie) {
	svc := config.Connect()

	// marshall that data into a map of AttributeValue objects.
	av, err := dynamodbattribute.MarshalMap(movie)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create the input for PutItem and call it.
	tableName := "Movies"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	year := strconv.Itoa(movie.Year)

	fmt.Println("Successfully added '" + movie.Title + "' (" + year + ") to table " + tableName)

}

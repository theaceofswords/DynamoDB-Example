package repository

import (
	"fmt"
	"os"

	"code.qburst.com/navaneeth.k/DynamoDB-example/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteRecord() {
	svc := config.Connect()

	tableName := "Movies"
	movieName := "The Big New Movie"
	movieYear := "2015"

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(movieYear),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Error Deleting Item")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Deleted")
}

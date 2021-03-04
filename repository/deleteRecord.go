package repository

import (
	"fmt"

	"code.qburst.com/navaneeth.k/DynamoDB-example/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteRecord(movieName string, movieYear string) error {
	svc := config.Connect()

	tableName := "Movies"
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
		return err
	}

	fmt.Println("Deleted")
	return nil
}

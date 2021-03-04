package repository

import (
	"fmt"
	"strconv"

	"code.qburst.com/navaneeth.k/DynamoDB-example/config"
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UpdateRecord(movie models.Movie) {
	svc := config.Connect()
	tableName := "Movies"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(strconv.FormatFloat(movie.Rating, 'E', -1, 64)),
			},
			":d": {
				S: aws.String(movie.Director),
			},
			":a": {
				S: aws.String(movie.LeadActor),
			},
			":c": {
				S: aws.String(movie.Category),
			},
			":p": {
				S: aws.String(movie.Plot),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(strconv.Itoa(movie.Year)),
			},
			"Title": {
				S: aws.String(movie.Title),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rating = :r, Director = :d, LeadActor = :a, Category = :c, Plot = :p"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated")

}

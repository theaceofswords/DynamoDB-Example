package repository

import (
	"fmt"
	"strconv"

	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
)

func (r *repo) UpdateRecord(movie models.Movie) {

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

	_, err := r.svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully updated")

	shardInput := &dynamodbstreams.GetShardIteratorInput{
		StreamArn:         aws.String("arn:aws:dynamodb:ddblocal:000000000000:table/Movies/stream/2021-03-10T04:25:19.082"),
		ShardId:           aws.String("shardId-00000001615350319121-a639bb13"),
		ShardIteratorType: aws.String("LATEST"),
	}
	shardId, err := r.svc2.GetShardIterator(shardInput)

	if err != nil {
		fmt.Println(err.Error())
	}
	str := fmt.Sprintf("%v", shardId)
	fmt.Println(str)

	recInput := &dynamodbstreams.GetRecordsInput{
		//ShardIterator: aws.String("arn:aws:dynamodb:ddblocal:000000000000:table/Movies/stream/2021-03-10T04:25:19.082|001|c2hhcmRJZC0wMDAwMDAwMTYxNTM1MDMxOTEyMS1hNjM5YmIxM3wwMDAwMDAwMDAwMDAwMDAwMDAwMDV8MDAwMDAwMDAwMDAwMDAwMDAxNjE1MzUwNTAwNTU3"),
		ShardIterator: aws.String(str),
	}

	result, err := r.svc2.GetRecords(recInput)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
	fmt.Printf("%T\n", result)

}

package repository

import (
	"fmt"
	"strconv"

	// "code.qburst.com/navaneeth.k/DynamoDB-example/config"
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	
)

func (r *repo) UpdateRecord(movie models.Movie) {



	// tableDescrIp := &dynamodb.DescribeTableInput{
	// 	TableName: aws.String(tableName),
	// }

	// tblDescr, err := r.svc.DescribeTable(tableDescrIp)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// streamArn := *tblDescr.Table.LatestStreamArn
	// fmt.Println(streamArn)

	// streamDescIp := &dynamodbstreams.DescribeStreamInput{
	// 	StreamArn: aws.String(streamArn),
	// }
	// //fmt.Println(streamDescIp)

	// streamDescr, err := r.svc2.DescribeStream(streamDescIp)
	// //fmt.Println(streamDescr)
	// shardId := *streamDescr.StreamDescription.Shards[0].ShardId
	// fmt.Println(shardId)

	// shardInput := &dynamodbstreams.GetShardIteratorInput{
	// 	StreamArn:         aws.String(streamArn),
	// 	ShardId:           aws.String(shardId),
	// 	ShardIteratorType: aws.String("LATEST"),
	// }
	// shardIterator, err := r.svc2.GetShardIterator(shardInput)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }


    var shardIterator string
    if nextIterator == ""{
		fmt.Println("new iterator")
		shardIterator = r.GetIterator()
	}else{
		
		shardIterator = nextIterator
	}


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

	r.psqlWrite(shardIterator)

	
	// str := fmt.Sprintf("%v", shardIterator)

	//fmt.Println(*shardIterator.ShardIterator)

	// recInput := &dynamodbstreams.GetRecordsInput{
	// 	ShardIterator: aws.String(shardIterator),
	// 	//ShardIterator: aws.String("arn:aws:dynamodb:ddblocal:000000000000:table/Movies/stream/2021-03-10T14:57:44.267|001|c2hhcmRJZC0wMDAwMDAwMTYxNTM4ODI2NDMwMC02MWFmM2VjMHwwMDAwMDAwMDAwMDAwMDAwMDAwMDZ8MDAwMDAwMDAwMDAwMDAwMDAxNjE1Mzg5NzU1MTY2"),
	// }
	// //recInput.SetShardIterator("arn:aws:dynamodb:ddblocal:000000000000:table/Movies/stream/2021-03-10T14:57:44.267|001|c2hhcmRJZC0wMDAwMDAwMTYxNTM4ODI2NDMwMC02MWFmM2VjMHwwMDAwMDAwMDAwMDAwMDAwMDAwMDZ8MDAwMDAwMDAwMDAwMDAwMDAxNjE1Mzg4MjY5MDE0")

	// result, err := r.svc2.GetRecords(recInput)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// nextIterator = *result.NextShardIterator

	// fmt.Println(result)
	// strc := models.Records{
	// 	OldImage: fmt.Sprintf("%v",result.Records[0].Dynamodb.OldImage),
	// 	NewImage: fmt.Sprintf("%v",result.Records[0].Dynamodb.NewImage),
	// 	EventId: *result.Records[0].EventID,
	// 	EventName: result.Records[0].EventName,
	// }

	// // psqlDB := config.PsqlConnect()
	// // defer psqlDB.Close()

	// err = r.psqlDB.Create(&strc).Error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

}

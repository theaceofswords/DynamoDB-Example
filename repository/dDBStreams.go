package repository

import (
	"fmt"

	
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
)

func (r *repo)GetIterator() string{
	tableDescrIp := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	tblDescr, err := r.svc.DescribeTable(tableDescrIp)
	if err != nil {
		fmt.Println(err.Error())
	}
	streamArn := *tblDescr.Table.LatestStreamArn
	fmt.Println(streamArn)

	streamDescIp := &dynamodbstreams.DescribeStreamInput{
		StreamArn: aws.String(streamArn),
	}
	//fmt.Println(streamDescIp)
	var shardId string

	streamDescr, err := r.svc2.DescribeStream(streamDescIp)
	//fmt.Println(streamDescr)
	for i,id := range streamDescr.StreamDescription.Shards{
		
		fmt.Println("************************************************=>",i)
	    shardId = *id.ShardId
	}
	fmt.Println(shardId)

	shardInput := &dynamodbstreams.GetShardIteratorInput{
		StreamArn:         aws.String(streamArn),
		ShardId:           aws.String(shardId),
		ShardIteratorType: aws.String("LATEST"),
	}
	shardIterator, err := r.svc2.GetShardIterator(shardInput)

	if err != nil {
		fmt.Println(err.Error())
	}
    fmt.Println("Done")
	return *shardIterator.ShardIterator

}
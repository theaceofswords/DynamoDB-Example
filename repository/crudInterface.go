package repository

import (
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
)

type CRUD interface {
	AddRecord(movie models.Movie)
	GetRecord(movieName string, movieYear string) (models.Movie, error)
	UpdateRecord(movie models.Movie)
	DeleteRecord(movieName string, movieYear string) error
	InitaliseData()
	CreateTable()
}

type repo struct {
	svc  *dynamodb.DynamoDB
	svc2 *dynamodbstreams.DynamoDBStreams
}

func CreateRepository(dDB *dynamodb.DynamoDB, dDBS *dynamodbstreams.DynamoDBStreams) CRUD {
	return &repo{
		svc:  dDB,
		svc2: dDBS,
	}
}

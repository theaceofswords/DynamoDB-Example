package repository

import (
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/jinzhu/gorm"
)

var nextIterator string

type CRUD interface {
	AddRecord(movie models.Movie)
	GetRecord(movieName string, movieYear string) (models.Movie, error)
	UpdateRecord(movie models.Movie)
	DeleteRecord(movieName string, movieYear string) error
	InitaliseData()
	CreateTable()
	GetIterator() string
	psqlWrite(string)
	psqlRead(string)
	iteratorExpCheck(string) bool
}

type repo struct {
	svc  *dynamodb.DynamoDB
	svc2 *dynamodbstreams.DynamoDBStreams
	psqlDB *gorm.DB
}

func CreateRepository(dDB *dynamodb.DynamoDB, dDBS *dynamodbstreams.DynamoDBStreams,pDB *gorm.DB) CRUD {
	return &repo{
		svc:  dDB,
		svc2: dDBS,
		psqlDB: pDB,
	}
}

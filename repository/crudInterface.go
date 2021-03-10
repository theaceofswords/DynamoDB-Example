package repository

import (
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type CRUD interface {
	AddRecord(movie models.Movie)
	GetRecord(movieName string, movieYear string) (models.Movie, error)
	UpdateRecord(movie models.Movie)
	DeleteRecord(movieName string, movieYear string) error
	InitaliseData()
}

type repo struct {
	svc *dynamodb.DynamoDB
}

func CreateRepository(dDB *dynamodb.DynamoDB) CRUD {
	return &repo{
		svc: dDB,
	}
}

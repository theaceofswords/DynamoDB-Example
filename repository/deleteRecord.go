package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (r *repo) DeleteRecord(movieName string, movieYear string) error {

	var shardIterator string
    if nextIterator == ""{
		fmt.Println("new iterator")
		shardIterator = r.GetIterator()
	}else{
		
		shardIterator = nextIterator
	}

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

	_, err := r.svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Error Deleting Item")
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Deleted")
	r.psqlWrite(shardIterator)
	return nil
}

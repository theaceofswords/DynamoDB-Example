package repository

import(
	"fmt"
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"encoding/json"
	

	// "errors"
	// "strings"
)

func (r* repo) psqlWrite(shardIterator string){
	// var shardIterator string
    // if nextIterator == ""{
	// 	shardIterator = r.GetIterator()
	// }else{
	// 	shardIterator = nextIterator
	// }

	recInput := &dynamodbstreams.GetRecordsInput{
		ShardIterator: aws.String(shardIterator),
		
	}

   // var expItrErr = errors.New("ExpiredIteratorException:")

	
	result, err := r.svc2.GetRecords(recInput)
	if err != nil {
		fmt.Println(err.Error())
	}

	// if err != nil {
	// 	if  strings.Compare(err.Error(), expItrErr.Error()) == 0 {
	// 		shardIterator = r.GetIterator()
	// 	}else{
	// 		fmt.Println(err.Error(),"=",expItrErr.Error())
	// 	}
	// }

	// if err != nil{
	// 	if aerr, ok := err.(awserr.Error); ok {
	// 		switch aerr.Code() {
	// 		case dynamodbstreams.ErrCodeExpiredIteratorException:
	// 			shardIterator = r.GetIterator()
	// 			recInput = &dynamodbstreams.GetRecordsInput{
	// 				ShardIterator: aws.String(shardIterator),
					
	// 			}
	// 			result, _ = r.svc2.GetRecords(recInput)
	// 		default:
	// 			fmt.Println(aerr.Error())
	// 		}
	// 	}else{
	// 		fmt.Println(err.Error())
	// 	}
	// }
	nextIterator = *result.NextShardIterator

	fmt.Println(result)
	var strc models.Records

    for _,record := range result.Records{
		// strc = models.Records{
		// 	OldImage: fmt.Sprintf("%v",record.Dynamodb.OldImage),
		// 	NewImage: fmt.Sprintf("%v",record.Dynamodb.NewImage),
		// 	EventId: *record.EventID,
		// 	EventName: *record.EventName,
		// }

		var newImage models.Movie
        err = dynamodbattribute.UnmarshalMap(record.Dynamodb.NewImage, &newImage)
        if err != nil {
            fmt.Println(err.Error())
        }
		
		
		var oldImage models.Movie
        err = dynamodbattribute.UnmarshalMap(record.Dynamodb.OldImage, &oldImage)
        if err != nil {
            fmt.Println(err.Error())
        }
        str1, err := json.Marshal(oldImage)
		str2, err := json.Marshal(newImage)
	
		
		strc = models.Records{
			OldImage: str1,
		 	NewImage: str2,
			EventId: *record.EventID,
		 	EventName: *record.EventName,
		}

		fmt.Println(strc)

		
		err = r.psqlDB.Create(&strc).Error
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	
	r.psqlRead(strc.EventId)

}

func (r *repo) psqlRead(eventId string){
	var rec models.Records
	err := r.psqlDB.Where("event_id=?", eventId).Find(&rec).Error
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(rec)
}

func (r *repo)iteratorExpCheck(shardIterator string) bool {
	recInput := &dynamodbstreams.GetRecordsInput{
		ShardIterator: aws.String(shardIterator),
		
	}
	_, err := r.svc2.GetRecords(recInput)
	if err != nil{
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodbstreams.ErrCodeExpiredIteratorException:
				return true
			default:
				fmt.Println(aerr.Error())
				return true
			}
		}else{
			fmt.Println(err.Error())
			return false
		}
	}
	return false

}
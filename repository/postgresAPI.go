package repository

import(
	"fmt"
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
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

	result, err := r.svc2.GetRecords(recInput)
	if err != nil {
		fmt.Println(err.Error())
	}
	nextIterator = *result.NextShardIterator

	fmt.Println(result)
	strc := models.Records{
		OldImage: fmt.Sprintf("%v",result.Records[0].Dynamodb.OldImage),
		NewImage: fmt.Sprintf("%v",result.Records[0].Dynamodb.NewImage),
		EventId: *result.Records[0].EventID,
		EventName: *result.Records[0].EventName,
	}


	err = r.psqlDB.Create(&strc).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
package lambda



import (
    "fmt"
	"context"

    "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    //"encoding/json"
)

type Records struct {
	OldImage string  `json:"OldImage"`
	NewImage string `json:"NewImage"`
	EventId string   `json:"EventId"`
	EventName string  `json:"EventName"`
	s_no int `gorm:"primaryKey"`

}

// type Movie struct {
// 	Year      int   `json:"Year"`
// 	Title     string `json:"Title"`
// 	Category  string  `json:"Category"`
// 	Plot      string  `json:"Plot"`
// 	Rating    float64 `json:"Rating"`
// 	Director  string  `json:"Director"`
// 	LeadActor string  `json:"LeadActor"`
// 	Duration  int   `json:"Duration"`
// }

func handleRequest(ctx context.Context, e events.DynamoDBEvent) {
    var records Records
    // var err error
    for _, record := range e.Records {
       // fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)
        
        // var newImage Movie
         err = dynamodbattribute.UnmarshalMap(record.Change.NewImage, &newImage)
        // if err != nil {
        //     fmt.Println(err.Error())
        // }
		
		
		// var oldImage Movie
         err = dynamodbattribute.UnmarshalMap(record.Change.OldImage, &oldImage)
        // if err != nil {
        //     fmt.Println(err.Error())
        // }
        // str1, err := json.Marshal(oldImage)
		// str2, err := json.Marshal(newImage)
        
        
        
        records = Records{
			OldImage: fmt.Sprintf("%v",record.Change.OldImage),
		 	NewImage: fmt.Sprintf("%v",record.Change.NewImage),
			EventId: record.EventID,
		 	EventName: record.EventName,
		}

        // Print new values for attributes of type String
        // for name, value := range record.Change.NewImage {
        //     if value.DataType() == events.DataTypeString {
        //         fmt.Printf("Attribute name: %s, value: %s\n", name, value.String())
        //     }
        // }

        fmt.Println(records)
    }
}

func main(){
	lambda.Start(handleRequest)
}
 package main

// import (
//   "os"
//   "log"
//   "context"
//   "encoding/json"
//   runtime "github.com/aws/aws-lambda-go/lambda"
//   "github.com/aws/aws-lambda-go/events"
//   "github.com/aws/aws-lambda-go/lambdacontext"
//   "github.com/aws/aws-sdk-go/aws/session"
//   "github.com/aws/aws-sdk-go/service/lambda"
//   "github.com/aws/aws-sdk-go/aws"
//   "code.qburst.com/navaneeth.k/DynamoDB-example/models"
//   "github.com/aws/aws-sdk-go/aws/credentials"
//   "fmt"
//   "gopkg.in/yaml.v2"
// )



// // var client = lambda.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")})
// var client = getClient()


// func init() {
    
//   callLambda()
// }

// func getClient() *lambda.Lambda{
//     f, err := os.Open("./config/cred.yml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer f.Close()

// 	var crd models.Cred
// 	decoder := yaml.NewDecoder(f)
// 	err = decoder.Decode(&crd)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	sess, err := session.NewSession(&aws.Config{
// 		Region:      aws.String(crd.Region),
// 		Credentials: credentials.NewStaticCredentials(crd.AWS_ACCESS_KEY_ID, crd.AWS_SECRET_ACCESS_KEY, ""),
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
//     client := lambda.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")})
//     return client
 
// }

// func callLambda() (string, error) {
//   input := &lambda.GetAccountSettingsInput{}
//   req, resp := client.GetAccountSettingsRequest(input)
//   err := req.Send()
//   output, _ := json.Marshal(resp.AccountUsage)
//   return string(output), err
// }

// func handleRequest(ctx context.Context, event events.SQSEvent) (string, error) {
//   // event
//   eventJson, _ := json.MarshalIndent(event, "", "  ")
//   log.Printf("EVENT: %s", eventJson)
//   // environment variables
//   log.Printf("REGION: %s", os.Getenv("AWS_REGION"))
//   log.Println("ALL ENV VARS:")
//   for _, element := range os.Environ() {
//     log.Println(element)
//   }
//   // request context
//   lc, _ := lambdacontext.FromContext(ctx)
//   log.Printf("REQUEST ID: %s", lc.AwsRequestID)
//   // global variable
//   log.Printf("FUNCTION NAME: %s", lambdacontext.FunctionName)
//   // context method
//   deadline, _ := ctx.Deadline()
//   log.Printf("DEADLINE: %s", deadline)
//   // AWS SDK call
//   usage, err := callLambda()
//   if err != nil {
//     return "ERROR", err
//   }
//   return usage, nil
// }

// func main() {
//   runtime.Start(handleRequest)
// }





import (
	"fmt"
	
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "user123"
	dbname   = "datb3"
)

var (
	db  *gorm.DB
	err error
)
type Records struct {
	OldImage string  `json:"OldImage"`
	NewImage string `json:"NewImage"`
	EventId string   `json:"EventId"`
	EventName string  `json:"EventName"`
	s_no int `gorm:"primaryKey"`

}




func LambdaHandler() {
	fmt.Println("start")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("connected")
	strc := Records{
		OldImage: "test",
		 NewImage: "test",
		EventId: "test",
		 EventName: "test",
	}
	fmt.Println(strc)

	
	
	err = db.Create(&strc).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}


func main ()  {
	lambda.Start(LambdaHandler)
}
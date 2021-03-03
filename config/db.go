package config

import (
	"fmt"
	"os"

	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gopkg.in/yaml.v2"
)

func Connect() *dynamodb.DynamoDB {
	// Initialize a session
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))
	f, err := os.Open("./config/cred.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	var crd models.Cred
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&crd)
	if err != nil {
		fmt.Println(err.Error())
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(crd.Region),
		Credentials: credentials.NewStaticCredentials(crd.AWS_ACCESS_KEY_ID, crd.AWS_SECRET_ACCESS_KEY, ""),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Create DynamoDB client
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")})

	return svc
}

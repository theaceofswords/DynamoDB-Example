package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(){
     fmt.Println("Triggered")
}

func main ()  {
	lambda.Start(handler)
}
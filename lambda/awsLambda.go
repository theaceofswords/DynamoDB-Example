package lambda

import (
	"fmt"
	
	

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
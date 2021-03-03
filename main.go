package main

import (
	"code.qburst.com/navaneeth.k/DynamoDB-example/repository"
)

func main() {

	// repository.CreateTable()
	// repository.InitaliseData()

	year := "2012"
	movieTitle := "movie 2"
	repository.GetRecord(movieTitle, year)

	//repository.DeleteRecord()
	//repository.GetRecord()
}

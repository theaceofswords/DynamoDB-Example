package main

import (
	"code.qburst.com/navaneeth.k/DynamoDB-example/controller"
)

func main() {

	//repository.CreateTable()
	//repository.InitaliseData()

	// year := "2012"
	// movieTitle := "movie 2"
	// repository.GetRecord(movieTitle, year)

	// movie := models.Movie{
	// 	Year:      2012,
	// 	Title:     "movie 2",
	// 	Category:  "action",
	// 	Plot:      "lots of fighting",
	// 	Rating:    4,
	// 	Director:  "so and so",
	// 	LeadActor: "the other guy",
	// 	Duration:  193,
	// }
	// repository.UpdateRecord(movie)

	//repository.DeleteRecord()
	//repository.GetRecord(movieTitle, year)

	controller.RequestHandler()
}

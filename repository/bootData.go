package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
)

func InitaliseData() {

	raw, err := ioutil.ReadFile("./models/raw.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var movies []models.Movie

	json.Unmarshal(raw, &movies)

	for _, movie := range movies {
		AddRecord(movie)
	}

	// 	newMovie := models.Movie{Year: 2015,
	// 		Title:  "The New Movie",
	// 		Plot:   "something happens.",
	// 		Rating: 0.0,
	// 	}
	// 	repository.AddRecord(newMovie)

	// 	newMovie = models.Movie{
	// 		Year:      2012,
	// 		Title:     "movie 2",
	// 		Plot:      "some plot",
	// 		Rating:    5,
	// 		Director:  "director name",
	// 		LeadActor: "actor name",
	// 		Duration:  157,
	// 	}
	// 	repository.AddRecord(newMovie)

	// 	newMovie = models.Movie{
	// 		Year:     2012,
	// 		Title:    "movie 3",
	// 		Plot:     "some other plot",
	// 		Rating:   3,
	// 		Director: "director name2",
	// 		Duration: 176,
	// 	}
	// 	repository.AddRecord(newMovie)
}

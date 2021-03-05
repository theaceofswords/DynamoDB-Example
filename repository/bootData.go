package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
)

func (r *repo) InitaliseData() {

	raw, err := ioutil.ReadFile("./models/raw.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var movies []models.Movie

	json.Unmarshal(raw, &movies)

	for _, movie := range movies {
		r.AddRecord(movie)
	}

}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"code.qburst.com/navaneeth.k/DynamoDB-example/repository"
)

func RequestHandler() {
	http.HandleFunc("/movies", requestHandler)
	fmt.Println("Running,.. ")
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		title := r.URL.Query().Get("title")
		year := r.URL.Query().Get("year")
		//fmt.Println(title, year)

		movie, err := repository.GetRecord(title, year)
		if err != nil {
			msg := models.ErrorMsg{err.Error(), http.StatusNotFound, "Record does not exist"}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(msg)
		} else {
			json.NewEncoder(w).Encode(movie)
		}

	case "POST":
		var movie models.Movie

		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			msg := models.ErrorMsg{err.Error(), http.StatusUnprocessableEntity, "Invalid Body type"}
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(msg)
		} else {
			repository.AddRecord(movie)
			msg := models.ErrorMsg{"nil", http.StatusOK, "Record added"}
			json.NewEncoder(w).Encode(msg)
		}
	case "PUT":
		var movie models.Movie

		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			msg := models.ErrorMsg{err.Error(), http.StatusUnprocessableEntity, "Invalid Body type"}
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(msg)
		} else {
			repository.UpdateRecord(movie)
			msg := models.ErrorMsg{"nil", http.StatusOK, "Record Updated"}
			json.NewEncoder(w).Encode(msg)
		}
	case "DELETE":
		title := r.URL.Query().Get("title")
		year := r.URL.Query().Get("year")
		//fmt.Println(title, year)

		err := repository.DeleteRecord(title, year)
		if err != nil {
			msg := models.ErrorMsg{err.Error(), http.StatusNotFound, "Record does not exist"}
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := models.ErrorMsg{"nil", http.StatusOK, "Record Deleted"}
			json.NewEncoder(w).Encode(msg)
		}

	}

}

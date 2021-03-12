package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"code.qburst.com/navaneeth.k/DynamoDB-example/config"
	"code.qburst.com/navaneeth.k/DynamoDB-example/models"
	"code.qburst.com/navaneeth.k/DynamoDB-example/repository"
)

type handler struct {
	crud repository.CRUD
}

func RequestHandler() {
	svc, svc2 := config.Connect()
	psqlDB := config.PsqlConnect()
	defer psqlDB.Close()
	crud := repository.CreateRepository(svc, svc2,psqlDB)
	
	//crud.CreateTable()
	//crud.InitaliseData()
	t := handler{crud}
	http.HandleFunc("/movies", t.requestHandler)
	fmt.Println("Running,.. ")
	http.ListenAndServe(":8080", nil)
}

func (t *handler) requestHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		title := r.URL.Query().Get("title")
		year := r.URL.Query().Get("year")
		//fmt.Println(title, year)

		movie, err := t.crud.GetRecord(title, year)
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
			t.crud.AddRecord(movie)
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
			t.crud.UpdateRecord(movie)
			msg := models.ErrorMsg{"nil", http.StatusOK, "Record Updated"}
			json.NewEncoder(w).Encode(msg)
		}
	case "DELETE":
		title := r.URL.Query().Get("title")
		year := r.URL.Query().Get("year")
		//fmt.Println(title, year)

		err := t.crud.DeleteRecord(title, year)
		if err != nil {
			msg := models.ErrorMsg{err.Error(), http.StatusNotFound, "Record does not exist"}
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := models.ErrorMsg{"nil", http.StatusOK, "Record Deleted"}
			json.NewEncoder(w).Encode(msg)
		}

	}

}

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Router()  {
	router := mux.NewRouter()
	router.HandleFunc("/home",ABC).Methods("GET")
}

func ABC(w http.ResponseWriter,r *http.Request) {

	var a string
	a  = "Hello World"
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(a)

}
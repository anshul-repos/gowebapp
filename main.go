package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Hello")
	})
	fmt.Println("Starting server at 9090\n")
	if err := http.ListenAndServe(":9090",nil) ; err != nil {
		log.Fatal(err)
	}
	Router()
}
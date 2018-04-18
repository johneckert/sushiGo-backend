package main

//uncomment this --
import (
	"log"      //error logging
	"net/http" //random number gen
	//convert to string

	"github.com/gorilla/mux" //router
)

func main() {
	//initialize router
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi/", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

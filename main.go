package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammad-siraj/crud_api/operations"
)

func main() {
	fmt.Printf("Hello world")
	m := mux.NewRouter()
	m.HandleFunc("/cars/{id}", operations.Getcar).Methods("GET")
	m.HandleFunc("/cars/{model}/{make}/{year}", operations.Createcar).Methods("POST")
	m.HandleFunc("/cars/{id}/{model}/{make}/{year}", operations.Updatecar).Methods("PUT")
	m.HandleFunc("/cars/{id}", operations.Deletecar).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8500", m))
}

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammad-siraj/crud_api/database"
	"github.com/mohammad-siraj/crud_api/entities"
)

func Getcar(w http.ResponseWriter, req *http.Request) {
	var c entities.Car
	database.Getdata_db(&c, req)
	//vars := mux.Vars(req)["model"]
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(c)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//fmt.Printf("\n %s \n", vars)
}

func Createcar(w http.ResponseWriter, req *http.Request) {
	var c entities.Car
	c.Model = mux.Vars(req)["model"]
	c.Make = mux.Vars(req)["make"]
	c.Year = mux.Vars(req)["year"]
	database.Postdata_db(&c, req)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(c)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//fmt.Printf("\n %s \n", vars)
}

func Updatecar(w http.ResponseWriter, req *http.Request) {
	var c entities.Car
	c.Model = mux.Vars(req)["model"]
	c.Make = mux.Vars(req)["make"]
	c.Year = mux.Vars(req)["year"]
	database.Updatedata_db(&c, req)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(c)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//fmt.Printf("\n %s \n", vars)
}

func Deletecar(w http.ResponseWriter, req *http.Request) {
	var c entities.Car
	database.Deletedata_db(&c, req)
}

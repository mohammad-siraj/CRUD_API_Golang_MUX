package operations

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohammad-siraj/crud_api/database"
	"github.com/mohammad-siraj/crud_api/entities"
)

var db database.Database

func Init() error {
	var err error
	db, err = database.NewDatabase()
	return err
}

func Getcar(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	carid, err := strconv.Atoi(id)
	if err != nil {

	}
	c, err := db.Getdata_db(carid)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func Createcar(w http.ResponseWriter, req *http.Request) {
	car := entities.Car{Model: mux.Vars(req)["model"], Make: mux.Vars(req)["make"], Year: mux.Vars(req)["year"]}
	err := db.Postdata_db(car)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(car)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func Updatecar(w http.ResponseWriter, req *http.Request) {
	car := entities.Car{Model: mux.Vars(req)["model"], Make: mux.Vars(req)["make"], Year: mux.Vars(req)["year"]}
	id, err := strconv.Atoi((mux.Vars(req)["id"]))
	car, err = db.Updatedata_db(car, id)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(car)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func Deletecar(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi((mux.Vars(req)["id"]))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Deletedata_db(id)
}

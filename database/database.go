package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/mohammad-siraj/crud_api/entities"
)

const (
	host     = "localhost"
	port     = "8600"
	user     = "root"
	password = "admin"
	db_name  = "root"
)

func db_intializer() *sql.DB {
	connStr := "host=" + host + " port=" + port + " dbname=" + db_name + " user=" + user + " password=" + password + " sslmode=disable"
	fmt.Printf(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("\n DATABASE CONNECTION SUCESSFULL=1")
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("it went wrong here \n\n\n")
		panic(err)
	}
	fmt.Printf("DATABASE CONNECTION SUCESSFULL")
	return db
}

func Getdata_db(c *entities.Car, req *http.Request) {
	db := db_intializer()
	defer db.Close()
	//var id intn
	var model, make, year string
	con := mux.Vars(req)["id"]
	//log.Fatal(err)
	q := "SELECT model,make,year FROM car WHERE id=" + con
	row, err := db.Query(q)
	fmt.Printf("\nthe query is %s \n", q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if row == nil {
		fmt.Printf("the record in not present")
	}
	//fmt.Printf("\nthe query is %s \n", q)
	for row.Next() {
		err = row.Scan(&model, &make, &year)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("\nread sucessful %s \n", model)
	}
	c.Make = make
	c.Model = model
	c.Year = year
}

func Postdata_db(c *entities.Car, req *http.Request) {
	db := db_intializer()
	defer db.Close()
	//con := mux.Vars(req)["id"]
	//log.Fatal(err)
	q := "INSERT INTO car (id,model,make,year) VALUES ((SELECT MAX(id) FROM car)+1,'" + c.Model + "','" + c.Make + "','" + c.Year + "')"
	row, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if row == nil {
		fmt.Printf("the record in not present")
	}
	fmt.Printf("\nthe query is %s \n", q)
}

func Updatedata_db(c *entities.Car, req *http.Request) {
	db := db_intializer()
	defer db.Close()
	con := mux.Vars(req)["id"]
	//log.Fatal(err)
	q := "UPDATE car SET model='" + c.Model + "',make='" + c.Make + "',year='" + c.Year + "' WHERE id=" + con
	row, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if row == nil {
		fmt.Printf("the record in not present")
	}
	//fmt.Printf("\nthe query is %s \n", q)
}

func Deletedata_db(c *entities.Car, req *http.Request) {
	db := db_intializer()
	defer db.Close()
	con := mux.Vars(req)["id"]
	//log.Fatal(err)
	q := "DELETE FROM car WHERE id=" + con
	row, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if row == nil {
		fmt.Printf("the record in not present")
	}
	//fmt.Printf("\nthe query is %s \n", q)
}

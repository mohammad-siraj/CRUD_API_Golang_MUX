package database

import (
	"database/sql"
	"fmt"
	"log"

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

type database struct {
	conn *sql.DB
}

type Database interface {
	Getdata_db(id int) (entities.Car, error)
	Postdata_db(c entities.Car) error
	Updatedata_db(car entities.Car, id int) (entities.Car, error)
	Deletedata_db(id int) error
}

func NewDatabase() (*database, error) {
	connStr := "host=" + host + " port=" + port + " dbname=" + db_name + " user=" + user + " password=" + password + " sslmode=disable"
	fmt.Printf(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("\n DATABASE CONNECTION SUCESSFULL=1")
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()

	return &database{conn: db}, err
}

func (db *database) Getdata_db(id int) (entities.Car, error) {

	var model, make, year string
	q := "SELECT model,make,year FROM car WHERE id=$1"
	err := db.conn.QueryRow(q, id).Scan(&model, &make, &year)
	if err != nil {
		log.Fatal(err)
	}
	car := entities.Car{
		Model: model,
		Make:  make,
		Year:  year,
	}
	return car, err
}

func (db *database) Postdata_db(c entities.Car) error {
	q := "INSERT INTO car (id,model,make,year) VALUES ((SELECT MAX(id) FROM car)+1,'" + c.Model + "','" + c.Make + "','" + c.Year + "')"
	row, err := db.conn.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return err
	//fmt.Printf("\nthe query is %s \n", q)
}

func (db *database) Updatedata_db(car entities.Car, id int) (entities.Car, error) {
	q := "UPDATE car SET model='" + car.Model + "',make='" + car.Make + "',year='" + car.Year + "' WHERE id=$1"
	row, err := db.conn.Query(q, id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return car, err
}

func (db *database) Deletedata_db(id int) error {
	q := "DELETE FROM car WHERE id=$1"
	row, err := db.conn.Query(q, id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return err
}

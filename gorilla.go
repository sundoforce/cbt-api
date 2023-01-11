package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{table}/{id}", getRecord).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "host=localhost user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		fmt.Fprintln(w, "Error connecting to the database:", err)
		return
	}
	defer db.Close()

	vars := mux.Vars(r)
	tableName := vars["table"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, "Invalid record ID")
		return
	}

	var record interface{}
	switch tableName {
	case "apple":
		record = &Apple{}
	case "google":
		record = &Google{}
	default:
		fmt.Fprintln(w, "Invalid table name")
		return
	}

	if err := db.Table(tableName).First(record, id).Error; err != nil {
		fmt.Fprintln(w, "Error fetching record:", err)
		return
	}

	// TODO: print out the record
}

type Apple struct {
	gorm.Model
	Name string
}

type Google struct {
	gorm.Model
    Name string
}

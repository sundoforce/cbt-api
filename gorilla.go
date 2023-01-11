package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
    if err != nil {
    	log.Fatal(err)
    }
    defer db.Close()

    r := mux.NewRouter()

	r.HandleFunc("/api/{table}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tableName := vars["table"]

		rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer rows.Close()

		var data []map[string]interface{}
		for rows.Next() {
			columns, _ := rows.Columns()
			count := len(columns)
			values := make([]interface{}, count)
			valuePtrs := make([]interface{}, count)
			for i := 0; i < count; i++ {
				valuePtrs[i] = &values[i]
			}
			rows.Scan(valuePtrs...)

			entry := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]
				b, ok := val.([]byte)
				if ok {
					v = string(b)
				} else {
					v = val
				}
				entry[col] = v
			}
			data = append(data, entry)
		}

		json.NewEncoder(w).Encode(data)
	}).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", r))
}



package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

func main() {
    err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file")
    }
   db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
       os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DBNAME")))

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := gin.Default()

    router.GET("/api/:table", func(c *gin.Context) {
        tableName := c.Param("table")

        rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

        c.JSON(http.StatusOK, gin.H{"data": data})
    })

    router.Run(":8081")
}


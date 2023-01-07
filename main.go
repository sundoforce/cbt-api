package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Quiz struct {
	ID               int
	Title            string
	PassageA         string
	PassageB         string
	PassageC         string
	PassageD         string
	PassageE         string
	PassageF         string
	Answer           string
	AnswerCandidates string
	Description      string
}

type Quiz2 struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	PassageA         string `json:"passage_a"`
	PassageB         string `json:"passage_b"`
	PassageC         string `json:"passage_c"`
	PassageD         string `json:"passage_d"`
	PassageE         string `json:"passage_e"`
	PassageF         string `json:"passage_f"`
	Answer           string `json:"answer"`
	AnswerCandidates string `json:"answer_candidates"`
	Description      string `json:"description"`
}

func (q *Quiz2) GetQuiz(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM quiz WHERE id=$1", q.ID).Scan(
		&q.ID, &q.Title, &q.PassageA, &q.PassageB, &q.PassageC, &q.PassageD, &q.PassageE, &q.PassageF, &q.Answer, &q.AnswerCandidates, &q.Description,
	)
}

func init() {
    // Load the environment variables from the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    // Retrieve the database connection details from environment variables
    host := os.Getenv("POSTGRES_HOST")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DBNAME")

    // Build the connection string
    connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

    // Connect to the database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Check the connection status
    err = db.Ping()
    if err != nil {
        log.Println("Error: Could not establish a connection to the database")
    } else {
        log.Println("Success: Connected to the database")
    }
}

func GetQuizzes(c *gin.Context) {
    var quizzes []Quiz
    if err := db.Find(&quizzes).Error; err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        fmt.Println(err)
    } else {
        c.JSON(http.StatusOK, quizzes)
    }
}

func main() {
	fmt.Printf("hello, world\n")
    v1 := router.Group("/api")
    {
        v1.GET("/quiz", GetQuizzes)
    }
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Start the server on port 8080
	router.Run(":8080")
}


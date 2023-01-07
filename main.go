package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/joho/godotenv"
)

type Quiz struct {
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

var db *gorm.DB

func GetQuizzes(c *gin.Context) {
    var quizzes []Quiz
    if err := db.Find(&quizzes).Error; err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        fmt.Println(err)
    } else {
        c.JSON(http.StatusOK, quizzes)
    }
}

func GetQuiz(c *gin.Context) {
    var quiz Quiz
    quizID := c.Params.ByName("id")
    if err := db.Where("id = ?", quizID).First(&quiz).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        fmt.Println(err)
    } else {
        c.JSON(http.StatusOK, quiz)
    }
}

func main() {
    fmt.Printf("hello, world\n")
    router := gin.Default()

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
    db, err = gorm.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Set some gorm options
    db.LogMode(true)
    db.SingularTable(true)

    router.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })

    v1 := router.Group("/api")
    {
        v1.GET("/quiz", GetQuizzes)
        v1.GET("/quiz/:id", GetQuiz)
    }

	// Start the server on port 8080
	router.Run(":8080")
}


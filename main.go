package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
    "database/sql"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
//	"github.com/rs/cors"
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

type QuizManagement struct {
    ID               int    `json:"id"`
    Title            string `json:"title"`
    Description      string `json:"description"`
    Version          string `json:"version"`

}

type User struct {
    gorm.Model
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

type DB struct {
    *sql.DB
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

func GetQuizList(c *gin.Context) {
    var quizzes []QuizManagement
    if err := db.Find(&quizzes).Error; err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        fmt.Println(err)
    } else {
        c.JSON(http.StatusOK, quizzes)
    }
}



func GetTable(c *gin.Context) {
//    tableID := c.Params.ByName("table")
//    start := c.Params.ByName("start")
//    end := c.Params.ByName("end")

   // Execute the raw SQL query.
   query := fmt.Sprintf("SELECT * FROM quiz ")
	var quizzes []Quiz
	db.Raw(query)

	// Return the results as JSON.
    c.JSON(http.StatusOK, quizzes)

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

func GetRandomQuizze(c *gin.Context) {
	start, err := strconv.Atoi(c.Params.ByName("start"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	end, err := strconv.Atoi(c.Params.ByName("end"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	randomIndex := rand.Intn(end-start+1) + start
	var quiz Quiz
	if err := db.Where("id = ?", randomIndex).First(&quiz).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, quiz)
}

func GetRandomQuizzes(c *gin.Context) {
    // SELECT * FROM quiz WHERE id >= 1 AND id <= 10 ORDER BY RANDOM() LIMIT 10;
	start, err := strconv.Atoi(c.Params.ByName("start"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	end, err := strconv.Atoi(c.Params.ByName("end"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var quizzes []Quiz
	if err := db.Where("id BETWEEN ? AND ?", start, end).Order("random()").Find(&quizzes).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, quizzes)
}

func GetTableList(c *gin.Context)  {
    // GET /table/:name
        // Get the table name from the URL parameter
    tableName := c.Params.ByName("name")

        // Execute a SELECT query on the table
        var users []User
        if err := db.Table(tableName).Find(&users).Error; err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error finding users in table %s: %s", tableName, err))
            return
        }

        // Return the results
        c.JSON(http.StatusOK, users)

}

func connectDB() *gorm.DB {
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
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Set some gorm options
	db.LogMode(true)
	db.SingularTable(true)
	return db
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


func main() {
	fmt.Printf("hello, world\n")
	router := gin.Default()
    gin.SetMode(gin.ReleaseMode)

    router.Use(CORSMiddleware())


	err := godotenv.Load()
//    if err != nil {
//        log.Fatal("Error loading .env file", err)
//    }
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
	//    db := connectDB()
	//    defer db.Close()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	v1 := router.Group("/api")
	{
		v1.GET("/quiz", GetQuizzes)
		v1.GET("/quiz/:id", GetQuiz)
		v1.GET("/quiz/randoms/:start/:end", GetRandomQuizzes)
		v1.GET("/quiz/random/:start/:end", GetRandomQuizze)
        v1.GET("/table/:table/:start/:end", GetTable)
        v1.GET("/quizs", GetQuizList)
        v1.GET("/table2/:name", GetTableList)
	}

	// Start the server on port 8080
	router.Run(":8081")
}

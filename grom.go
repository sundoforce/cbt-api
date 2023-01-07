package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Quiz struct {
	ID int json:"id"
	Title string json:"title"
	PassageA string json:"passage_a"
	PassageB string json:"passage_b"
	PassageC string json:"passage_c"
	PassageD string json:"passage_d"
	PassageE string json:"passage_e"
	PassageF string json:"passage_f"
	Answer string json:"answer"
	AnswerCandidates string json:"answer_candidates"
	Description string json:"description"
}

func init() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Quiz{})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.GET("/quiz", GetQuizzes)
		v1.GET("/quiz/:id", GetQuiz)
		v1.POST("/quiz", CreateQuiz)
		v1.PUT("/quiz/:id", UpdateQuiz)
		v1.DELETE("/quiz/:id", DeleteQuiz)
	}

	router.Run()
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

func GetQuiz(c *gin.Context) {
	var quiz Quiz
	quizID := c.Params.ByName("id")
	if err := db.Where("id = ?", quizID).First(&quiz).Error; err != nil {
		c.AbortWith

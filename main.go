package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
)

type Quiz struct {
	ID            int
	PassageA      string
	PassageB      string
	PassageC      string
	PassageD      string
	PassageE      string
	PassageF      string
	Answer        string
	AnswerChoices string
}

func main() {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
	})
	defer db.Close()

	router := gin.Default()

	router.GET("/api/quiz", func(c *gin.Context) {
		var quizzes []Quiz
		err := db.Model(&quizzes).Select()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, quizzes)
		}
	})

	router.GET("/api/quiz/:id", func(c *gin.Context) {
		var quiz Quiz
		quizID := c.Params.ByName("id")
		err := db.Model(&quiz).Where("id = ?", quizID).Select()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, quiz)
		}
	})

	err := router.Run(":" + os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}

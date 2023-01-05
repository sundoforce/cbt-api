package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
)

type Quiz struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	PassageA      string `json:"passage_a"`
	PassageB      string `json:"passage_b"`
	PassageC      string `json:"passage_c"`
	PassageD      string `json:"passage_d"`
	PassageE      string `json:"passage_e"`
	PassageF      string `json:"passage_f"`
	Answer        string `json:"answer"`
	AnswerChoices string `json:"answer_choices"`
}

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	})
	defer db.Close()

	router.GET("/api/quiz", func(c *gin.Context) {
		var quiz []Quiz
		err := db.Model(&quiz).Select()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, quiz)
	})

	router.Run()
}

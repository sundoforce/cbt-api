package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
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

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "mypassword",
		Database: "mydatabase",
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

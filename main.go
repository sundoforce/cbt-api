package main

import (
	"fmt"
	"net/http"
	"os"

	//	"log"
	//	"os"
	"github.com/gin-gonic/gin"
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
	fmt.Printf("hello, world\n")

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
	//
	//	router.GET("/api/quiz/:id", func(c *gin.Context) {
	//		var quiz Quiz
	//		quizID := c.Params.ByName("id")
	//		err := db.Model(&quiz).Where("id = ?", quizID).Select()
	//		if err != nil {
	//			c.AbortWithStatus(http.StatusInternalServerError)
	//			fmt.Println(err)
	//		} else {
	//			c.JSON(http.StatusOK, quiz)
	//		}
	//	})
	//
	//	err := router.Run(":" + os.Getenv("API_PORT"))
	//	if err != nil {
	//		log.Fatalln(err)
	//	}

	//	router.POST("/api/:table/:id", func(c *gin.Context) {
	//		var quiz Quiz
	//		table := c.Params.ByName("table")
	//		id := c.Params.ByName("id")
	//		if err := c.BindJSON(&quiz); err != nil {
	//			c.AbortWithStatus(http.StatusBadRequest)
	//			fmt.Println(err)
	//			return
	//		}
	//		_, err := db.Model(&quiz).Where("id = ?", id).Update()
	//		if err != nil {
	//			c.AbortWithStatus(http.StatusInternalServerError)
	//			fmt.Println(err)
	//		} else {
	//			c.JSON(http.StatusOK, gin.H{"id #" + id: "updated"})
	//		}
	//	})
	//
	//	router.GET("/api/:table/random/:start/:end", func(c *gin.Context) {
	//		tableName := c.Params.ByName("table")
	//		start, _ := strconv.Atoi(c.Params.ByName("start"))
	//		end, _ := strconv.Atoi(c.Params.ByName("end"))
	//
	//		// 입력받은 tableName과 column 이름을 이용해서 SELECT 쿼리 생성
	//		query := fmt.Sprintf("SELECT * FROM %s WHERE id BETWEEN %d AND %d ORDER BY random()", tableName, start, end)
	//
	//		// db.Query() 함수를 이용해서 쿼리 수행
	//		_, err := db.Query(query)
	//		if err != nil {
	//			c.AbortWithStatus(http.StatusInternalServerError)
	//			fmt.Println(err)
	//		} else {
	//			c.JSON(http.StatusOK, gin.H{"message": "success"})
	//		}
	//	})
	//
	//	router.GET("/api/:table/:start/:end", func(c *gin.Context) {
	//		tableName := c.Params.ByName("table")
	//		start, _ := strconv.Atoi(c.Params.ByName("start"))
	//		end, _ := strconv.Atoi(c.Params.ByName("end"))
	//
	//		// Create a slice of structs with the same type as the table
	//		var records []struct{}
	//
	//		// Use ORM to execute the SELECT query
	//		err := db.Model(&records).Where("id BETWEEN ? AND ?", start, end).Order("random()").Select()
	//		if err != nil {
	//			c.AbortWithStatus(http.StatusInternalServerError)
	//			fmt.Println(err)
	//		} else {
	//			c.JSON(http.StatusOK, gin.H{"records": records})
	//		}
	//	})
}

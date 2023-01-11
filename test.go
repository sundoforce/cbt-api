package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

type Quiz struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	PassageA          string `json:"passage_a"`
	PassageB          string `json:"passage_b"`
	PassageC          string `json:"passage_c"`
	PassageD          string `json:"passage_d"`
	PassageE          string `json:"passage_e"`
	PassageF          string `json:"passage_f"`
	Answer            string `json:"answer"`
	AnswerCandidates  string `json:"answer_candidates"`
	Description       string `json:"description"`
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt


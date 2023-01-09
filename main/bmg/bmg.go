package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"pcrTest/auth"
)

type BodyRequest struct {
	Iin         string `json:"iin"`
	PhoneNumber string `json:"phone_number"`
}

type Result struct {
	Iin    string
	Status string
}

func main() {

	const connStr = "postgres://admin:admin@localhost:5432/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	token := auth.GetToken

	stringToken := token()

	query := "select iin , phone_number from bmg_request"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		var request BodyRequest
		err := rows.Scan(&request.Iin, &request.PhoneNumber)
		if err != nil {
			panic(err)
		}

		marshal, err := json.Marshal(request)
		if err != nil {
			return
		}
		id := uuid.New()
		isValid := auth.IsValid(stringToken)

		if !isValid {

			token = auth.GetToken
			stringToken = token()
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://secure2.1cb.kz/bmg/service/", bytes.NewReader(marshal))
		if err != nil {
			panic(err)
		}
		req.Header.Add("RequestID", id.String())
		req.Header.Add("Consent-Confirmed", "1")
		req.Header.Add(
			"Authorization",
			"Bearer "+stringToken,
		)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var result Result

		result.Iin = request.Iin

		result.Status = string(body)

		log.Print(result.Iin + " - " + result.Status)

		db.Exec("insert into bmg_response(iin, status) VALUES ($1,$2)", result.Iin, result.Status)

	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}

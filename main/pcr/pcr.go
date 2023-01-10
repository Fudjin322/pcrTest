package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

type PcrResponse struct {
	CreditRating int    `json:"creditRating"`
	RiskClass    string `json:"riskClass"`
	RiskCode     int    `json:"riskCode"`
	Message      string `json:"message"`
	Code         int    `json:"code"`
}

type RequestIin struct {
	iin string
}

func main() {

	const connStr = "postgres://admin:admin@localhost:5432/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	query := "select iin from pcr_request"
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

		var requestIin RequestIin

		err := rows.Scan(&requestIin.iin)
		if err != nil {
			return
		}

		uri := fmt.Sprintf("https://test2.1cb.kz/pcr/rating/%v", requestIin.iin)
		fmt.Println(uri)

		client := &http.Client{}
		req, err := http.NewRequest("GET", uri, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Consent-Confirmed", "1")
		req.Header.Add("Authorization", "Basic NzA3NDAxNDc1Nzo3MDc0MDE0NzU3")

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
		fmt.Println(string(body))

		var pcrResponse PcrResponse

		json.Unmarshal(body, &pcrResponse)

		//log.Print(iin + "  " + " " + " " + " ")

		db.Exec(
			"insert into pcr_response(iin, credit_rating,risk_class,risk_code,message,code) VALUES ($1,$2,$3,$4,$5,$6)",
			requestIin.iin,
			pcrResponse.CreditRating, pcrResponse.RiskClass, pcrResponse.RiskCode, pcrResponse.Message,
			pcrResponse.Code,
		)
	}
}

///*
//{
//    "creditRating": 713,
//    "riskClass": "CR10 - «Хорошо»",
//    "riskClassKK": "CR10 - «Жақсы»",
//    "riskClassRU": "CR10 - «Хорошо»",
//    "riskClassEN": "CR10 - «Good»",
//    "color": "#92D050",
//    "riskCode": 4
//}
//*/

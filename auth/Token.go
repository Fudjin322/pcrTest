package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Token struct {
	Access struct {
		Hash      string    `json:"hash"`
		ExpiresAt time.Time `json:"expires_at"`
	} `json:"access"`
	Refresh struct {
		Hash      string    `json:"hash"`
		ExpiresAt time.Time `json:"expires_at"`
	} `json:"refresh"`
	PassChangeNeeded bool `json:"pass_change_needed"`
}

type ValidResponse struct {
	Message string `json:"message"`
}

type ValidRequest struct {
	TokenHash string `json:"token_hash"`
}

func GetToken() string {

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://secure2.1cb.kz/bmg/auth/login/", nil)

	if err != nil {
		fmt.Println(err)
		return "e"
	}

	req.SetBasicAuth("7074014757", "7074014757")
	//req.Header.Add("Authorization", "Basic NzA3NDAxNDc1Nzo3MDc0MDE0NzU3")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "e"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "e"
	}

	var result Token

	json.Unmarshal(body, &result)

	token := result.Access.Hash

	return token
}

func IsValid(token string) bool {

	request := ValidRequest{TokenHash: token}
	marshal, err := json.Marshal(&request)
	if err != nil {
		return false
	}

	//hash := request.TokenHash

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://secure2.1cb.kz/bmg/auth/is/valid/", bytes.NewReader(marshal))

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(body))

	var validResponse ValidResponse

	json.Unmarshal(body, &validResponse.Message)
	if err != nil {
		return false
	}
	fmt.Println(validResponse.Message)

	return true

}

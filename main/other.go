package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//)
//
//type Bodyy struct {
//	Iin          string
//	Phone_number string
//}
//
//func main() {
//	body := Bodyy{Iin: "930106300258", Phone_number: "77074014757"}
//	marshal, err := json.MarshalIndent(body, "", "")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(marshal))
//
//	//MakeRequest(body)
//}
//
//func MakeRequest(body Bodyy) {
//
//	client := &http.Client{}
//	req, _ := http.NewRequest("POST", "https://test2.1cb.kz/bmg/service/", nil)
//
//	req.Header.Set("Consent-Confirmed", "1")
//	req.Header.Set("RequestID", "3eecb5da-65c4-4f09-9e71-c9fdc7433098")
//	req.Header.Set("Authorization", "Basic NzA3NDAxNDc1Nzo3MDc0MDE0NzU3")
//	r, _ := client.Do(req)
//	defer r.Bodyy.Close()
//	io.Copy(os.Stdout, r.Bodyy)
//
//}

//func main() {
//
//	url := "https://bmg.1cb.kz/v1/bmg/check_vshep"
//	method := "GET"
//
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, nil)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Bodyy.Close()
//
//	body, err := ioutil.ReadAll(res.Bodyy)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//}
//
//import (
//	"encoding/csv"
//	"fmt"
//	"io"
//	"strconv"
//	"strings"
//)
//
//type Person struct {
//	Name      string
//	Age       int
//	IsAwesome bool
//}
//
//func MakePerson(record []string) (Person, error) {
//	if len(record) != 3 {
//		return Person{}, fmt.Errorf("invalid person slice: %v", record)
//	}
//	name := record[0]
//	age, err := strconv.Atoi(record[1])
//	if err != nil {
//		return Person{}, fmt.Errorf("invalid Age: %v", record[1])
//	}
//	isAwesome, err := strconv.ParseBool(record[2])
//	if err != nil {
//		return Person{}, fmt.Errorf("invalid IsAwesome: %v", record[2])
//	}
//	return Person{name, age, isAwesome}, nil
//}
//
//func main() {
//	src := `Alice;25;true
//Emma;23;false
//Grace;27;false`
//
//	r := csv.NewReader(strings.NewReader(src))
//	r.Comma = ';'
//
//	var people []Person
//	for {
//		record, err := r.Read()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			panic(err)
//		}
//		person, err := MakePerson(record)
//		if err != nil {
//			panic(err)
//		}
//		people = append(people, person)
//	}
//
//	fmt.Println(people)
//	// [{Alice 25 true} {Emma 23 false} {Grace 27 false}]
//}

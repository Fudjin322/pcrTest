package narcology

import (
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Status string `xml:"status"`
}

func Narcology() string {

	payload := strings.NewReader(
		`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://web.service.narcology.erdb.fcb.com/">
    <soapenv:Header>
        <web:header>
            <Password></Password>
            <UserName></UserName>
        </web:header>
    </soapenv:Header>
    <soapenv:Body>
        <web:sendToErdbNarcology>
            <IIN>680604400733</IIN>
        </web:sendToErdbNarcology>
    </soapenv:Body>
</soapenv:Envelope>`,
	)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://narcology.1cb.kz/narcology/service", payload)

	if err != nil {
		fmt.Println(err)
		return "500"
	}
	req.Header.Add("Content-Type", "text/xml")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "500"
	}
	defer res.Body.Close()

	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))

	return res.Status
}

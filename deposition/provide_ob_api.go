package deposition

import (
	"bytes"
	"log"
	"net/http"
)

func ProvideOpenBankingApi(jsonData []byte) *http.Response {

	resp, err := http.Post("http://localhost:5000/v2_1_1.1/payments/v2_1_1.1/domestic", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	return resp

}

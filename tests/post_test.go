package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	model "github.com/test/transaction/model"
)

func POST(url string, headers map[string]string, content interface{}) ([]byte, error) {
	var postBody io.Reader = nil
	var bodyBytes []byte
	if content != nil {
		bodyBytes, _ = json.Marshal(content)
		postBody = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest("POST", url, postBody)
	if err != nil {
		log.Println("Error crating http request. ", err)
		return nil, err
	}

	req.Header.Set("Cache-Control", "no-cache")
	if content != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error reading response. ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func TestTranAdd(t *testing.T) {

	url := "http://localhost:8000/api/v1/trans"

	var transaction = model.TransactionReq{}
	transaction.Origin = "mobile-android"
	transaction.User_ID = 3
	transaction.Amount = "255.00"
	transaction.Op_Type = "credit"
	transaction.Registered_At = "2022-10-11 04:05:06"

	_, err := POST(url, nil, &transaction)
	if err != nil {
		t.Fatalf(`Error while running TransAdd. %v`, err)
	}
}


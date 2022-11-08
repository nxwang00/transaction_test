package tests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/test/transaction/model"
)

func GET(url string, headers map[string]string) ([]byte, error) {
	var postBody io.Reader = nil

	req, err := http.NewRequest("GET", url, postBody)
	if err != nil {
		log.Println("Error crating http request. ", err)
		return nil, err
	}

	req.Header.Set("Cache-Control", "no-cache")

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

func TestTransGet(t *testing.T) {

	url := "http://localhost:8000/api/v1/trans?page_num=2&page_size=2&origin=mobile-android"

	resp, err := GET(url, nil)
	if err != nil {
		log.Printf("GetAllTransactions: %v\n", err)
	}

	var emptyTransResp string

	var transactions []*model.Transaction
	err = json.Unmarshal(resp, &transactions)
	if err != nil {
		err1 := json.Unmarshal(resp, &emptyTransResp)
		if err1 != nil {
			log.Printf("GetAllTransactions: Unmarshal %v\n", err)
		}
	}

	if emptyTransResp == "There ins`t nothing in DB now" {
		t.Logf("%s", emptyTransResp)
		return
	}

	var trans []model.Transaction
	for _, tran := range transactions {
		trans = append(trans, *tran)
	}
	if err != nil {
		t.Fatalf(`Error while running TransactionsGet. %v`, err)
	}

	t.Logf("%+v\n", trans)
}

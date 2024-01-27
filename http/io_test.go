package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestError(t *testing.T) {

	testRequest := httptest.NewRequest(http.MethodPost, "/something", errReader(0))
	HandlePostRequest(nil, testRequest)
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v\n", err)
	}
	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

}

func TestStr(t *testing.T) {
	bytes := []byte("123")
	fmt.Println(bytes)
	sprintf := fmt.Sprintf("%v", bytes)
	fmt.Println(sprintf)

}

func TestTime(t *testing.T) {
	parse, err := time.Parse("2006-1-02", "2021-11-01")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(parse)

}

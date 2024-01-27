package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	request := http.Request{}
	request.Close = true
	get, err := http.DefaultClient.Get("http://localhost:28000")
	if err != nil {
		panic(err)
	}
	reader := get.Body

	now := time.Now()
	fmt.Println(now)
	_, err = ioutil.ReadAll(reader)
	fmt.Println(time.Since(now))
}

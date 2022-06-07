package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func sendRequest() {
	// Request (GET https://the-trivia-api.com/api/categories)

	json := []byte(`{}`)
	body := bytes.NewBuffer(json)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://the-trivia-api.com/api/categories", body)

	// Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=--AaB03x")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


func main() {
	sendRequest()
}
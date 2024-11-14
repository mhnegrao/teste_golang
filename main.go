package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.example.com/api/v1/create"
	contentType := "application/json"
	data := []byte(`{"name": "Test User", "email": "test@example.com"}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	query:=req.URL.Query()
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")
	query.Add("chave","valor")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

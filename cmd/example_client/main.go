package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Status string `json:"status"`
}

func main() {
	req, _ := http.NewRequest("GET", "http://192.168.1.63:9090/getPr", nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	response := response{}

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParsingData(url string) (string, error) {
	r1, err := http.Get(url)
	defer r1.Body.Close()
	if err != nil {
		return "", err
	}
	//fmt.Println(r1.Status)

	body, err := io.ReadAll(r1.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(body))
	return string(body), nil
}
func WorkwithParsingData() {
	result, err := ParsingData("https://www.google.com/robots.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

type Status struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func ParsingJson(url string) (Status, error) {
	res, err := http.Get(url)
	if err != nil {
		return Status{}, err
	}
	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		return Status{}, err
	}
	defer res.Body.Close()
	return status, nil

}
func WorkwithParsingJson() {
	result, err := ParsingJson("http://127.0.0.1:8080/json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
func main() {

}

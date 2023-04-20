package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Paterns struct {
	NumberInAggoritm string
	PathInEX         string
	PaternInEX       string
}

func main() {

	dan := Paterns{}
	dan.NumberInAggoritm = "first"
	dan.PathInEX = "${7*7}"
	dan.PaternInEX = "49"

	//...................................
	//Writing struct type to a JSON file
	//...................................
	content, err := json.Marshal(dan)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("package.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//...................................
	//Reading into struct type from a JSON file
	//...................................
	content, err = ioutil.ReadFile("package.json")
	if err != nil {
		log.Fatal(err)
	}
	user2 := Paterns{}
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("number = %s \npath = %s \npatern = %s\n", user2.NumberInAggoritm, user2.PathInEX, user2.PaternInEX)

}

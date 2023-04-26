package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	result := make(chan bool)
	url1 := "f'https://directory.web.actf.co/{i}.html'"
	pattern := "actf{\\S+"
	for i := 0; 1 < 5022; i++ {
		url1 = "https://directory.web.actf.co/" + strconv.Itoa(i) + ".html"
		go CheckPusReq(pattern, url1, result)
		rusult2 := <-result
		fmt.Println(i, "\n", rusult2)
		if rusult2 == true {
			fmt.Println("Нашли : ", i)
			break
		}
	}
}
func CheckPusReq(patern string, myurl string, result chan bool) {
	resp, _ := http.Get(myurl) // переделать 2 строку на тестовый путь
	bytes, _ := ioutil.ReadAll(resp.Body)
	Body := string(bytes)
	matched, _ := regexp.MatchString(patern, Body) // пустой «приемник» ошибки, ведь мы уверены, что пример отработает нормально
	//fmt.Println(patern)//отлаживал алгоритм
	result <- matched
} 

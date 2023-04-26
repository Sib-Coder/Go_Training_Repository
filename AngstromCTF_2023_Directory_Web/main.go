package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	url1 := "f'https://directory.web.actf.co/{i}.html'"
	pattern := "actf{\\S+"
	for i := 0; 1 < 5022; i++ {
		url1 = "https://directory.web.actf.co/" + strconv.Itoa(i) + ".html"
		result := CheckPusReq(pattern, url1)
		fmt.Println(i, "\n", result)
		if result == true {
			fmt.Println("Нашли : ", i)
			break
		}
	}
}
func CheckPusReq(patern string, myurl string) bool {
	resp, _ := http.Get(myurl) // переделать 2 строку на тестовый путь
	bytes, _ := ioutil.ReadAll(resp.Body)
	Body := string(bytes)
	matched, _ := regexp.MatchString(patern, Body) // пустой «приемник» ошибки, ведь мы уверены, что пример отработает нормально
	//fmt.Println(patern)//отлаживал алгоритм
	return matched
}

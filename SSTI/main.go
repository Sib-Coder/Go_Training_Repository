package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//surl := flag.String("url", "http://127.0.0.1:5000/", "Give URL Address")
	fmt.Println("Start programm......")
	//flag.Parse() ///парсим
	myurl := "http://127.0.0.1:5000/"
	fmt.Println(myurl)

	///parsing
	result, _ := url.Parse(myurl)

	//fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	fmt.Println("new operation")
	result, _ = url.Parse(myurl + "${7+7}")

	//fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	fmt.Println("new operation")
	result, _ = url.Parse(myurl + "{{7*'7'}}")
	//надо парсить текст страницы
	//fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	///парсинг страницы
	resp, _ := http.Get(myurl + "{{7*'7'}}")
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	///парсинг г страницы
	fmt.Println("Govno Parsing")
	///page1
	result, _ = url.Parse(string(myurl) + "kateloveandpisck69")
	resp1, _ := http.Get(string(myurl) + "kateloveandpisck69")
	bytes1, _ := ioutil.ReadAll(resp1.Body)
	///page2
	result, _ = url.Parse(string(myurl) + "zxhgsudgfdhfhdufgdhgd70")
	resp2, _ := http.Get(string(myurl) + "zxhgsudgfdhfhdufgdhgd70")
	bytes2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println("HTML:\n\n", string(bytes1))
	fmt.Println("HTML:\n\n", string(bytes2))
}

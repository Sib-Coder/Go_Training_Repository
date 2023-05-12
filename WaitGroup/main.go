package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)


func main() {
	urls := []string{ // урлы к которым нужно сделать запросы
		"https://www.google.com/",
		"https://ya.ru/",
		"https://www.youtube.com/",
	}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1) // добавляем горутину в группу
		go func(url string) {
			doHTTP(url)
			wg.Done() // сигнализировать, что элемент группы завершил свое выполнение
		}(url)
	}
	wg.Wait() // ожидаем завершения всех горутин
}

func doHTTP(url string) {
	t := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error %s = %s", url, err.Error())
	}
	defer resp.Body.Close()
	fmt.Println("%s - Status code %d - Latency %d ms", url, resp.StatusCode, time.Since(t).Milliseconds())
}

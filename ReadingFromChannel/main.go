package main

import (
	"fmt"
	"time"
)

func main() {

	massage1 := make(chan string)
	massage2 := make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			massage1 <- "Прошло пол секунды"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			massage2 <- "Прошло 2 секунды"
		}
	}()

	for { // чтение из канала при появлении информации в канале
		select {
		case msg := <-massage1:
			fmt.Println(msg)
		case msg := <-massage2:
			fmt.Println(msg)
		}

	}
	//for { чтение будет последовательным , а не ассинхронным
	//	fmt.Println(<-massage1)
	//	fmt.Println(<-massage2)
	//}
}

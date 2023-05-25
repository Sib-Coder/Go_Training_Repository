package main

import "fmt"

// /функция генерации QR кодов
func CreateSum(x int, y int) int {
	sum := x + y
	return sum
}

// /"exaple/21/32" - проверка работоспособности сервиса
func main() {
	result := CreateSum(3, 4)
	fmt.Println(result)
}

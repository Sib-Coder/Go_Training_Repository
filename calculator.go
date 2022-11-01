package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expression := [][]string{
		[]string{"2", "+", "3"},
		[]string{"3", "*", "3"},
		[]string{"3", "-", "2"},
		[]string{"3", "/", "3"},
		[]string{"3", "mul", "3"},
		[]string{"0X1", "-", "3"},
	}
	for _, expression := range expression {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]     ///занесли оператор
		opFunc, ok := opMap[op] ///поняли какой оператор нужен и вставили его в функцию opFunc
		if !ok {
			fmt.Println("unsupported operator:", ok)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}

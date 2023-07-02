package main

import (
	"fmt"
)

func main() {
	var x, y int
	x = 1
	y = 2
	fmt.Println(comparison(x, y))
}

type ComparisonType interface { //пример написание сообственного типа через интерфейс
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64
}

// как аналог можно использовать   тип Ordered из пакета "golang.org/x/exp/constraints"
func comparison[T ComparisonType](x, y T) T {
	if x > y {
		return x
	}
	return y
}

package main

import (
	"fmt"
	"math"
)

type Share interface {
	Area() float32
}
type Squere struct {
	sideLength float32
}

func (s Squere) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return float32(math.Pi) * c.radius
}
func printShareArea(s Share) {
	fmt.Println("Площадь фигуры : %.2f cm", s.Area())
}
func main() {
	squere := Squere{10.5}
	circle := Circle{5.7}
	printShareArea(squere)
	printShareArea(circle)
}
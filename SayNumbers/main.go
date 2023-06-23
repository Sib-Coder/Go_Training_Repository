package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&z)
	//your code goes here
	numbers := map[int]string{
		0:  "Zero",
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
	}

	fmt.Println(numbers[x])
	fmt.Println(numbers[y])
	fmt.Println(numbers[z])

}
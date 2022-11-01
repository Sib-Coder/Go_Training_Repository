package main

import (
	"fmt"
	"sort"
)

type Student struct {
	last_name string
	firs_name string
	age       int
}

func main() {
	students := []Student{
		{"Daniil", "Sin", 31},
		{"Roman", "Lider", 27},
		{"Anna", "Sin", 12},
	}
	fmt.Println(students)

	sort.Slice(students, func(i int, j int) bool {
		return students[i].last_name < students[j].last_name
	})
	fmt.Println(students)

	sort.Slice(students, func(i int, j int) bool {
		return students[i].age < students[j].age
	})
	fmt.Println(students)

	sort.Slice(students, func(i int, j int) bool {
		return students[i].firs_name < students[j].firs_name
	})
	fmt.Println(students)

}

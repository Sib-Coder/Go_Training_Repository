package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("start programm ...")
	dan := flag.String("name", "Alice", "Name person:")
	bullshet := flag.Bool("bools", false, "True of False?")
	flag.Parse()///парсим

	daniil := *dan + " sin"
	fmt.Println(daniil)
	fmt.Println(*bullshet == true)///обращаемся к содержимому через *
}

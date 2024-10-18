package main

import (
	"fmt"
)

func main() {
	request()
	fmt.Println("-----------------------------------------------------")
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}
	printArea(t)
	printArea(s)
}

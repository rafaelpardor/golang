package main

import (
	"fmt"
)

func main() {
	x := [6]int{98, 93, 77, 82, 83}

	fmt.Println(x)
	fmt.Println("Lenght of the array:", len(x)) // We print the lenght of the array
	fmt.Printf("Type-%T\n", x)
}
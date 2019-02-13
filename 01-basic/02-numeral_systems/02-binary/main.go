package main

import "fmt"

func main() {
	fmt.Println("Binary numbers")
	fmt.Println("------")

	for i := 1; i < 11; i++ {
		fmt.Printf("%b \n", i) // We use %b to print decimalless scientific notation with exponent a power of two.
	}
}

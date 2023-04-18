package main

import (
	"calcu/calcu"
	"fmt"
)

func main() {
	var a, b, ch int
	fmt.Println("Enter 1st Number: ")
	fmt.Scan(&a)
	fmt.Println("Enter 2nd Number: ")
	fmt.Scan(&b)

	fmt.Println("1: Addition")
	fmt.Println("2: Subtraction")
	fmt.Println("3: Multiplication")
	fmt.Println("4: Division")
	fmt.Println("Enter choice: ")
	fmt.Scan(&ch)

	switch ch {
	case 1:
		fmt.Printf("Additon: %d \n", calcu.Addition(a, b))
	case 2:
		fmt.Printf("Additon: %d \n", calcu.Subtraction(a, b))
	case 3:
		fmt.Printf("Additon: %d \n", calcu.Multiplication(a, b))
	case 4:
		fmt.Printf("Additon: %d \n", calcu.Division(a, b))
	default:
		fmt.Println("Invalid value")
	}
}

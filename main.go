package main

import (
	"calc/calc"
	"fmt"
)

func main() {
	var ch int
	var a, b float64
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%f", &b)

	calc := calc.MathFloat{No1: a, No2: b}

	fmt.Println("1: Addition")
	fmt.Println("2: Subtraction")
	fmt.Println("3: Multiplication")
	fmt.Println("4: Division")
	fmt.Println("Enter choice: ")
	fmt.Scan(&ch)

	switch ch {
	case 1:
		fmt.Printf("Additon: %f \n", calc.Addition())
	case 2:
		fmt.Printf("Subtraction: %f \n", calc.Subtraction())
	case 3:
		fmt.Printf("Multiplication: %f \n", calc.Multiplication())
	case 4:
		fmt.Printf("Division: %f \n", calc.Division())
	default:
		fmt.Println("Invalid value")
	}
}

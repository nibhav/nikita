package main

import (
	"fmt"
	"mcalci/mcalci"
)

func main() {
	input := mcalci.CalculatorInput{A: 10, B: 5}

	sum := input.Add()
	fmt.Printf("%.2f + %.2f = %.2f\n", input.A, input.B, sum)

	diff := input.Subtract()
	fmt.Printf("%.2f - %.2f = %.2f\n", input.A, input.B, diff)

	product := input.Multiply()
	fmt.Printf("%.2f * %.2f = %.2f\n", input.A, input.B, product)

	quotient := input.Divide()
	fmt.Printf("%.2f * %.2f = %.2f\n", input.A, input.B, quotient)
}

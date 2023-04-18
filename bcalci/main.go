package main

import (
	fcalculator "fcalculator/bcalci"
	"fmt"
)

type Calculator struct {
	a int
	b int
}


func main() {
	var ch int 
	var a, b int
   fmt.Print("Enter the first number: ")
   fmt.Scanf("%d", &a)
   fmt.Print("Enter the second number: ")
   fmt.Scanf("%d", &b)
   cal := Calculator{
      a: a,
      b: b,
   }

	fmt.Println("1: Addition")
	fmt.Println("2: Subtraction")
	fmt.Println("3: Multiplication")
	fmt.Println("4: Division")
	fmt.Println("Enter choice: ")
	fmt.Scan(&ch)

	switch ch {
	case 1:
		fmt.Printf("Additon: %d \n", fcalculator.Add(cal.a, cal.b))
	case 2:
		fmt.Printf("Subtraction: %d \n", fcalculator.Sub(cal.a, cal.b))
	case 3:
		fmt.Printf("Multiplication: %d \n", fcalculator.Mul(cal.a, cal.b))
	case 4:
		fmt.Printf("Division: %d \n", fcalculator.Div(cal.a, cal.b))
	default:
		fmt.Println("Invalid value")
	}
}

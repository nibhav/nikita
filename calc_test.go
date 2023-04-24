package calc_test

import (
	"calc/calc"
	"testing"
)

func TestAddition(t *testing.T) {
	//Input := calc.MathInt{Num1: 10, Num2: 5}
	Input := calc.MathFloat{No1: 2.5, No2: 2.5}
	//expected := 15
	expected := 5.0
	result := Input.Addition()
	if result != expected {
		t.Errorf("Add(%v): expected %v", Input, expected)
	}
}

func TestSubtraction(t *testing.T) {
	//Input := calc.MathInt{Num1: 10, Num2: 5}
	Input := calc.MathFloat{No1: 2.5, No2: 2.0}
	expected := 0.5
	result := Input.Subtraction()
	if result != expected {
		t.Errorf("Add(%v): expected %v", Input, expected)
	}
}

func TestMultiplication(t *testing.T) {
	//Input := calc.MathInt{Num1: 10, Num2: 5}
	Input := calc.MathFloat{No1: 2.5, No2: 1.5}
	expected := 3.75
	result := Input.Multiplication()
	if result != expected {
		t.Errorf("Add(%v): expected %v", Input, expected)
	}
}

func TestDivision(t *testing.T) {
	//Input := calc.MathInt{Num1: 10, Num2: 5}
	Input := calc.MathFloat{No1: 2.5, No2: 2.0}
	expected := 1.25
	result := Input.Division()
	if result != expected {
		t.Errorf("Add(%v): expected %v", Input, expected)
	}
}

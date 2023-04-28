package mcalci_test

import (
	"mcalci/mcalci"
	"testing"
)

func TestAdd(t *testing.T) {
	input := mcalci.CalculatorInput{A: 10.0, B: 5.0}
	expected := 15.0
	if actual := input.Add(); actual != expected {
		t.Errorf("Add(%v): expected %v, actual %v", input, expected, actual)
	}
}

func TestSubtract(t *testing.T) {
	input := mcalci.CalculatorInput{A: 10.0, B: 5.0}
	expected := 5.0
	if actual := input.Subtract(); actual != expected {
		t.Errorf("Subtract(%v): expected %v, actual %v", input, expected, actual)
	}
}

func TestMultiply(t *testing.T) {
	input := mcalci.CalculatorInput{A: 10.0, B: 5.0}
	expected := 50.0
	if actual := input.Multiply(); actual != expected {
		t.Errorf("Multiply(%v): expected %v, actual %v", input, expected, actual)
	}
}

func TestDivide(t *testing.T) {
	input := mcalci.CalculatorInput{A: 10.0, B: 5.0}
	expected := 2.0
	if actual := input.Divide(); actual != expected {
		t.Errorf("Multiply(%v): expected %v, actual %v", input, expected, actual)
	}
}

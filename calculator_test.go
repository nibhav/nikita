package main

import (
   "testing"
)

func TestAdd(t *testing.T) {
	expected := 5
	result := addition(2, 3)
	if expected != result {
	  t.Errorf("%d was expect but got %d .\n", expected, result)
	}
}
func TestSubtract(t *testing.T) {
	expected := 2
	result := subtract(5, 3)
	if result != expected {
		t.Errorf("%d was expect but got %d .\n", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	expected := 10
	result := multiply(2, 5)
	if result != expected {
		t.Errorf("%d was expect but got %d .\n", expected, result)
	}
}

func TestDivision(t *testing.T) {
	expected := 2
	result := division(6, 3)
	if result != expected {
		t.Errorf("%d was expect but got %d .\n", expected, result)
	}
}


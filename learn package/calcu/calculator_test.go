package calcu_test

import (
	"calcu/calcu"
	"testing"
)

func TestAddition(t *testing.T) {
	expected := 5
	result := calcu.Addition(2, 3)
	if expected != result {
		t.Errorf("%d was expect but got %d .\n", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	expected := 2
	result := calcu.Subtraction(5, 3)
	if result != expected {
		t.Errorf("%d was expect but got %d .\n", expected, result)
	}
}

func TestMultiplication(t *testing.T) {
	expected := 10
	result := calcu.Multiplication(2, 5)
	if result != expected {
		t.Errorf("%d was expect but got %d .\n", expected, result)
	}
}

func TestDivision(t *testing.T) {
	expected := 2
	result := calcu.Division(6, 3)
	if result != expected {
		t.Errorf("%d was expect but got %d .\n", expected, result)
	}
}

func BenchmarkAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcu.Addition(4, 6)
	}
}

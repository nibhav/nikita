package fcalculator_test

import (
	fcalculator "fcalculator/bcalci"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	got := fcalculator.Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSub(t *testing.T) {

	got := fcalculator.Sub(6, 4)
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fcalculator.Add(4, 6)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fcalculator.Sub(6, 4)
	}
}

func ExampleAdd() {
	fmt.Println(fcalculator.Add(4, 6))
	// Output: 10
}
func ExampleSub() {
	fmt.Println(fcalculator.Sub(6, 4))
	// Output: 2
}

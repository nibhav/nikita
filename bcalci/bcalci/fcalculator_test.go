package fcalculator_test

import (
	"fcalculator/bcalci"
	"fmt"
	"testing"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'


// type subTest struct {
// 	arg1, arg2, expected int
// }
// var subTests = []subTest{
// 	subTest{2, 3, 5},
// }



func TestAdd(t *testing.T){

    got := fcalculator.Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
 func TestSub(t *testing.T){

     got := fcalculator.Sub(4, 6)
     want := -2

     if got != want {
         t.Errorf("got %q, wanted %q", got, want)
     }
 }


func BenchmarkAdd(b *testing.B){
    for i :=0; i < b.N ; i++{
        fcalculator.Add(4, 6)
    }
}


func BenchmarkSub(b *testing.B){
     for i :=0; i < b.N ; i++{
         fcalculator.Sub(4, 6)
     }
 }


func ExampleAdd() {
    fmt.Println(fcalculator.Add(4, 6))
    // Output: 10
}
func ExampleSub() {
     fmt.Println(fcalculator.Sub(5,9))
     // Output: 2
}

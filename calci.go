package mcalci

type CalculatorInput struct {
	A float64
	B float64
}

func (c CalculatorInput) Add() float64 {
	return c.A + c.B
}

func (c CalculatorInput) Subtract() float64 {
	return c.A - c.B
}

func (c CalculatorInput) Multiply() float64 {
	return c.A * c.B
}

func (c CalculatorInput) Divide() float64 {
	return c.A / c.B
}

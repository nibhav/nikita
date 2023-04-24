package calc

type Maths interface {
	Addtion()
	Subtraction()
	Multiplication()
	division()
}

//type MathInt struct {
//	Num1 int
//	Num2 int
//}
type MathFloat struct {
	No1 float64
	No2 float64
}

//func (value *MathInt) Addition() int {
//	return value.Num1 + value.Num2
//}

//func (value *MathInt) Subtraction() int {
//	return value.Num1 - value.Num2
//}

//func (value *MathInt) Multiplication() int {
//	return value.Num1 * value.Num2
//}

//func (value *MathInt) Division() int {
//	return value.Num1 / value.Num2
//}

func (value *MathFloat) Addition() float64 {
	return value.No1 + value.No2
}

func (value *MathFloat) Subtraction() float64 {
	return value.No1 - value.No2
}

func (value *MathFloat) Multiplication() float64 {
	return value.No1 * value.No2
}

func (value *MathFloat) Division() float64 {
	return value.No1 / value.No2
}

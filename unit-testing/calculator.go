package unit_testing

type Calculator struct {
	//CalculatorInterface
}

//type CalculatorInterface interface {
//	Add(a, b int) int
//	Subtract(a, b int) int
//	Multiply(a, b int) int
//}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Subtract(a, b int) int {
	return a - b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

package mathutil

func Factorial(a int) int {
	var result int
	result = 1
	for i := 1; i <= a; i++ {
		result *= i
	}
	return result
}

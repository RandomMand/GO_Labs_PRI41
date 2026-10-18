// Метод вращения (Гивенса)
// Вариант: A = | 1.03    0.991  |   b = | 2.51 |
//
//	| 0.2973  0.2829 |       | 2.43 |
//
// Точность 0.001
package main

import "fmt"

const N = 2

type TMatr [N][N]float64
type TVector [N]float64

// Модуль числа (без пакета math)
func myAbs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}

// Квадратный корень методом Ньютона (без пакета math)
func mySqrt(v float64) float64 {
	if v <= 0 {
		return 0
	}
	r := v
	if r < 1 {
		r = 1
	}
	for i := 0; i < 100; i++ {
		next := 0.5 * (r + v/r)
		if myAbs(next-r) < 1e-15*next {
			return next
		}
		r = next
	}
	return r
}

// Метод вращения: приведение к верхнетреугольному виду
// поворотами Гивенса и обратный ход
func vrash(nn int, aa TMatr, bb TVector) (x TVector) {
	var m, l, r float64
	for i := 0; i < nn-1; i++ {
		for k := i + 1; k < nn; k++ {
			if aa[i][i] != 0 || aa[k][i] != 0 {
				m = mySqrt(aa[i][i]*aa[i][i] + aa[k][i]*aa[k][i])
				l = -aa[k][i] / m // -sin
				m = aa[i][i] / m  //  cos
			} else {
				m, l = 1, 0
			}
			for j := 0; j < nn; j++ {
				r = m*aa[i][j] - l*aa[k][j]
				aa[k][j] = l*aa[i][j] + m*aa[k][j]
				aa[i][j] = r
			}
			r = m*bb[i] - l*bb[k]
			bb[k] = l*bb[i] + m*bb[k]
			bb[i] = r
		}
	}

	// Обратный ход
	for i := nn - 1; i >= 0; i-- {
		s := 0.0
		for j := i + 1; j < nn; j++ {
			s += aa[i][j] * x[j]
		}
		x[i] = (bb[i] - s) / aa[i][i]
	}
	return x
}

func main() {
	a := TMatr{
		{1.03, 0.991},
		{0.2973, 0.2829},
	}
	b := TVector{2.51, 2.43}

	x := vrash(N, a, b)
	fmt.Println("Метод вращения (Гивенса)")
	fmt.Printf("x1 = %.4f\n", x[0])
	fmt.Printf("x2 = %.4f\n\n", x[1])

	// Возмущение правой части на 0.005
	bv := TVector{b[0] + 0.005, b[1] + 0.005}
	x = vrash(N, a, bv)
	fmt.Println("При возмущении правой части на 0.005:")
	fmt.Printf("x1 = %.4f\n", x[0])
	fmt.Printf("x2 = %.4f\n", x[1])
}

package main

import (
	"fmt"
)

func chetnoeNechetnoe(chislo int) string {
	if chislo%2 == 0 {
		return "Chetnoe"
	}
	return "NeChetnoe"
}

func znak(chislo int) string {
	if chislo < 0 {
		return "Negative"
	}
	if chislo > 0 {
		return "Positive"
	}
	if chislo == 0 {
		return "Zero"
	}
	return "Ввели не число"
}

func dlinaStroki(stroka string) int {
	return len(stroka)
}

type Rectangle struct {
	a int
	b int
}

func Ploshad(r Rectangle) int {
	return r.a * r.b
}

func avgNumbers(number1 int, number2 int) int {
	return number1 + number2/2
}
func main() {
	var chislo int
	fmt.Println("Введите любое число для задания 1: ")
	fmt.Scan(&chislo)
	fmt.Println(chetnoeNechetnoe(chislo))

	fmt.Println("Введите любое число для задания 2: ")
	fmt.Scan(&chislo)
	fmt.Println(znak(chislo))

	for i := 1; i <= 10; i++ {
		fmt.Print(i)
	}
	fmt.Println("")

	var stroka string
	fmt.Println("Введите строку для измерения длинны")
	fmt.Scan(&stroka)
	fmt.Println(dlinaStroki(stroka))

	var rect Rectangle
	rect.a = 10
	rect.b = 10
	fmt.Println("Площадь прямоугольника ", Ploshad(rect))

	var chislo1 int
	var chislo2 int
	fmt.Println("Введите первое число")
	fmt.Scan(&chislo1)
	fmt.Println("Введите второе число")
	fmt.Scan(&chislo2)

	fmt.Println("Среднее из 2ух чисел ", avgNumbers(chislo1, chislo2))

}

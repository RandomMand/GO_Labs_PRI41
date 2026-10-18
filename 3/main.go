package main

import (
	"Laba3/mathutil"
	"Laba3/stringutil"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Введите число для факторила: ")
	var factNumber int
	fmt.Scan(&factNumber)
	fmt.Println("Факториал числа ", factNumber, " равно ",
		mathutil.Factorial(factNumber))

	fmt.Print("Введите строку для переворота: ")
	var stroka string
	fmt.Scan(&stroka)

	fmt.Println("Перевернутая строка ", stringutil.ReverseStr(stroka))

	var array [5]int
	fmt.Print("Массив чисел: ")
	for i := 0; i < 5; i++ {
		array[i] = rand.Intn(50)
	}
	fmt.Println(array)

	//splice[:]
	var srez int
	fmt.Println("Введите срез массива: ")
	fmt.Scan(&srez)
	fmt.Print("Введенный срез: ")
	for i := 0; i < srez; i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Println("")
	fmt.Println("Добавить элемент или удалить?")
	var choose string
	fmt.Scan(&choose)
	if choose == "Добавить" {
		var addArray []int
		addArray = append(array[:], rand.Intn(50))
		fmt.Println("Обновленный: ", addArray)
	} else if choose == "Удалить" {
		var delArray []int
		delArray = array[1:]
		fmt.Println("Удаленный: ", delArray)
	}

	words := []string{"яблоко", "банан", "апельсин", "киви", "виноград"}
	fmt.Println("Срез строк: ", words)

	longest := words[0]
	for _, w := range words {
		if len([]rune(w)) > len([]rune(longest)) {
			longest = w
		}
	}
	fmt.Println("Самая длинная строка: ", longest)
}

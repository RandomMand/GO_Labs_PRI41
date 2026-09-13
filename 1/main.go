package main

import (
	"fmt"
	"time"
)

func currentDateAndTime() string {
	today := time.Now()
	return today.Format("2006-01-02 15:04")
}
func arifmeticheskoeInt(a int, b int, operationType string) int {
	var result int
	if operationType == "+" {
		result = a + b
	}
	if operationType == "-" {
		result = a - b
	}
	if operationType == "/" {
		result = a / b
	}
	if operationType == "*" {
		result = a * b
	}
	return result
}
func arifmeticheskoeFloat(a float32, b float32, operationType string) float32 {
	var result float32
	if operationType == "+" {
		result = a + b
	}
	if operationType == "-" {
		result = a - b
	}
	if operationType == "/" {
		result = a / b
	}
	if operationType == "*" {
		result = a * b
	}
	return result
}

func avg3Numbers(numbers [3]float32) float32 {
	var result float32
	for _, number := range numbers {
		result += number
	}
	return result / float32(len(numbers))
}
func main() {
	fmt.Println("Задание 1: \nТекущая дата и время " + currentDateAndTime())

	var integer int = 67
	var floatSixFour float64 = 1.212412
	var stroka string = "boy next door"
	var is_true bool = true

	fmt.Println("\nЗадание 2:")
	fmt.Println(" int", integer, "\n",
		"float64 ", floatSixFour, "\n",
		"boolean ", is_true, "\n",
		"string ", stroka)

	fmt.Println("\nЗадание 3:")
	navernoeStroka := "strokovoe znachenie"
	fmt.Println(navernoeStroka)

	fmt.Println("\nЗадание 4:")
	fmt.Println(" Сложение", arifmeticheskoeInt(2, 2, "+"), "\n",
		"Вычитание", arifmeticheskoeInt(2, 2, "-"), "\n",
		"Деление", arifmeticheskoeInt(2, 2, "/"), "\n",
		"Умножение", arifmeticheskoeInt(2, 2, "*"), "\n")

	fmt.Println("\nЗадание 5:")
	fmt.Println(" Сложение", arifmeticheskoeFloat(2.15, 2.15, "+"), "\n",
		"Вычитание", arifmeticheskoeFloat(2.15, 2.15, "-"), "\n",
		"Деление", arifmeticheskoeFloat(2.15, 2.15, "/"), "\n",
		"Умножение", arifmeticheskoeFloat(2.15, 2.15, "*"), "\n")

	fmt.Println("\nЗадание 6:")
	fmt.Println(avg3Numbers([3]float32{5, 2.15, 2.15}))
}

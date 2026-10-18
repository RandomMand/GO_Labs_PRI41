// Лабораторная работа 4
//
// Запуск: go run main.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// общий читатель ввода для всех заданий
var reader = bufio.NewReader(os.Stdin)

// readLine считывает одну строку с ввода без символа перевода строки
func readLine() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\r\n")
}

// printPeople выводит все записи карты (по алфавиту, чтобы порядок был всегда одинаковый)
func printPeople(people map[string]int) {
	names := make([]string, 0, len(people))
	for name := range people {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, people[name])
	}
}

// Задание 2: средний возраст всех людей в карте
func averageAge(people map[string]int) float64 {
	if len(people) == 0 {
		return 0
	}
	sum := 0
	for _, age := range people {
		sum += age
	}
	return float64(sum) / float64(len(people))
}

// readInts считывает строку и превращает её в срез целых чисел (через пробел)
func readInts() []int {
	for {
		fields := strings.Fields(readLine())
		nums := make([]int, 0, len(fields))
		ok := true
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				fmt.Printf("\"%s\" не является целым числом, введите заново: ", f)
				ok = false
				break
			}
			nums = append(nums, n)
		}
		if ok {
			return nums
		}
	}
}

func main() {
	// ---------- Задание 1 ----------
	fmt.Println("Задание 1: ")
	people := map[string]int{
		"Иван":  25,
		"Мария": 31,
		"Пётр":  19,
	}
	people["Анна"] = 42 // добавляем нового человека
	printPeople(people)
	fmt.Println()

	// ---------- Задание 2 ----------
	fmt.Println("Задание 2: ")
	fmt.Printf("Средний возраст: %.2f\n", averageAge(people))
	fmt.Println()

	// ---------- Задание 3 ----------
	fmt.Println("Задание 3: ")
	fmt.Print("Введите имя для удаления: ")
	name := strings.TrimSpace(readLine())
	if _, found := people[name]; found {
		delete(people, name)
		fmt.Printf("Запись \"%s\" удалена\n", name)
	} else {
		fmt.Printf("Человека с именем \"%s\" нет в карте\n", name)
	}
	printPeople(people)
	fmt.Println()

	// ---------- Задание 4 ----------
	fmt.Println("Задание 4: ")
	fmt.Print("Введите строку: ")
	text := readLine()
	fmt.Println(strings.ToUpper(text))
	fmt.Println()

	// ---------- Задание 5 ----------
	fmt.Println("Задание 5: ")
	fmt.Print("Введите числа через пробел: ")
	numbers := readInts()
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println("Сумма:", sum)
	fmt.Println()

	// ---------- Задание 6 ----------
	fmt.Println("Задание 6: ")
	fmt.Print("Введите целые числа через пробел: ")
	arr := readInts()
	fmt.Print("В обратном порядке:")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(" ", arr[i])
	}
	fmt.Println()
}

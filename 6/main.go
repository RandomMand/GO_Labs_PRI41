package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Факториал числа", n, "равен", result)
}

func randomNumbers(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		fmt.Println("Случайное число:", rand.Intn(100))
		time.Sleep(150 * time.Millisecond)
	}
}

func sumRow(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Println("Сумма ряда от 1 до", n, "равна", sum)
}

func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
		time.Sleep(50 * time.Millisecond)
	}
	close(ch)
}

func printNumbers(ch chan int, done chan bool) {
	for num := range ch {
		fmt.Print(num, " ")
	}
	fmt.Println("")
	done <- true
}

type Request struct {
	a  float64
	b  float64
	op string
}

func calcServer(requests chan Request, results chan string) {
	for req := range requests {
		go func(r Request) {
			time.Sleep(100 * time.Millisecond)
			var text string
			if r.op == "+" {
				text = fmt.Sprintf("%.2f + %.2f = %.2f", r.a, r.b, r.a+r.b)
			} else if r.op == "-" {
				text = fmt.Sprintf("%.2f - %.2f = %.2f", r.a, r.b, r.a-r.b)
			} else if r.op == "*" {
				text = fmt.Sprintf("%.2f * %.2f = %.2f", r.a, r.b, r.a*r.b)
			} else if r.op == "/" {
				if r.b == 0 {
					text = fmt.Sprintf("%.2f / %.2f = деление на ноль", r.a, r.b)
				} else {
					text = fmt.Sprintf("%.2f / %.2f = %.2f", r.a, r.b, r.a/r.b)
				}
			} else {
				text = "Неизвестная операция " + r.op
			}
			results <- text
		}(req)
	}
}

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func worker(id int, tasks chan string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range tasks {
		time.Sleep(100 * time.Millisecond)
		results <- fmt.Sprintf("Воркер %d: %s -> %s", id, line, reverseStr(line))
	}
}

func main() {
	fmt.Println("Задание 1: ")
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(3)
	go factorial(5, &wg)
	go randomNumbers(4, &wg)
	go sumRow(10, &wg)
	wg.Wait()
	fmt.Println("Время выполнения:", time.Since(start))
	fmt.Println("")

	fmt.Println("Задание 2: ")
	fibCh := make(chan int)
	done := make(chan bool)
	go fibonacci(10, fibCh)
	go printNumbers(fibCh, done)
	<-done
	// close() закрывает канал, после этого range в printNumbers
	// понимает что чисел больше не будет и завершается.
	// Без close() горутина вечно ждала бы новое число (deadlock)
	fmt.Println("Канал закрыт через close(), чтение завершено")
	fmt.Println("")

	fmt.Println("Задание 3: ")
	numCh := make(chan int)
	checkCh := make(chan int)
	msgCh := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			n := rand.Intn(100)
			numCh <- n
			checkCh <- n
			time.Sleep(100 * time.Millisecond)
		}
		close(checkCh)
	}()
	go func() {
		for n := range checkCh {
			if n%2 == 0 {
				msgCh <- fmt.Sprint(n, " - чётное")
			} else {
				msgCh <- fmt.Sprint(n, " - нечётное")
			}
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case n := <-numCh:
			fmt.Println("Получено число:", n)
		case msg := <-msgCh:
			fmt.Println("Сообщение:", msg)
		}
	}
	fmt.Println("")

	fmt.Println("Задание 4: ")
	counter := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			counter++
		}()
	}
	wg2.Wait()
	fmt.Println("Без мьютекса счётчик =", counter, "(должно быть 1000)")

	counter = 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Println("С мьютексом счётчик =", counter)
	fmt.Println("")

	fmt.Println("Задание 5: ")
	requests := make(chan Request)
	results := make(chan string)
	go calcServer(requests, results)

	tasks := []Request{
		{10, 5, "+"},
		{20, 8, "-"},
		{3, 7, "*"},
		{15, 4, "/"},
		{9, 0, "/"},
	}
	for i := 0; i < len(tasks); i++ {
		go func(r Request) {
			requests <- r
		}(tasks[i])
	}
	for i := 0; i < len(tasks); i++ {
		fmt.Println("Результат:", <-results)
	}
	close(requests)
	fmt.Println("")

	fmt.Println("Задание 6: ")
	text := "привет\nмир\ngolang\nгорутины\nканалы\nворкер\nлабораторная"
	os.WriteFile("input.txt", []byte(text), 0644)

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	fmt.Print("Введите количество воркеров: ")
	var workersCount int
	fmt.Scan(&workersCount)
	if workersCount < 1 {
		workersCount = 1
	}

	taskCh := make(chan string, len(lines))
	resultCh := make(chan string, len(lines))
	var wg3 sync.WaitGroup
	for i := 1; i <= workersCount; i++ {
		wg3.Add(1)
		go worker(i, taskCh, resultCh, &wg3)
	}
	for _, line := range lines {
		taskCh <- line
	}
	close(taskCh)
	wg3.Wait()
	close(resultCh)

	output := ""
	for res := range resultCh {
		fmt.Println(res)
		output = output + res + "\n"
	}
	os.WriteFile("output.txt", []byte(output), 0644)
	fmt.Println("Результаты записаны в файл output.txt")
}

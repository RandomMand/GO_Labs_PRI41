package main

import "fmt"

type Person struct {
	name string
	age  int
}

func birthday(pers *Person) {
	pers.age = pers.age + 1
}

type Circle struct {
	radius float32
}

func ploshadCircle(krug *Circle) float32 {
	return krug.radius * krug.radius * 3.14
}

type Shape interface {
	Area() float32
}

type Rectangle struct {
	shirina float32
	visota  float32
}

func (pryamougolnik Rectangle) Area() float32 {
	return pryamougolnik.shirina * pryamougolnik.visota
}

func (krug Circle) Area() float32 {
	return krug.radius * krug.radius * 3.14
}

func vivodPloshadey(figuri []Shape) {
	for i := 0; i < len(figuri); i++ {
		fmt.Println("Площадь фигуры", i+1, "равна : ", figuri[i].Area())
	}
}

type Stringer interface {
	String() string
}

type Book struct {
	nazvanie string
	avtor    string
	god      int
}

func (kniga Book) String() string {
	return "Книга: " + kniga.nazvanie + ", автор: " + kniga.avtor + ", год: " + fmt.Sprint(kniga.god)
}

func main() {
	fmt.Println("Задание 1: ")
	var diman Person
	diman.name = "Dmitriy"
	diman.age = 21
	fmt.Println(diman)

	fmt.Println("Задание 2: ")
	birthday(&diman)
	fmt.Println(diman)

	fmt.Println("Задание 3: ")
	var krug Circle
	krug.radius = 15
	fmt.Println("Площадь круга равна : ", ploshadCircle(&krug))

	fmt.Println("Задание 4: ")
	var pryamougolnik Rectangle
	pryamougolnik.shirina = 4
	pryamougolnik.visota = 6
	var figura Shape
	figura = pryamougolnik
	fmt.Println("Площадь прямоугольника равна : ", figura.Area())
	figura = krug
	fmt.Println("Площадь круга равна : ", figura.Area())

	fmt.Println("Задание 5: ")
	figuri := []Shape{pryamougolnik, krug}
	vivodPloshadey(figuri)

	fmt.Println("Задание 6: ")
	var kniga Book
	kniga.nazvanie = "Мастер и Маргарита"
	kniga.avtor = "Булгаков"
	kniga.god = 1967
	var str Stringer
	str = kniga
	fmt.Println(str.String())
}

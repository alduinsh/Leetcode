package main

import (
	"fmt"
)

//1

func main() {
		var (
			nameOfFilm string
			yearOfProd int16
			genre      string
			rating     float32
			favourite  bool
		)
		fmt.Println(nameOfFilm, yearOfProd, genre, rating, favourite)
	}

//2 нет

//3 нужно переприсвоить переменную

//4 float

const (
	January = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func main() {
	fmt.Println("January:", January)
	fmt.Println("February:", February)
	fmt.Println("March:", March)
	fmt.Println("April:", April)
	fmt.Println("May:", May)
	fmt.Println("June:", June)
	fmt.Println("July:", July)
	fmt.Println("August:", August)
	fmt.Println("September:", September)
	fmt.Println("October:", October)
	fmt.Println("November:", November)
	fmt.Println("December:", December)
}


//5 1

//6

var (
	a int
	b int = 9
	c int = 8
	d int = 7
)

func main() {
	a = a + 1
	b = a + 2*b
	c = (b - a) + c
	d = a*b + c - d
	fmt.Println(a, b, c, d)
}


//7

var a, b int = 3, 7

func main() {
	a, b = b, a
	fmt.Println(a, b)
}


//8

const r = 6371

func main() {
	d := 2 * r
	fmt.Printf("Диамметр земли равен %d", d)
}


//9

func main() {
	// Общее количество конфет в килограммах
	totalCandies := 4.86

	// Количество детей
	numberOfChildren := 5

	// Рассчитываем количество конфет на одного ребенка
	candiesPerChild := totalCandies / float64(numberOfChildren)

	// Выводим результат
	fmt.Printf("Каждый ребенок получит %.2f килограммов конфет.\n", candiesPerChild)
}


//10

func main() {

	day := 1
	hour := 18
	min := 56

	eveNewYearDay := day * 24 * 60
	eveNewYearHour := hour * 60
	eveNewYearMin := eveNewYearDay + eveNewYearHour + min
	fmt.Println("Минут до нового года:", eveNewYearMin)
}


//11

func main() {
	min := 472
	minOfHour := 60
	hourOfExm := min / minOfHour
	minOfExam := min % minOfHour
	fmt.Printf("До важного экзамена: %d часов и %d минуты\n", hourOfExm, minOfExam)
}


//12

func main() {
	a := 5.78
	b := int(a)
	c := b * 34
	fmt.Println(b, c)
	e := a * 34
	g := int(e)
	fmt.Println(e, g)
}


//13

func main() {
a := true
b := true
c := false
fmt.Println(c == b, a != b, a != c)

a := 1
b := 2
a1 := 1
fmt.Println(a == a1,
	a != a1,
	a == b,
	a != b,
	a <= a1,
	a < b,
	a >= a1)

a := 2.67
b := 1.9
a1 := 2.67
fmt.Println(a == a1,
	a != a1, // true вероятно из-за особенностей языка?
	a == b,
	a != b,
	a <= a1,
	a < b,
	a >= a1,
	a > b)

	a := "mango"
	b := "tomato"
	a1 := "mango"
	fmt.Println(a == a1,
		a != a1,
		a == b,
		a != b,
		a <= a1,
		a < b,
		a >= a1,
		a > b)
}

14
func main() {
	a := 1
    b := float32(1)
	fmt.Println(a == b)
}

func main() {
	helloworld := "helloworld"
	slyceByte = []byte(helloworld)
	fmt.Println(slyceByte)
}

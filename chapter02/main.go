package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

const (
	Pi               = 3.14
	Avogadro float32 = 6.022e23
)

type DayOfWeek uint8

const (
	Monday DayOfWeek = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println("Hello World")

	var a int
	a = 42

	var aa int = 100

	b := -25

	c := "this is a string"

	var d, e string
	d, e = "var d", "var e"

	f, g := true, false

	fmt.Println(a)
	fmt.Println(aa)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var aBool bool = true
	var aString string = "yXXxy"
	var aComplex complex64 = 5i
	var aRune = '€'

	fmt.Println(aBool)
	fmt.Println(aString)
	fmt.Println(aComplex)
	fmt.Println(aRune)
	fmt.Printf("%c \n", aRune)

	fmt.Println("What is the value of Pi? Pi is", Pi)
	fmt.Println(reflect.TypeOf(Pi))
	fmt.Println("Avogadro number is", Avogadro)
	fmt.Println(reflect.TypeOf(Avogadro))

	fmt.Printf("Monday is %d\n", Monday)
	fmt.Printf("Wednesday is %d\n", Wednesday)
	fmt.Printf("Friday is %d\n", Friday)

	result := sum(2, 2)
	fmt.Println(result)

	_sum, diff := ops(5, 3)
	fmt.Println(_sum, diff)

	total := sum_all(1, 2, 3, 4, 5)
	fmt.Println("the first five numbers sum is", total)

	_c := doit(2, 3, sum)
	_d := doit(2, 3, multiply)

	fmt.Println(_c)
	fmt.Println(_d)

	_a := accumulator(1)
	_b := accumulator(4)

	fmt.Println("a", "b")
	for i := 0; i < 5; i++ {
		fmt.Println(_a(), _b())
	}

	a = 10
	zero_value(a)
	fmt.Println(a)
	zero_point(&a)
	fmt.Println(a)

	rand.Seed(time.Now().UnixNano())
	x := rand.Float32()

	if x < 0.5 {
		fmt.Println("head")
	} else {
		fmt.Println("tail")
	}

	var finger int = 1
	switch finger {
	case 0:
		fmt.Println("Thumb")
	case 1:
		fmt.Println("Index")
	case 2:
		fmt.Println("Middle")
	case 3:
		fmt.Println("Ring")
	case 4:
		fmt.Println("Pinkie")
	default:
		fmt.Println("Not a valid finger")
	}

	rand.Seed(time.Now().UnixNano())
	_x := rand.Float32()

	switch {
	case _x < 0.25:
		fmt.Println("Q1")
	case _x < 0.5:
		fmt.Println("Q2")
	case _x < 0.75:
		fmt.Println("Q3")
	default:
		fmt.Println("Q4")
	}

	y := 5
	counter := y

	for counter > 0 {
		fmt.Println(counter)
		counter--
	}

	for i := 0; i < y; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for {
		if y%2 == 0 {
			fmt.Printf("%d is odd\n", y)
			y++
			continue
		}
		break
	}

	for {
		fmt.Println("print foreveer")
		break
	}

	defer CloseMsg()

	rand.Seed(time.Now().UnixNano())
	id := rand.Int() % 6
	mosq, err := GetMusketeer(id)
	if err == nil {
		fmt.Printf("[%d] %s\n", id, mosq)
	}

	defer fmt.Println("Certainly closed !!!")
}

func sum(a int, b int) int {
	return a + b
}

func ops(a int, b int) (int, int) {
	return a + b, a - b
}

func sum_all(nums ...int) int {
	total := 0
	for _, a := range nums {
		total = total + a
	}

	return total
}

func multiply(a int, b int) int {
	return a * b
}

func doit(a int, b int, operator func(int, int) int) int {
	return operator(a, b)
}

func accumulator(increment int) func() int {
	i := 0
	return func() int {
		i = i + increment
		return i
	}
}

func zero_value(a int) {
	a = 0
}

func zero_point(a *int) {
	*a = 0
}

var musketeers = []string{
	"Athos", "Porthos", "Aramis", "D’Artagnan",
}

func GetMusketeer(id int) (string, error) {
	if id < 0 || id >= len(musketeers) {
		return "", errors.New(
			fmt.Sprintf("Invalid id [%d]", id))
	}

	return musketeers[id], nil
}

func CloseMsg() {
	fmt.Println("Closed!!!")
}

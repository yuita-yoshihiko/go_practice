package main

import "fmt"

func plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
		return
}

func Returnfunc () func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	r := 0
	return func() int {
		r++
		return r
	}
}

func main() {
  i := plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()



	f := func(x, y int) int {
		return x + y
	}
	k := f(1, 2)
	fmt.Println(k)

	k2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(k2)



  f2 := Returnfunc()
	f2()

	CallFunction(func() {
		fmt.Println("I'm a function")
	})

	f3 := Later()
	fmt.Println(f3("Hello"))
	fmt.Println(f3("My"))
	fmt.Println(f3("Name"))
	fmt.Println(f3("Is"))
	fmt.Println(f3("Golang"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}

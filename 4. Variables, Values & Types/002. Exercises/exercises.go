package main

import "fmt"

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
}

func ex01() {

	x := 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y := "James Bond"
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	z := true
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(x, y, z)

}

var x int
var y string
var z bool

func ex02() {

	fmt.Println(x, y, z)

}

var a int = 12
var b string = "Test"
var c bool = true

func ex03() {

	fmt.Printf("%v\t%v\t%v\t\n", a, b, c)

	fmt.Printf("%T\t%T\t%T\t\n", a, b, c)

}

type xablau int

var n xablau

func ex04() {
	n = 145
	fmt.Printf("%v\t%T\n", n, n)
}

var l int

func ex05() {
	l = int(n)
	fmt.Println(l)
}

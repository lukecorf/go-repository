package main

import "fmt"

//Declare a variable. you can declarate outside of the function
var y = 22
var z int
var w = 12

//Creating a type
type xablau int

var a xablau

func main() {
	fmt.Println("Hello everyone!")

	foo()

	testFor()

	bar()

	testPrintln()

	declareValues()

	printXablau()

	convertTypes()
}

func testFor() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}

func foo() {
	fmt.Println("I`m Foo!")
}

func bar() {
	fmt.Println("And then we exited")
}

func testPrintln() {
	n, err := fmt.Println("test")
	fmt.Println(n)
	fmt.Println(err)
	//To ignore a return you just need use the _
	t, _ := fmt.Println("test")
	fmt.Println(t)
}

func printXablau() {
	a = 21
	fmt.Println(a)
	//Print the type of the variable
	fmt.Printf("%T\n", a)

}

func convertTypes() {
	//Convert the xablau type to int
	z = int(a)
	fmt.Println(z)
}

func declareValues() {
	// the : indicate a declaration. And you can declarate this just inside of the function
	x := 42
	fmt.Println(x)
	//After you declare you can set values to your variable using just =
	x = 32
	fmt.Println(x)
	//Using a global variable
	fmt.Println(y)
	//Print a global variable of int type
	fmt.Println(z)
}

// control flow:
// (1) SEQUENCES
// (2) LOOP
// (3) CONDITIONAL

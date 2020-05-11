package main

import "fmt"

func main() {
	fmt.Println("Hello everyone!")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I`m Foo!")
}

func bar() {
	fmt.Println("And then we exited")
}

// control flow:
// (1) SEQUENCES
// (2) LOOP
// (3) CONDITIONAL

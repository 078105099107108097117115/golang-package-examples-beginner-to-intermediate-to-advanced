package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	example()
	fmt.Println()
	fmt.Println()
	secondExample()
	fmt.Println()
	func1("Golang,Erlang and Rust baby!")
	fmt.Println()
	x := new(int)
	fmt.Println(x)
	c := complex(13.2, 11)
	fmt.Printf("complex no c:-> %v\n", c)
	fmt.Println()
}
func example() {
	i := 0
	defer fmt.Println(i) // Prints 0
	i++
	// defer fmt.Println(i)

	return
}
func secondExample() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//Defer can also be used to trace where we/you are in the code.
// ........................
// ........................

// Can also be used to defer to log parameter and return values from within the function
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(\"%q\") = %d, %v", s, n, err)
	}()
	return 120, io.ErrUnexpectedEOF
}

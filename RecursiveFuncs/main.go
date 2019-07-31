package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonnaci(10))
	fmt.Println(fibonnaci(1))
	fmt.Println()
	fmt.Println()
	out1, out2 := even(2333001)
	fmt.Printf(" For 12: \n%t ; %s", out1, out2)
	var someValue int
	someValue = -2
	factResult1, factResult2 := factorial(someValue)
	fmt.Println()
	fmt.Println()
	fmt.Printf("factorial (%d) -> %d,%v\n", someValue, factResult1, factResult2)
}
func fibonnaci(input int) (result int) {
	if input <= 1 {
		result = 1
	} else {
		result = fibonnaci(input-1) + fibonnaci(input-2)
	}
	return
}

// Mutually recursive functions:
func even(nr int) (res bool, s string) {
	defer func() {
		if res == true {
			s = fmt.Sprintf("The input no %d is even!", nr)
		} else {
			s = fmt.Sprintf("The input no %d is not even, it's odd!", nr)
		}
	}()
	if nr == 0 {
		return true, s
	}
	return odd(revSign(nr) - 1)
}
func odd(nr int) (res bool, s string) {
	defer func() {
		if res == true {
			s = fmt.Sprintf("The input no %d is odd!", nr)
		} else {
			s = fmt.Sprintf("The input no %d is not odd, it's even!", nr)
		}
	}()
	if nr == 0 {
		return false, s
	}
	return even(revSign(nr) - 1)
}
func revSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

// Factorial recursive fxn:
var val int

func factorial(n int) (res int, err error) {
	if n < 0 {
		res = 0
		err = fmt.Errorf("invalid input. Input less than 0")
	} else if n == 0 || n == 1 {
		res = 1
		err = nil
	} else {
		err = nil
		val, _ = factorial(n - 1)
		res = n * val
	}
	return
}

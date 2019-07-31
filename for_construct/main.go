package main

import "fmt"

func main() {
	str := "Go is a statically-typed language with strong concurrency ideal for networking!"
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("The character at position %d is %c.\n", ix, str[ix])
	}
	secStr := "Random Korean Letters: 새롭다기둥답답하다틀림없이" //Korean Non-ASCII characters
	// For - Range construct needed for Non-ASCII chars
	for pos, val := range secStr {
		fmt.Printf("The character at position/index %d is %c\n", pos, val)
	}
	fmt.Println()
	fmt.Printf("Index\tInt(Rune)\tRune\tChar\tBytes\n")
	for index, char := range secStr {
		fmt.Printf("%d\t%d\t\t%U\t%c\t%X\n", index, char, char, char, []byte(string(char)))
	}
	println()
	for i := 0; i < 5; i++ {
		var v int
		println(" ", v)
		v = 5
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		for j := 20; j > 0; j-- {
			if j == 10 {
				continue
			}
			if j < 4 {
				break
			}
			fmt.Println("i -> ", i, "; j -> ", j)

		}

	}
}

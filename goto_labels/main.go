package main

import "fmt"

func main() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is %d and j is %d\n", i, j)
		}

	}
	fmt.Println()
	a := 100
	b := 12
START:
	if a < 200 {
		goto TARGET
	} else {
		return
	}

TARGET:
	fmt.Printf("a is %d ; b is %d\n", a, b)
	a += b
	goto START

}

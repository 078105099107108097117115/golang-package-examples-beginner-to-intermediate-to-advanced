package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println(MinNum(2, 3, 1, 45, 78, 8, 80, 5))
	fmt.Println()
	fmt.Println("**********************************")
	fmt.Println()
	// variadicFunc('w', 's', 'q', 'r', 'r', '1', 'o', 'g')
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var strSl []string
	for _, val := range letters {
		out := fmt.Sprintf("%c", val)
		strSl = append(strSl, out)

	}
	fmt.Println(strSl, len(strSl))
}

//MinNum function below demonstrates the use of Variadic functions
func MinNum(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, val := range a {
		if val < min {
			min = val
		}
	}
	return min
}

// func variadicFunc(r ...rune) []string {
// 	variadicFunc3(r...)
// 	s := string(r)
// 	var inputVF2 []string
// 	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	for i:=0;i<len(r);i++{
// 		inputVF2[i] = s
// 		s = ""
// 	}
// 	variadicFunc2(inputVF2...)

// }

// func variadicFunc2(s ...string) {

// 	for i := 0; i < len(s); i++ {
// 		s[i] = s[i] + "abcde"
// 	}
// 	sort.Strings(s)
// 	fmt.Println(s)
// }
// func variadicFunc3(s ...rune) []string {
// 	max := s[0]
// 	var newSl []string
// 	for _, val := range s {
// 		if val > max {
// 			max = val
// 		}
// 		newSl = append(newSl, string(val))
// 	}
// 	fmt.Println(newSl)
// }

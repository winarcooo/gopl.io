package main

import "fmt"

func main() {
	// this variable is array literal
	a := [...]int{0, 1, 2, 3, 4, 5}

	// this b is slice literal
	// b := []int{6, 7, 8, 9, 10, 11}
	reverse(a[:])
	fmt.Println(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

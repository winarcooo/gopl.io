package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var a [3]int = [3]int{1, 2, 3}
	var b = [3]int{4, 5, 6}
	c := [4]int{9, 8, 7, 6}
	d := [...]int{4, 5, 6}

	for k, v := range c {
		fmt.Printf("key : %v  value : %v\n", k, v)
	}

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}

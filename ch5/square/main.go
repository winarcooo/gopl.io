package main

import (
	"fmt"
	"reflect"
)

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(reflect.TypeOf(f)) // func() int
	fmt.Println(f())
}

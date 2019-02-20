package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum(2))        //2
	fmt.Println(sum(2, 4, 5))  //11
	fmt.Println(sum(5, -2, 1)) //4

	values := []int{10, 12, 8}
	fmt.Println(sum(values...))
}

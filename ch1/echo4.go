package main

import (
	"fmt"
	"os"
)

func main() {
	for i, val := range os.Args[1:] {
		fmt.Printf("key:%v and value:%v\n", i, val)
	}
}

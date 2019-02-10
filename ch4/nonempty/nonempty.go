package main

import "fmt"

func nonemtpy(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonemtpy2(strings []string) []string {
	out := strings[:0] // zero-length slice of origins
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonemtpy2(data))
	fmt.Printf("%q\n", data)
}

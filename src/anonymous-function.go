//go:build ignore

package main

import "fmt"

func store() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func main() {
	s1 := store()
	s2 := store()
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2)) // s1: 1, s2: 2
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2)) // s1: 2, s2: 4
}

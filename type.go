package main

import "fmt"

func multiple(x, y string) (string, string) {
	return x, y
}

func main() {
	// x, y := 9.32, 12.77
	w1, w2 := "Hello", "Test"
	fmt.Println(multiple(w1, w2))
}

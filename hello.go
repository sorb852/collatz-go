package main

import "fmt"

func add(a float32, b float32) float32 {
	return a + b
}

func main() {
	var a float32
	var b float32
	fmt.Print("Please input 'a' and 'b': ")
	fmt.Scan(&a, &b)
	out := add(a, b)

	fmt.Printf("'a' is type %T and value %v\n", a, a)
	fmt.Printf("'b' is type %T and value %v\n", b, b)
	fmt.Printf("'add(a + b)' is type %T and value %v\n", out, out)
}

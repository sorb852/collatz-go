package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"path/filepath"
)

func apply_rule(x int) int {
	if x % 2 == 0 {
		return x / 2
	} else {
		return 3 * x + 1
	}
}

func create_chain(n int) []int {
	var out = []int{}
	for i := n; true; i = apply_rule(i) {
		out = append(out, i)
		if i == 1 { break }
	}
	return out
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "No argument given")
		os.Exit(1)
	}
	if args[0] == "-h" || args[0] == "--help" {
		fmt.Fprintf(os.Stderr, "Usage: %v [-h | --help | N]\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	}

	var n_str string = args[0]

	n, err := strconv.Atoi(n_str)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please input a valid string")
		os.Exit(1)
	}

	if n < 1 {
		fmt.Fprintln(os.Stderr, "N must be a natural number")
		os.Exit(1)
	}

	chain := create_chain(n)
	fmt.Println("Created chain:")
	fmt.Printf("[%v", chain[0])
	for i := 1; i < len(chain); i++ {
		fmt.Printf(", %v", chain[i])
	}
	fmt.Println("]")

	fmt.Printf("Element count: %v\n", len(chain))
	fmt.Printf("Peak of graph: %v\n", slices.Max(chain))
}

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"path/filepath"
)

// Applies the rule on the number and returns it.
func apply_rule(x int) int {
	if x % 2 == 0 {
		return x / 2
	} else {
		return 3 * x + 1
	}
}

// Creates the chain of numbers and taking `n` as its starting point
func create_chain(n int) []int {
	var out = []int{}
	// its lowkey clean tho right
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
	// ik this a little sloppy but hey, its a simple CLI so no body cares
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

	// print the chain
	// it its python esque so you can use it somewhere else
	fmt.Println("Created chain:")
	fmt.Printf("[%v", chain[0])
	for i := 1; i < len(chain); i++ {
		fmt.Printf(", %v", chain[i])
	}
	fmt.Println("]")

	// some stats
	fmt.Printf("Element count: %v\n", len(chain))
	fmt.Printf("Peak of graph: %v\n", slices.Max(chain))
}

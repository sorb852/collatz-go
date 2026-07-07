package processor

// Applies the rule on the number and returns it.
func ApplyRule(x int) int {
	if x%2 == 0 {
		return x / 2
	} else {
		return 3*x + 1
	}
}

// Creates the chain of numbers and taking `n` as its starting point
func CreateChain(n int) []int {
	var out = []int{}
	// its lowkey clean tho right
	for i := n; true; i = ApplyRule(i) {
		out = append(out, i)
		if i == 1 {
			break
		}
	}
	return out
}

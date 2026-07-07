package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"
)

// Applies the rule on the number and returns it.
func apply_rule(x int) int {
	if x%2 == 0 {
		return x / 2
	} else {
		return 3*x + 1
	}
}

// Creates the chain of numbers and taking `n` as its starting point
func create_chain(n int) []int {
	var out = []int{}
	// its lowkey clean tho right
	for i := n; true; i = apply_rule(i) {
		out = append(out, i)
		if i == 1 {
			break
		}
	}
	return out
}

// Error response to collatz conjecture request
type CollatzErrorResponse struct {
	Message string `json:"message"`
}
// Response to collatz conjecture request
type CollatzResponse struct {
	StartingNumber int   `json:"starting_number"`
	Length         int   `json:"length"`
	Peak           int   `json:"peak"`
	Chain          []int `json:"chain"`
}

// my favourite little conjecture
// i would never solve it obv but i think its so cool
func collatz_handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	n_str := params.Get("n")
    w.Header().Set("Content-Type", "application/json")

	if n_str == "" {
		res, _ := json.Marshal(CollatzErrorResponse{Message: "parameter n is missing"})
		http.Error(w, string(res), http.StatusBadRequest)
		return
	}

	n, parse_err := strconv.Atoi(n_str)
	if parse_err != nil || n < 1 {
		res, _ := json.Marshal(CollatzErrorResponse{Message: "parameter n is not a natural number"})
		http.Error(w, string(res), http.StatusBadRequest)
		return
	}

	chain := create_chain(n)
	response := CollatzResponse{
		StartingNumber: chain[0],
		Length:         len(chain),
		Peak:           slices.Max(chain),
		Chain:          chain,
	}

	json_parse_err := json.NewEncoder(w).Encode(response)
	if json_parse_err != nil {
        http.Error(w, json_parse_err.Error(), http.StatusInternalServerError)
        return
	}

    w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/collatz", collatz_handler)

    // Start that baby UP
	log.Fatal(http.ListenAndServe(":8080", nil))
}

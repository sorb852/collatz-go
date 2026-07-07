package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
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

// Parases the URI string to get the params
func get_params(uri string) map[string]string {
	split := strings.Split(uri, "?")
	if len(split) < 2 {
		return nil
	}
	paramstr := split[1]
	params_pair := strings.Split(paramstr, "&")
	params := map[string]string{}
	for _, pair := range params_pair {
		pair := strings.Split(pair, "=")
		if len(pair) < 2 {
			continue
		}
		params[pair[0]] = pair[1]
	}
	return params
}

type CollatzResponse struct {
	starting_number int
	length          int
	peak            int
	chain           []int
}

func collatz_handler(w http.ResponseWriter, r *http.Request) {
	params := get_params(r.RequestURI)
	n_str, exists := params["n"]
	if !exists {
		fmt.Fprint(w, "parameter n doesn't exist")
		return
	}

	n, parse_err := strconv.Atoi(n_str)
	if parse_err != nil || n < 1 {
		fmt.Fprint(w, "parameter n is not a natural number")
		return
	}

	chain := create_chain(n)

	response := CollatzResponse{chain[0], len(chain), slices.Max(chain), chain}

	response_str, json_parse_err := json.Marshal(response)
	if json_parse_err == nil {
		fmt.Fprint(w, string(response_str))
	}
}

func main() {
	http.HandleFunc("/collatz", collatz_handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

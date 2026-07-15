package routes

import (
	"collatz-go/processor"
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"
)

// Error response to collatz conjecture request
type CollatzErrorResponse struct {
	Message string `json:"message"`
}

// Response to collatz conjecture request
type CollatzResponse struct {
	StartingNumber int     `json:"starting_number"`
	Length         int     `json:"length"`
	Average        float32 `json:"average"`
	Peak           int     `json:"peak"`
	Chain          []int   `json:"chain"`
}

// my favourite little conjecture
// i would never solve it obv but i think its so cool
func CollatzHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request with method %s using %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)

	params := r.URL.Query()
	n_str := params.Get("n")
	w.Header().Set("Content-Type", "application/json")

	if n_str == "" {
		log.Printf("Missing parameter `n` from %s\n", r.RemoteAddr)
		res, _ := json.Marshal(CollatzErrorResponse{Message: "parameter n is missing"})
		http.Error(w, string(res), http.StatusBadRequest)
		return
	}

	n, parse_err := strconv.Atoi(n_str)
	if parse_err != nil || n < 1 {
		log.Printf("Non natural number `n` from %s\n", r.RemoteAddr)
		res, _ := json.Marshal(CollatzErrorResponse{Message: "parameter n is not a natural number"})
		http.Error(w, string(res), http.StatusBadRequest)
		return
	}

	log.Printf("Processing number %v for %s\n", n, r.RemoteAddr)
	chain := processor.CreateChain(n)
	var sum int = 0
	for _, v := range chain {
		sum += v
	}
	response := CollatzResponse{
		StartingNumber: chain[0],
		Length:         len(chain),
		Average:        float32(sum) / float32(len(chain)),
		Peak:           slices.Max(chain),
		Chain:          chain,
	}

	json_parse_err := json.NewEncoder(w).Encode(response)
	if json_parse_err != nil {
		log.Printf("JSON parse error (somehow) while sending to %s\n", r.RemoteAddr)
		http.Error(w, json_parse_err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully sent Response back to %s\n", r.RemoteAddr)
}

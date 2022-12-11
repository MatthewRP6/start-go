package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	HC = 0
)

type SendBack struct {
	Status string
	Value  int
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	TEST := []byte("OK\n")
	w.Write(TEST)
	HC++
}

func metricHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("called metric")
	res, err := json.Marshal(SendBack{Status: "OK", Value: 0})
	if err != nil {
		fmt.Println("error")

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func main() {
	fmt.Printf("starting server at port 8080\n")
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/metric", metricHandler)

	http.ListenAndServe(":8080", nil)

}

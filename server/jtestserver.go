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
	w.WriteHeader(http.StatusOK)
	HC++
}

func metricHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("called metric\n")
	res, err := json.Marshal(SendBack{Status: "OK", Value: 0}) //Returns JSON encoding of V
	if err != nil {
		fmt.Println("error")

	}
	json.NewEncoder(w).Encode(res)                     //moved up from line 34 . Creates an encoder writing to w + writes to W
	w.Header().Set("Content-Type", "application/json") //Adds  header + description to HeaderMap
	w.WriteHeader(http.StatusOK)                       //Sends http response header with (response code) http.StatusOK = constant int

}

func main() {
	fmt.Printf("starting server at port 8080\n")
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/metric", metricHandler)

	http.ListenAndServe(":8080", nil)

}

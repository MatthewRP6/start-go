package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	HC = 0
)

type SendBack struct {
	Status int
	Msg    string
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	HC++
	log.Println(HC)
	json.NewEncoder(w).Encode(SendBack{})

}

func metricHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("call to metrics ep")
	back := SendBack{Status: http.StatusOK, Msg: "How would you change this endpoint to pass it a value, and have it repeat it back to the client?"} // 2) cleaned this up
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(back) // <-- click into Encode method and see what it's actually doing to ...
	// 									to understand why it should look like this

}

func main() {
	fmt.Printf("starting server at port 8080\n")
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/metric", metricHandler)

	http.ListenAndServe(":8080", nil)

}

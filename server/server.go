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
	fmt.Printf("healthHandler running now")

}
func metricHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("metricHandler running now")
	log.Println("call to metrics ep")
	sendback := r.Header.Values("Send-Back")
	sendbackv := sendback[0]
	back := SendBack{Status: http.StatusOK, Msg: sendbackv}
	//back := SendBack{Status: http.StatusOK, Msg: "How would you change this endpoint to pass it a value, and have it repeat it back to the client?"} // 2) cleaned this up
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(back) // <-- click into Encode method and see what it's actually doing to ...
	// 									to understand why it should look like this
	//Encode writes the json value of back into the stream 'w' followed by a new line character
}

func main() {
	fmt.Printf("starting server at port 8080\n")
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/metric", metricHandler)

	http.ListenAndServe(":8080", nil)

}

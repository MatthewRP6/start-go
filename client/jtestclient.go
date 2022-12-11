package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type ExpectedVal struct {
	Status string
	Value  int
}

func main() {

	addr := os.Getenv("MP_PORT")

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	// resp, err := c.Get("http://localhost:" + port + userI)
	// if err != nil {
	// 	fmt.Printf("Error %s", err)
	// 	return
	// }

	//make a post request with body and read request body on server and do something.
	req, err := http.NewRequest("GET", addr, nil)

	if err != nil {
		fmt.Printf("error: %v \n", err)
		os.Exit(-1)
	}
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("http req to ep returned with error: %v\n", err)
		os.Exit(-1)
	}
	defer res.Body.Close()
	target := &ExpectedVal{}
	_ = json.NewDecoder(res.Body).Decode(target)

	fmt.Printf("FINAL:%v ", *target)
}

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

	//addr := os.Getenv("MP_PORT")

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	// resp, err := c.Get("http://localhost:" + port + userI)
	// if err != nil {
	// 	fmt.Printf("Error %s", err)
	// 	return
	// }

	//make a post request with body and read request body on server and do something.
	req, err := http.NewRequest("GET", "http://localhost:8080/metric", nil) //replaced middle var from addr

	if err != nil {
		fmt.Printf("error: %v \n", err)
		os.Exit(-1)
	}
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)

	//fmt.Printf("res: %v\n", res)

	//x := res.StatusCode
	//fmt.Println(x)

	//x := res.Header.Values("Content-Type")
	//fmt.Println(x) ??WHY DOES THIS PRINT "text/plain; char=utf-8" instead of "application/json"

	if err != nil {
		fmt.Printf("http req to ep returned with error: %v\n", err)
		os.Exit(-1)
	}
	defer res.Body.Close()
	target := &ExpectedVal{}
	_ = json.NewDecoder(res.Body).Decode(target)

	fmt.Println(target.Status) //Why does this print blank line instead of "OK"

	//fmt.Println(res.Status)

	//fmt.Println(res.Header)
	//fmt.Printf("FINAL:%s ", target.Status)

}

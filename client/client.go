package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"
)

type ExpectedVal struct {
	Status int
	Msg    string
}

func main() {
	//FIGURE OUT WHATS GOING ON HERE  VVVVVV

	var ep string
	flag.StringVar(&ep, "ep", "health", "usage: pick ep by name opts:[health, metric] ")
	flag.Parse()
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	// resp, err := c.Get("http://localhost:" + port + userI)
	// if err != nil {
	// 	fmt.Printf("Error %s", err)
	// 	return
	// }

	// useless but giving you ideas to play with
	url := ""
	switch ep {
	case "health":
		url = "http://localhost:8080/"
	case "metric":
		url = "http://localhost:8080/metric"

	}

	//make a post request with body and read request body on server and do something.
	req, err := http.NewRequest("GET", url, nil) //replaced middle var from addr
	if err != nil {
		log.Fatalf("Issue creating request struct: %v\n ", err)
	}
	//Accept header is used by HTTP clients to tell the server which type of content they expect/prefer as response.
	//Content-type can be used both by clients and servers to identify the format of the data -
	//															in their request (client) or response (server) and,
	//															therefore, help the other part interpret correctly the information.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Send-Back", "Return this value to the client")
	res, err := c.Do(req)

	//fmt.Printf("res: %v\n", res)

	//x := res.StatusCode
	//fmt.Println(x)

	//x := res.Header.Values("Content-Type")
	//fmt.Println(x) ??WHY DOES THIS PRINT "text/plain; char=utf-8" instead of "application/json"

	if err != nil {
		log.Fatalf("http req to ep returned with error: %v \n", err)
	}
	defer res.Body.Close()
	target := ExpectedVal{}
	err = json.NewDecoder(res.Body).Decode(&target) // <-------- get the error and see what it is.
	if err != nil {
		log.Fatalf("%v \n", err) //<---- started here 1) 2022/12/17 19:35:46 json: cannot unmarshal string into Go value of type main.ExpectedVal
		// <--- click into log.Fatalf and look what it does & why i clean up the fmts.
		//Changed fmts to log.Fatalf because log.Fatalf will  call fmt.Printf as well as call OS.Exit(1) which is what we want to happen when an error occurs
	}

	log.Printf("http res status code: %d, status from json body %d, message: %s, ", res.StatusCode, target.Status, target.Msg) // <- think about why this works(???) .... and let me know. specifically the first value
	//This Works because we have a variable "res" of type http.response which contains an integer variable called StatusCode. Similar for the second two values that are a variable "target" of our type ExpextedVal which contains the fields Status and Msg

	//Why does this print blank line instead of "OK"

	//fmt.Println(res.Status)

	//fmt.Println(res.Header)
	//fmt.Printf("FINAL:%s ", target.Status)

}

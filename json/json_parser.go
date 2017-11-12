package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

var p = fmt.Println

type ip struct {
	Origin string
}

func main() {
	response, err := grequests.Get("http://httpbin.org/ip", nil)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	resp := response.String()
	p(response.StatusCode == 200)
	p(resp)
	//spew.Dump(resp)
}

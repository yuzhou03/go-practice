package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func main() {
	resp, _ := http.Get("http://httpbin.org/ip")
}

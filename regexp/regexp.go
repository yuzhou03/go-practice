package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var p = fmt.Println

func main() {

	r, _ := regexp.Compile("p([a-z]+)ch")

	p(r.MatchString("peach"))

	p(r.FindString("peach punch"))

	p(r.FindStringIndex("peach punch"))

	p(r.FindStringSubmatchIndex("peach punch"))

	p(r.FindAllString("peach punch", -1))

	p(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")

	out := r.ReplaceAllFunc(in, bytes.ToUpper)

	p(string(out))
}

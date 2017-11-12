package main

import "fmt"

var p = fmt.Println

func main() {
	str := "abc" + "123"
	p(str)
	p(string(str[3]))
	p(demo(0))
}

func demo(x int) int {
	var ret int
	if x == 0 {
		ret = 5
	} else {
		ret = x
	}
	return ret
}

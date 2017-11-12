package main

import "fmt"

var p = fmt.Println

func MyPrint(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			p(arg, ": int value")
		case string:
			p(arg, " : string value")
		case int64:
			p(arg, ": int64 value")
		default:
			p("others")
		}
	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 100
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrint(v1, v2, v3, v4)
}

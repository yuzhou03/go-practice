package main

import (
	"fmt"
	"os"
	"strconv"
	"git.tvblack.com/pratice/calc/simplemath"
)

var p = fmt.Println
var Usage = func() {
	p("Usage: calc command [arguments] ...")
	p("\nThe commands are:add\t")
}

func main() {
	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[0] {
	case "add":
		if len(args) != 3 {
			p("usage:calc add <int1> <int2>")
			return
		}

		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			p("usage:calc add <int1> <int2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		p("result: ", ret)
	case "sqrt":
		if len(args) != 2 {
			p("Usage: calc sqrt <int>")
			return
		}

		v, err := strconv.Atoi(args[1])
		if err != nil {
			p("Usage: calc sqrt <int>")
			return
		}
		ret := simplemath.Sqrt(float64(v))
		p("result: ", ret)
	default:
		p("default")
		Usage()

	}


}

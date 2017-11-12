package main

import "fmt"

var p = fmt.Println

func sum(nums ...int) (sum int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	p(sum(1, 2, 3))
	p(sum(1, 2, 3000))
}

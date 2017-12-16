package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + ", " + p.last
}

//对象的类似于java的toSting()方法
func (p person) String() string {
	return fmt.Sprintf("Name:%s, Age:%d", p.fullName(), p.age)
}

var p = fmt.Println

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Tiger", "Woods", 28}
	p(p1)
	p(p2)
}

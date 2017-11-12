package main

import (
	"fmt"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

var p = fmt.Println

func (s *Student) Show() {
	p("show student:")
	p("\tName:\t", s.Name)
	p("\tAge:\t", s.Age)
	p("\tGuake\t", s.Guake)
	p("\tPrice\t", s.Price)
	p("\tClasses\t", s.Classes)
	for _, class := range s.Classes {
		fmt.Printf("%s", class)
	}
	p("")
}

func main() {
	student := &Student{
		"Xiao Ming",
		18,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
	p("before JSON encoding:")
	student.Show()

	b, err := json.Marshal(student)
	if err != nil {
		p("encoding failed")
	} else {
		p("encoded data:")
		p(b)
		p(string(b))
	}

	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))

	strData := <-ch
	p("---------")

	stu_bin := &Student{}

	stu_bin.Show()
	err = json.Unmarshal([]byte(strData), &stu_bin)

	if err != nil {
		p("Unmarshal failed")
	} else {
		p("Unmarshal success")
		stu_bin.Show()
	}
	p(strings.Repeat("=", 100))
	//spew.Dump(stu_bin)
	spew.Printf("%#v", student)
}

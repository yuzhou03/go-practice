package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
	"git.tvblack.com/pratice/sorter/algorithms/quicksort"
	"git.tvblack.com/pratice/sorter/algorithms/bubblesort"
)

var infile *string = flag.String("i", "infile", "Input File")
var outfile *string = flag.String("o", "outfile", "Output File")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")
var p = fmt.Println

func main() {

	flag.Parse()

	if infile != nil {
		p("infile:", *infile, ", outfile:", *outfile, ", algorithm:", *algorithm)
		values, err := readValues(*infile)
		if err == nil {
			p("Read Values", values)
		} else {
			p(err)
		}
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "quicksort":
			quicksort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			p("Sorting algorithm", *algorithm, " UNSUPPORTED")
		}

		t2 := time.Now()
		p("Sorting duration: ", t2.Sub(t1), " seconds.")
		writeValues(values, *outfile)
	} else {
		p(err)
	}
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		p("Failed to open the input file", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			p("A too long line, seems unexpected")
			return
		}

		str := string(line) //转换字符数组为字符串
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return

}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)

	if err != nil {
		p("failed to create the output file.", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

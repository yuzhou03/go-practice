package main

import (
	"os"
	"io"
	"flag"
	"fmt"
)

func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)

	if err != nil {
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)

}

var p = fmt.Println

func main() {
	srcFile := flag.String("src", "src", "source file")
	destFile := flag.String("dst", "dst", "dest file")
	p("src:", srcFile)
	p("dst:", destFile)
}

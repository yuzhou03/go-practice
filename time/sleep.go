package main

import (
	"fmt"
	"time"
)

func main() {
	//print("开始...\n")
	//time.Sleep(time.Second * 1)
	//print("结束")

	fmt.Println(NowStr())
}

func NowStr() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}

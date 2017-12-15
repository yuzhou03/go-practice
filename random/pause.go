package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	FORUM = iota
	THREAD
	IMAGE
)

func Pause(fetchType int) {

	var sleepSeconds time.Duration = 1
	var msg string = ""

	if fetchType == FORUM {
		//pausePerForum := config.Config.Fetch.Pause.PausePerForum
		pausePerForum := 2
		sleepSeconds = time.Duration(rand.Intn(pausePerForum) + 1)
		msg = fmt.Sprintf("等待处理下一个forum，Sleeping for %v seonds...\n", pausePerForum)
		fmt.Println(msg)
		time.Sleep(time.Second * sleepSeconds)
		fmt.Println("Sleep Done")

	} else if fetchType == THREAD {
		//pausePerThread := config.Config.Fetch.Pause.PausePerThread
		pausePerThread := 2
		sleepSeconds = time.Duration(rand.Intn(pausePerThread) + 1)
		msg = fmt.Sprintf("等待处理下一个Thread，Sleeping for %v seonds...\n", pausePerThread)
		fmt.Println(msg)
		time.Sleep(time.Second * sleepSeconds)
		fmt.Println("Sleep Done")
	} else if fetchType == IMAGE {
		//pausePerImage := config.Config.Fetch.Pause.PausePerImage
		pausePerImage := 3000
		milliSecond := time.Duration(rand.Intn(pausePerImage) + 1000)
		msg = fmt.Sprintf("等待处理下一个Image，Sleeping for %v milliSecond...\n", pausePerImage)
		fmt.Println(msg)
		time.Sleep(time.Millisecond * milliSecond)
		fmt.Println("Sleep Done")
	}

}

func main() {
	Pause(FORUM)
	Pause(THREAD)
	Pause(IMAGE)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		randN := rand.Intn(5)
		sleepInSeconds := 1 + randN
		fmt.Printf("Random sleep %v seconds.\n", sleepInSeconds)
		time.Sleep(time.Microsecond * time.Duration(sleepInSeconds)) //产生0-100的随机整数
	}

	//fmt.Println(rand.Float64()) //产生0.0-1.0的随机浮点点
	//
	//s1 := rand.NewSource(42) //用指定值创建一个随机数种子
	//r1 := rand.New(s1)
	//fmt.Print(r1.Intn(100), ",")
	//fmt.Print(r1.Intn(100))
	//fmt.Println()
	//
	//s2 := rand.NewSource(42) //同前面一样的种子
	//r2 := rand.New(s2)
	//fmt.Print(r2.Intn(100), ",")
	//fmt.Print(r2.Intn(100))
	//fmt.Println()

}

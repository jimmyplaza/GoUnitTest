package main

import (
	"fmt"
	"time"
)

func main() {
	go Hello()

	time.Sleep(time.Second * 1) // 暫停一秒鐘
}

func Hello() {
	fmt.Println("Hello")
}

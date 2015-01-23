package main

import "fmt"

func main() {

	//	runtime.GOMAXPROCS(4)

	go A()

	for {
		fmt.Println("B")
	}
}

func A() {
	for {
		fmt.Println("A")
	}
}

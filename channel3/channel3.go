package main

import "fmt"

func pingpong(ch chan int) {
	n := <-ch
	fmt.Println("Received %d", n)
	ch <- n
}

func main() {
	ch := make(chan int)
	go pingpong(ch)
	ch <- 42
	<-ch

}

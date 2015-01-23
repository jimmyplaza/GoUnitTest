package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(50 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(150 * time.Millisecond)
	ticker.Stop()

}

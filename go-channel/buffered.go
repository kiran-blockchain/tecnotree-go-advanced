package main

import (
	"fmt"
	//"time"
)

func main() {
	ch := make(chan int, 5) // 5 is the cap of a channel

	for i := 0; i < 5; i++ {
		ch <- i
	}
	//time.Sleep(2*time.Second)
	
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

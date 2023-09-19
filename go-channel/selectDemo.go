package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	x := 12
	go func(x int) {
		for i := 0; i < x; i++ {
			fmt.Println(<-ch) // read the data from the channel
		}
		quit <- false
	}(x)
	f(ch, quit)

}

func f(ch chan int, quit chan bool) {
	a, b := 0, 1
	for {
		select {
		case ch <- a: //write to the channel
			a, b = b, a+b
		case <-quit:
			fmt.Println("Completed")
			return
		}
	}
}

package main

import (
	"fmt"
	"sync"
	//"time"
)

func first(w *sync.WaitGroup) {
	//w.Add(1)
	defer w.Done()
	fmt.Println("I am the first function")
}
func second(w *sync.WaitGroup) {
	//w.Add(1)
	defer w.Done()
	fmt.Println("I am the second function")
}
func third(w *sync.WaitGroup) {
	//w.Add(1)
	defer w.Done()
	fmt.Println("I am the third function")
}

func main() {
	var wg sync.WaitGroup
		wg.Add(3)
		go first(&wg)
		go second(&wg)
		go third(&wg)
	fmt.Println("I am main")
	wg.Wait()
}

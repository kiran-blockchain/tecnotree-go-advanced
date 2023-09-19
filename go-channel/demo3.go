package main

import (
	"fmt"
	"sync"
)


var (
	msg string
	wg sync.WaitGroup
	mu sync.Mutex
) 
func first(){
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	msg ="Hello First"
	fmt.Println("i am in the first",msg)
	
}
func second(){
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	msg ="Hello Second"
	fmt.Println("i am in the second",msg)

}

func main(){
	msg="Welcome to concurrency"
	wg.Add(2)
	go first()
	mu.Lock()
	fmt.Println(msg)
	mu.Unlock()
	go second()
	wg.Wait()
	
	fmt.Println(msg)

}

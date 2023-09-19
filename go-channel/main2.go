package main

import (
	"fmt"
	//"time"
)

func main(){
 gen := channelGenerator()
 receiver(gen)
}

func receiver(c <-chan int){
  for i := range c{
	fmt.Println(i)
  }
}

func channelGenerator() <-chan int{
  c:= make(chan int)
    go func(){
		for i:=0;i<100;i++{
			fmt.Println("generating",i)
			c <-i
			//time.Sleep(1*time.Second)
		}
		close(c)
	}()
	return c
}
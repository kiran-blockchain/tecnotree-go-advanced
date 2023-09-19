package main

import (
	"fmt"
	//"time"
)

func main(){
	mesg := make(chan string)// channel creation
	 go greetings(mesg)
	reply:= <- mesg// main func acting as receiver of data from the greetigs
	reply2:=<- mesg
	reply3:= <-mesg
	_,ok:=<-mesg
	 if ok {
		fmt.Println("Channel is open")
	 }else{
		fmt.Println("Channel is closed")
	 }
	// _,ok := <-
	//reply3:=<- mesg
	//time.Sleep(1*time.Second)
	fmt.Println(reply)
	fmt.Println(reply2)
	fmt.Println(reply3)
}

func greetings(ch chan string){
	fmt.Println("Welcome to channel programming")
	ch <- "How are you"// sending data to the channel
	ch <- "How are you one"
	ch <- "How are you two"
	close(ch)
	ch <- "How are you Three"
	fmt.Println("Greetings completed")
}
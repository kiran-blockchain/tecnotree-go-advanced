package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct {
}

var SingleInstance *Singleton

func GetInstance() *Singleton {
	lock.Lock()
	defer lock.Unlock()
	if SingleInstance == nil {
		fmt.Println("Creating the instance")
		SingleInstance = &Singleton{}
	} else {
		fmt.Println("Instance already created")
	}
	return SingleInstance
}
func main(){
	for i:=0;i<10;i++{
		go GetInstance()
	}
	fmt.Scanln()
}

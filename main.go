package main

import (
	"fmt"
)
func init(){
	fmt.Println("I am init 1 \n")
}
func init(){
	fmt.Printf("I am init 2 \n")
}
func main() {
	fmt.Println("Welcome to go ")
}

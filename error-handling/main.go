package main

import (
	"fmt"

	"github.com/kiran-blockchain/error-handling/fileio"
)

func closeConnection(){
	fmt.Println("Closing the file")
}

func main(){
  
//   fmt.Println("I am opening the file")
//   defer closeConnection()
//   fmt.Println("File reading in progress")
fileio.CreateFile("demo.txt","Welcome to golang file io operations")
fileio.ReadFile("demo.txtt")
fmt.Println("Main is completed")

}
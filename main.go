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

func forRange(){
	numbers := []int{1,34,5,7,6}
     
	for index,item := range numbers{
		 fmt.Println("item - ",item)
		 fmt.Println("index = ",index)
	}

}

func slices (){
	 marks := [8]int{40,50,60,69,8,90,10,2}
	 s := marks[1:4]
	 fmt.Println(s)
	 fmt.Println(len(s))
	 fmt.Println(cap(s))
	//  for _,item := range marks{
	// 	fmt.Println(item)
	//  }
	//  for _,item1 := range s{
	// 	fmt.Println(item1)
	//  }
}


func main() {
	item, price := "Mobile", 2000
	fmt.Println("Welcome to go: ",item,price)
	if price>=2000 {
		//do this
		fmt.Println("price is greater than 2000\n")
	}
	//forRange()
	slices()
}

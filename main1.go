package main

import (
	"fmt"
)

func forRange() {
	numbers := []int{1, 34, 5, 7, 6}

	for index, item := range numbers {
		fmt.Println("item - ", item)
		fmt.Println("index = ", index)
	}

}

func Slices() {
	marks := [8]int{40, 50, 60, 69, 8, 90, 10, 2}
	numbers := []int{1, 34, 5, 7, 6}
	s := marks[0:4]
	fmt.Println(s)
	fmt.Println(len(s))
	//  fmt.Println(cap(s))
	s = append(s, numbers...)
	//s = append(s,4)
	fmt.Println(s)
	fmt.Println(cap(s))

	fmt.Println(s)
	//  for _,item := range marks{
	// 	fmt.Println(item)
	//  }
	//  for _,item1 := range s{
	// 	fmt.Println(item1)
	//  }

	//x:= []int{}
	x := make([]int, 6, 6)
	fmt.Println(x)

}

// func main() {
// 	item, price := "Mobile", 2000
// 	fmt.Println("Welcome to go: ",item,price)
// 	if price>=2000 {
// 		//do this
// 		fmt.Println("price is greater than 2000\n")
// 	}
// 	//forRange()
// 	slices()
// }

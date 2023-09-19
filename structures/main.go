package main

import (
	"fmt"

	"github.com/kiran-blockchain/structures/lists"
)

func main() {
	// r:= lists.Rectangle{Length: 10, Width:30 }
	// // var r lists.Rectangle
	// // r.Length=10
	// // r.Width=30
	// r:= new (lists.Rectangle)
	// r.Length=10
	//
	//	r.Width=30
	//
	// // r:=&lists.Rectangle{}
	// // (*r).Length =10
	// // (*r).Width=40
	// fmt.Println(r.Length)
	// fmt.Println(r.Width)
	// fmt.Println(r.Area())
	var list = lists.List{}
	list.Add("Kiran")
	list.Add("Matti")
	list.Add("john")
	PrintList(&list)
}
func PrintList(l *lists.List) {
	current := l.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
	fmt.Println("i am done")
}

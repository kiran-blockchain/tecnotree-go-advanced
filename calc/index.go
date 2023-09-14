package calc

import "fmt"

func  Add(a int,b int) (int){
	return a+b
}

func  Sub(a int,b int) (int){
	data:=Mul(a,b)
	fmt.Println(data)
	return a-b
}
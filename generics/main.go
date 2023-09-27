package main

import (
	"fmt"
	"generics/stack"
)

func main() {
    // Create a stack for integers.
    intStack := stack.NewStack[int]()
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)

    // Create a stack for strings.
    stringStack := stack.NewStack[string]()
    stringStack.Push("apple")
    stringStack.Push("banana")
    stringStack.Push("cherry")

    // Pop elements from the integer stack.
    fmt.Println("Popped from intStack:", 
	intStack.Pop()) // 3
    fmt.Println("Popped from intStack:",
	 intStack.Pop()) // 2

    // Pop elements from the string stack.
    fmt.Println("Popped from stringStack:", 
	stringStack.Pop()) // "cherry"
    fmt.Println("Popped from stringStack:",
	 stringStack.Pop()) // "banana"
}
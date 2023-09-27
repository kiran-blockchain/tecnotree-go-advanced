package stack

// Define a generic stack type with a type parameter T.
type Stack[T any] struct {
    elements []T
}
// Push adds an element to the stack.
func (s *Stack[T]) Push(item T) {
    s.elements = append(s.elements, item)
}
// Pop removes and returns the top element from the stack.
func (s *Stack[T]) Pop() T {
    if len(s.elements) == 0 {
        panic("Stack is empty")
    }
    top := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return top
}
// NewStack creates a new instance of the generic stack.
func NewStack[T any]() *Stack[T] {
    return &Stack[T]{}
}


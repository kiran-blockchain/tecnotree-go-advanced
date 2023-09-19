package lists

type Node struct {
	Data string
	Next *Node
}
type List struct {
	Head *Node
}

func (l *List) Add(value string) {
	newNode := &Node{Data: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode

}

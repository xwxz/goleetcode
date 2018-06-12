package structure

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
	Size uint
}

func (list *List) Init() {
	list.Size = 0
	list.Head = nil
	list.Tail = nil
}

func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}
	node.Next = nil
	if list.Size == 0 {
		list.Head = node
	} else {
		oldNode := list.Tail
		oldNode.Next = node
	}
	list.Tail = node
	list.Size++
	return true
}

func (list List) Print() {
	if list.Size != 0 {
		p := list.Head
		fmt.Printf("[")
		for p != nil {
			fmt.Printf("%d", p.Data)
			p = p.Next
			if p != nil {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("]\n")
	}
}
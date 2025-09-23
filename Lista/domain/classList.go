package domains

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

type LinkedList struct {
	Count int
	First *Node
}

func (l *LinkedList) IsEmpty() bool {
	if l.First == nil || l.Count <=0{
		return true
	}else{
		return false
	}
}

func (l *LinkedList) Add(value string, index int) {
	newNode := &Node{Value: value}
	if  index > l.Count || index < 0 {
		fmt.Println("não foi possível adicionar o elemento a lista")
		return
	}
	if l.IsEmpty() || index <= 0 {
		newNode.Next = l.First
		l.First = newNode
	} else {
		changeNode := l.First
		for i := 0; i < index-1 && changeNode != nil; i++ {
			changeNode = changeNode.Next
		}
		newNode.Next = changeNode.Next
		changeNode.Next = newNode
	}
	l.Count++
}

func (l *LinkedList) Show() {
	if l.Count == 0 {
		println("Lista Vazia")
		return
	}
	currentNode := l.First
	for currentNode != nil {
		println("Node =", currentNode.Value)
		currentNode = currentNode.Next
	}
}

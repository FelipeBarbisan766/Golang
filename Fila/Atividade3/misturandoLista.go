package main

import (
	"fmt"

	"Fila/domains"
)

func main() {
	var lista1 domains.Queue
	lista1.Enqueue("1")
	lista1.Enqueue("2")
	lista1.Enqueue("3")
	lista1.Enqueue("4")
	lista1.Enqueue("5")
	
	var lista2 domains.Queue
	lista2.Enqueue("A")
	lista2.Enqueue("B")
	lista2.Enqueue("C")

	var lista3 = Mistura(&lista1, &lista2)
	fmt.Println("Lista misturada:")
	lista3.Show()
}
func Mistura(q1, q2 *domains.Queue) *domains.Queue {
	var lista3 domains.Queue

	// Example logic: alternate enqueue from q1 and q2
	node1 := q1.First
	node2 := q2.First
	for node1 != nil {
		if node1 != nil {
			lista3.Enqueue(node1.Value)
			node1 = node1.Next
		}
	}
	for node2 != nil {
		if node2 != nil {
			lista3.Enqueue(node2.Value)
			node2 = node2.Next
		}
	}

	return &lista3
}

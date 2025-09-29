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

	fmt.Println("Fila original:")
	lista1.Show()

	lista2 := Invert(&lista1)
	fmt.Println("Fila invertida:")
	lista2.Show()
}
func Invert(q *domains.Queue) domains.Queue {
	var prev *domains.Node = nil
    current := q.First
    q.Last = q.First

    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }

    q.First = prev
	return *q
	
}
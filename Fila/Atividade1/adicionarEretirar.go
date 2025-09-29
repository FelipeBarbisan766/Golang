package main

import (
	"fmt"

	"Fila/domains"
)

func main() {
	var lista domains.Queue
	lista.Enqueue("10")
	lista.Enqueue("20")
	lista.Enqueue("30")
	lista.Enqueue("40")
	lista.Show()
	lista.Dequeue()
	lista.Dequeue()
	lista.Dequeue()
	if !lista.IsEmpty() {
		fmt.Println("Fila depois da retirada:")
		lista.Show()
	} else {
		fmt.Println("Fila está vazia")
	}
}

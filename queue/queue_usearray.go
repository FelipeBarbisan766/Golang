package main

import (
	"fmt"

	"prof.ronaldo/domains"
)

func main() {
	var qArr domains.QueueArray
	qArr.Enqueue(1)
	qArr.Enqueue(2)
	qArr.Enqueue(3)
	qArr.Show()
	qArr.Dequeue()
	qArr.Dequeue()
	qArr.Dequeue()
	if !qArr.IsEmpty() {
		fmt.Println("Fila depois da retirada:")
		qArr.Show()
	} else {
		fmt.Println("Fila está vazia")
	}
}

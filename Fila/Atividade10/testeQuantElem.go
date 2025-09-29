package main

import (
	"Fila/domains"
	"fmt"
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

	verifyQuantEntitys(&lista1,&lista2)
}
func verifyQuantEntitys(q1,q2 *domains.Queue) {
	fmt.Println("Verificacao de quantidade de elementos")
	if q1.Count > q2.Count {
		fmt.Print("Lista1 Maior que Lista2")
	}
	if q1.Count == q2.Count {
		fmt.Print("Lista1 Igual Lista2")
	}
	if q1.Count < q2.Count {
		fmt.Print("Lista1 Menor que Lista2")
	}
	
}
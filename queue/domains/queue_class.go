package domains

import "fmt"

// Representação para os elementos (nodes) que compõe a Fila
type Node struct {
	Value string
	Next  *Node
}

// Gerenciamento dos mecanismos de controle da Fila como um todo, sua quantidade de elementos, quem
// Qual o primeiro elemento da Fila, inserir, excluir...
type Queue struct {
	Count int
	First *Node
	Last  *Node
}

// Verifica se a Fila está vazia
func (q *Queue) IsEmpty() bool {
	if q == nil || q.Count <= 0 {
		return true
	} else {
		return false
	}
}

// Inserir novo item na Fila
func (q *Queue) Enqueue(data string) {
	var newNode Node

	newNode.Value = data

	if q.IsEmpty() {
		q.First = &newNode
	} else {
		q.Last.Next = &newNode
	}
	q.Last = &newNode
	q.Count++
}

// remover item da pilha
func (q *Queue) Dequeue() Node {
	var nodeRemoved Node
	if q.IsEmpty() {
		return nodeRemoved
	}
	q.Count--
	nodeRemoved = *q.First
	q.First = nodeRemoved.Next
	return nodeRemoved
}

func (q *Queue) Show() {
	if q.IsEmpty() {
		fmt.Println("Fila vazia!")
	} else {
		var nodeExibir = q.First
		for {
			if nodeExibir == nil {
				break
			}
			fmt.Println("Node = ", nodeExibir.Value)
			nodeExibir = nodeExibir.Next
		}
	}
}

func (q *Queue) Size() int {
	return q.Count
}

func (q *Queue) Begining() *Node {
	if q.IsEmpty() {
		return nil
	} else {
		return q.First
	}
}

func (q *Queue) Ending() *Node {
	if q.IsEmpty() {
		return nil
	} else {
		return q.Last
	}
}

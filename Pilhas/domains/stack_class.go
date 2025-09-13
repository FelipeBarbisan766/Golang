package domains

import "fmt"

// Representação para os elementos (nodes) que compõe a pilha
type Node struct {
	Value    string
	Previous *Node
}

// Gerenciamento dos mecanismos de controle da pilha como um todo, sua quantidade de elementos, quem
// está no topo, inserir novo node, excluir....
type Stack struct {
	Count int
	Top   *Node
}

// Inserir novo item na pilha
func (s *Stack) Push(data string) {
	s.Count++
	var newNode Node
	newNode.Value = data
	newNode.Previous = s.Top
	s.Top = &newNode
}

// remover item da pilha
func (s *Stack) Pop() Node {
	var nodeAux Node
	if s.IsEmpty() {
		return nodeAux
	}
	s.Count--
	nodeAux = *s.Top
	s.Top = nodeAux.Previous
	return nodeAux
}

func (s *Stack) IsEmpty() bool {
	if s == nil || s.Count <= 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Show() {
	if s.IsEmpty() {
		fmt.Println("Pilha vazia!")
	} else {
		var nodeExibir = s.Top
		for {
			if nodeExibir == nil {
				break
			}
			fmt.Println("Node = ", nodeExibir.Value)
			nodeExibir = nodeExibir.Previous
		}
	}
}

func (s *Stack) Size() int {
	return s.Count
}

package main

import (
	"Fila/domains"
)

func main() {

	// 4 - (WILKER, 2019) Considere a estrutura de dados fila, do tipo FIFO. Entidades são inseridas nessa estrutura com a operação enqueue( ) e removidas com a operação dequeue( ). A opção a seguir que mostra o conteúdo ordenado da fila após asequênciade operações: enqueue(8), enqueue(7), enqueue(5), enqueue(2), dequeue(),enqueue(8), enqueue(7), dequeue( ), enqueue(5), enqueue(2), dequeue( ), dequeue()é:
	
	// c) 8752 √

	var lista domains.Queue
	lista.Enqueue("8")
	lista.Enqueue("7")
	lista.Enqueue("5")
	lista.Enqueue("2")
	lista.Dequeue()
	lista.Enqueue("8")
	lista.Enqueue("7")
	lista.Dequeue()
	lista.Enqueue("5")
	lista.Enqueue("2")
	lista.Dequeue()
	lista.Dequeue()
	lista.Show()


}

// 5 - (WILKER, 2019) - FCC - 2008 -MPE-RS -Técnico em Informática -Área Sistemas Respeitando as ordens de inserção e de retirada dos dados, uma estruturade:
// c) fila é também denominada FIFO ou LIFO. √

// 6 - (WILKER, 2019) As estruturas do tipo LIFO (Last-In-First-Out) e FIFO(First-In-First-Out) são classificadas, respectivamente, como:
// c) pilha e fila; √

// 7 - (WILKER, 2019) TRE/MG –Analista de Sistemas –2005 .“É uma lista linear em que todas as inserções de novos elementos são realizadas numa extremidade da lista e todas as remoções de elementos são feitas na outra extremidade da lista”. Esta definição é adequada à:
// e) fila que é uma estrutura de dados do tipo LIFO (LastIn FirstOut) √

// 8 - Suponha que exista uma pilha de números inteiros, chamada “S” e uma fila de números inteiros “Q”. Mostre, por meio de uma ilustração, o conteúdo de Se Q depois das seguintes operações:
// push(S, 3)
// push(S, 12)
// enqueue(Q, 5)
// enqueue(Q, 8)
// x = pop(S)
// push(S, 2)
// enqueue(Q, x)
// push(S, x)
// y = top(S)
// push(S, 8)


// Pilha S (topo em cima):
// Topo
//  ┌───┐
//  │ 8 │
//  ├───┤
//  │12 │
//  ├───┤
//  │ 2 │
//  ├───┤
//  │ 3 │
//  └───┘
// Base


// Fila Q (frente → fim):
// Frente → [5] – [8] – [12] ← Fim


// 9 - Uma fila é uma estrutura de dados sujeita à seguinte regra de operação: “sempre que houver uma remoção, o elemento removido será aquele que estiver na estrutura há mais tempo.
// ( X ) certo ( ) errado

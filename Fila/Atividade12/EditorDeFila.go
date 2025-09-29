package main

import (
	"Fila/domains"
	"fmt"
)

func main() {
	var lista domains.Queue
	for {
		fmt.Println("\n--- EDITOR DE FILA ---")
		fmt.Println("1 - Enfileirar")
		fmt.Println("2 - Exibir primeiro da fila")
		fmt.Println("3 - Exibir a fila")
		fmt.Println("4 - Exibir último da fila")
		fmt.Println("5 - Esvaziar a fila")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao string
		fmt.Scanf("%s\n", &opcao)

		switch opcao {
		case "1":
			enqueue(&lista)
		case "2":
			begining(&lista)
		case "3":
			show(&lista)
		case "4":
			ending(&lista)
		case "5":
			dequeue(&lista)
		case "0":
			fmt.Println("Encerrando ...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
func enqueue(lista *domains.Queue) {
	fmt.Println("Digite o valor a Adicionar")
	var val string
	fmt.Scanf("%s\n", &val)
	lista.Enqueue(val)
	fmt.Println("Valor Adicionado : ", val)
}
func begining(lista *domains.Queue) {
	node := lista.Begining()
	fmt.Println("Valor do Primeiro Item")
	fmt.Println(node.Value)
}
func show(lista *domains.Queue) {
	lista.Show()
}
func ending(lista *domains.Queue) {
	node := lista.Ending()
	fmt.Println("Valor do Ultimo Item")
	fmt.Println(node.Value)
}
func dequeue(lista *domains.Queue) {
	if lista.IsEmpty() {
		fmt.Println("Fila vazia! Nada a remover.")
		return
	}
	fmt.Println("Removendo item da lista")
	lista.Dequeue()
	fmt.Println("Item Removido")
}

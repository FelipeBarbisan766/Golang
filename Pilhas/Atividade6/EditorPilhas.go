package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Pilhas/domains"
)

func main() {
	var pilha domains.Stack
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- EDITOR DE PILHA ---")
		fmt.Println("1 - Empilhar (Push)")
		fmt.Println("2 - Desempilhar (Pop)")
		fmt.Println("3 - Exibir elemento do topo")
		fmt.Println("4 - Exibir a pilha")
		fmt.Println("5 - Esvaziar a pilha")
		fmt.Println("6 - Sair")
		fmt.Print("Escolha uma opção: ")

		opcao, _ := reader.ReadString('\n')
		opcao = strings.TrimSpace(opcao)

		switch opcao {
		case "1":
			fmt.Print("Digite o valor a empilhar: ")
			valor, _ := reader.ReadString('\n')
			valor = strings.TrimSpace(valor)
			pilha.Push(valor)
			fmt.Println("Empilhado:", valor)

		case "2":
			if pilha.IsEmpty() {
				fmt.Println("A pilha está vazia, nada para desempilhar.")
			} else {
				removido := pilha.Pop()
				fmt.Println("Desempilhado:", removido.Value)
			}

		case "3":
			if pilha.IsEmpty() {
				fmt.Println("A pilha está vazia, não há topo.")
			} else {
				fmt.Println("Elemento do topo:", pilha.Top.Value)
			}

		case "4":
			fmt.Println("Conteúdo da pilha (do topo até a base):")
			pilha.Show()

		case "5":
			for !pilha.IsEmpty() {
				removido := pilha.Pop()
				fmt.Println("Removendo:", removido.Value)
			}
			fmt.Println("A pilha foi esvaziada.")

		case "6":
			fmt.Println("Saindo... Esvaziando pilha final.")
			for !pilha.IsEmpty() {
				removido := pilha.Pop()
				fmt.Println("Removendo:", removido.Value)
			}
			return

		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}

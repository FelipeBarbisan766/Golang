package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Pilhas/domains"
)

func main() {
	
	pecas := []string{
		"Base do corpo",
		"Armadura do tronco",
		"Braço esquerdo",
		"Perna esquerda",
		"Braço direito",
		"Perna direita",
		"Elmo",
	}

	
	var cavaleiro domains.Stack

	
	fmt.Println("Montando cavaleiro medieval...")
	for _, peca := range pecas {
		cavaleiro.Push(peca)
		fmt.Println("Montando:", peca)
	}

	fmt.Println("\nCavaleiro montado com sucesso!\n")

	
	fmt.Println("Estado atual da pilha de peças:")
	cavaleiro.Show()

	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nDigite o nome da peça que deseja trocar: ")
	input, _ := reader.ReadString('\n')
	pecaDefeituosa := strings.TrimSpace(input)

	
	var removidas domains.Stack

	fmt.Println("\nIniciando manutenção...")
	
	for !cavaleiro.IsEmpty() {
		peca := cavaleiro.Pop()
		fmt.Println("Removendo:", peca.Value)
		if peca.Value == pecaDefeituosa {
			fmt.Println("Peça defeituosa encontrada:", peca.Value)
			fmt.Println("Substituindo por uma nova peça...")
			cavaleiro.Push(peca.Value + " (nova)")
			break
		}
		removidas.Push(peca.Value)
	}

	
	for !removidas.IsEmpty() {
		peca := removidas.Pop()
		fmt.Println("Recolocando:", peca.Value)
		cavaleiro.Push(peca.Value)
	}

	
	fmt.Println("\nProduto final remontado com sucesso!")
	fmt.Println("Estado final da pilha de peças:")
	cavaleiro.Show()
}

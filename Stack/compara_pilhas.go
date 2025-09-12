/* =========================================
   Disciplina de Estrutura de Dados
   Resposta Exercício = Lista sobre Pilhas
   Compara se duas pilhas são iguais
============================================*/

// Importa módulos e classes
package main

import (
	"fmt"

	"ronaldo.prof/domains"
)

func main() {

	var Pilha1 domains.Stack
	var Pilha2 domains.Stack
	menu(&Pilha1, &Pilha2)
}

// mostra menu de opções do programa e faz chamada ao
// prompt ṕara receber as entradas feitas pelo usuário
func menu(Pilha1, Pilha2 *domains.Stack) {
	opcaoUsuario := ""

	for {
		if opcaoUsuario == "0" {
			return
		}

		fmt.Println("    Pilhas Iguais")
		fmt.Println("=====================")
		fmt.Println("1 - Informar pilha 1")
		fmt.Println("2 - Informar pilha 2 ")
		fmt.Println("3 - Compara pilhas")
		fmt.Println("4 - Mostra pilhas")
		fmt.Println("0 - Finalizar programa")
		fmt.Println("Digite uma opção")
		fmt.Scanf("%s", &opcaoUsuario)

		switch opcaoUsuario {
		case "1":
			fmt.Println("Inserir item na Pilha 1")
			AddInStack(Pilha1)
		case "2":
			fmt.Println("Inserir item na Pilha 2")
			AddInStack(Pilha2)
		case "3":
			fmt.Println("Comparar Pilhas 1 e 2")
			if IsSameSize(Pilha1, Pilha2) {
				fmt.Println("Pilhas com tamanhos iguais")
			} else {
				fmt.Println("Pilhas com tamanhos diferentes")
			}

			IsSameItems(Pilha1, Pilha2)
		case "4":
			fmt.Println("Mostrar Pilhas 1 e 2")
			Pilha1.Show()
			Pilha2.Show()
		case "0":
			fmt.Println("Bye!!!")
			opcaoUsuario = "0"
		default:
			fmt.Println("Opção Inválida")
		}
	}
}

func AddInStack(xStack *domains.Stack) {
	valor := ""
	fmt.Println("Digite valor a inserir: ")
	fmt.Scanf("%s", &valor)
	xStack.Push(valor)
	fmt.Println("valor %s inserido", valor)
	return
}

func IsSameSize(xStack1, xStack2 *domains.Stack) bool {
	fmt.Println("Tamanho: ")
	fmt.Printf("           Pilha 1 = %d elementos \n", xStack1.Size())
	fmt.Printf("           Pilha 2 = %d elementos \n", xStack2.Size())
	samesize := xStack1.Size() == xStack2.Size()
	return samesize
}

func IsSameItems(xStack1, xStack2 *domains.Stack) {
	zeroSize := xStack1.IsEmpty() || xStack2.IsEmpty()
	sameSize := xStack1.Size() == xStack2.Size()
	sameValue := true
	for {
		if sameValue && xStack1.Size() > 0 && xStack2.Size() > 0 {
			nodeStack1 := xStack1.Pop()
			nodeStack2 := xStack2.Pop()
			sameValue = (nodeStack1.Value == nodeStack2.Value)
		}

		if sameValue && sameSize {
			if zeroSize {
				fmt.Println("Não existem elementos para serem comparados.  As pilhas estã vazias")
				break
			} else {
				fmt.Println("Pilhas com elementos iguais")
				fmt.Println("PILHAS SÃO IGUAIS")
				break
			}
		} else {
			fmt.Println("Pilhas com elementos diferentes")
			fmt.Println("PILHAS DIFERENTES")
			break
		}
	}
}

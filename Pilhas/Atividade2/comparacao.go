/* =========================================
   Disciplina de Estrutura de Dados
   Resposta Exercício = Lista sobre Pilhas
   Compara se duas pilhas são iguais
============================================*/

// Importa módulos e classes
package main

import (
	"fmt"

	"Pilhas/domains"
)

func main() {

	var Pilha1 domains.Stack
	var Pilha2 domains.Stack
	Pilha1.Push("10")
	Pilha1.Push("20")
	Pilha1.Push("30")

	Pilha2.Push("10")
	Pilha2.Push("20")
	Pilha2.Push("30")
	Pilha2.Push("40")

	IsSameItems(&Pilha1, &Pilha2)
	IsSameSize(&Pilha1, &Pilha2)
	
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

package main

import (
	"fmt"

	"prof.ronaldo/domains"
)

func main() {

	var filaNaLoterica domains.Queue

	//inserindo pessoas na fila da Lotérica

	filaNaLoterica.Enqueue("José")
	filaNaLoterica.Enqueue("Alcione")
	filaNaLoterica.Enqueue("Getulino")
	filaNaLoterica.Enqueue("Whatterson")

	fmt.Printf("Tamanho da Fila = %d \n", filaNaLoterica.Size())

	// Listando pessoas na Fila da Loterica
	contador := 1
	pessoa := filaNaLoterica.Begining()

	for {
		if pessoa == nil {
			break
		}
		fmt.Printf("Pessoa na posição: %d  => %s \n", contador, pessoa.Value)
		pessoa = pessoa.Next
		contador++
	}
	fmt.Println()

	// Removendo pessoas da Fila da Lotérica (atendimento)
	atendimento := 1
	pessoaAtendida := filaNaLoterica.Dequeue()

	for {
		fmt.Printf("Pessoa %d atendida. Nome: %s \n", atendimento, pessoaAtendida.Value)
		if pessoaAtendida.Next == nil {
			break
		}
		pessoaAtendida = filaNaLoterica.Dequeue()
		atendimento++

	}

}

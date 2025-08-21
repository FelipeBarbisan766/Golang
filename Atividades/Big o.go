package main

import (
	"fmt"
	"strings"
)

func mainBig() {
	// Ex 01
	// fmt.Println("Digite um Valor de Soma:")
	// var valor int
	// fmt.Scanf("%d\n", &valor)
	// array := []int{7,4,6,5,1,2}
	// for i := 0; i < len(array); i++ {
	// 	for j := 0; j < len(array); j++ {
	// 		teste := array[i] + array[j]
	// 		if teste == valor {
	// 			fmt.Println("O valor da soma:", array[i], "+",array[j]," é incontrado nos indices :", i, j)
	// 		}
	// 	}
	// }
	// Nivel de Complexidade O(n^2)

	// Ex 02
	// 	Backpack := map[string]int{
	// 		"PedraMagica": 1,
	// 		"PedraFogo": 2,
	// 		"Artefato": 0,
	// 		"Bomba": 3,
	// 		"Armadilha": 4,
	// 	}
	// 	fmt.Println("Digite o Nome do Item:")
	// 	var item string
	// 	fmt.Scanf("%s\n", &item)
	// 	inBackpack(item, Backpack)
	// }
	// func inBackpack(item string,Backpack map[string]int) {
	// 	if Backpack[item] >= 1 {
	// 		quant := Backpack[item]
	// 		fmt.Println("Você tem ",quant," do item ", item)
	// 	} else {
	// 		fmt.Println("O item", item, "não está na mochila")
	// 	}
	// }
	// Ex 03
	act := map[string]int{
		"Faculdade": 5,
		"Trabalho":  2,
		"Estudo":    3,
		"Exercício": 1,
		"Descanso":  4,
	}
	fmt.Println("Digite o nome de uma função (add,list,remove):")
	var funcao string
	fmt.Scanf("%s\n", &funcao)
	if strings.ToLower(funcao) == "add" {
		Add_Act(act)
	} else if strings.ToLower(funcao) == "list" {
		List_Act(act)
	} else if strings.ToLower(funcao) == "remove" {
		Remove_Act(act)
	}
}
func Add_Act(act map[string]int) {
	fmt.Println("Digite o nome da atividade:")
	var atividade string
	fmt.Scanf("%s\n", &atividade)
	fmt.Println("Digite o tempo da atividade:")
	var tempo int
	fmt.Scanf("%d\n", &tempo)
	act[atividade] = tempo
	fmt.Println("Atividade adicionada com sucesso!")
	fmt.Println("A lista de atividades é:", act)

}
func List_Act(act map[string]int) {
	fmt.Println("Atividades:")
	for atividade, tempo := range act {
		fmt.Printf("%s: %d horas\n", atividade, tempo)
	}
}
func Remove_Act(act map[string]int) {
	fmt.Println("Digite o nome da atividade:")
	var atividade string
	fmt.Scanf("%s\n", &atividade)
	delete(act, atividade)
	fmt.Println("Atividade deletada com sucesso!")
	fmt.Println(act)

}

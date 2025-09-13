package main

import (
	"fmt"
	"Pilhas/domains" 
	"strings"
)
func main() {
	valores:= []string{
		"arara",
		"ovo",
		"Ame a ema",
		"socorram me subi no onibus em marrocos",
		"golang",
		"casa",
	}
	fmt.Println("--- Verificador de Palíndromos com Pilha ---")

	for _, texto := range valores {
		if IsPalindrome(texto) {
			fmt.Printf("'%s' é um palíndromo.\n", texto)
		} else {
			fmt.Printf("'%s' NÃO é um palíndromo.\n", texto)
		}
	}
}
func IsPalindrome(text string) bool {
	cleanedText := strings.ToLower(strings.ReplaceAll(text, " ", ""))
	stack := domains.Stack{}
	for _, char := range cleanedText {
		stack.Push(string(char))
	}
	var reversedText strings.Builder
	for !stack.IsEmpty() {
		node := stack.Pop()
		reversedText.WriteString(node.Value)
	}
	return cleanedText == reversedText.String()
}
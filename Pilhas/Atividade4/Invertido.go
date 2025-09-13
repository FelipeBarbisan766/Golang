package main

import (
	"Pilhas/domains"
	"fmt"
	"strings"
)

func main() {
	var stack domains.Stack
	fmt.Println("Digite um Texto:")
	var value string
	fmt.Scanf("%s", &value)
	stack.Push(value)
	Invertido(&stack)
	stack.Show()

	//  resposta: 10
}

func Invertido(stack *domains.Stack) {
	var reversedText strings.Builder
	for !stack.IsEmpty() {
		node := stack.Pop()
		reversedText.WriteString(node.Value)
	}
	fmt.Println("Texto invertido:", reversedText.String())
}

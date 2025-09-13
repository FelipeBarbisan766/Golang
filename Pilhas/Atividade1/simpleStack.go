package main

import (
	"Pilhas/domains"
)

func main() {
	var stack domains.Stack
	stack.Push("10")
	stack.Push("20")
	stack.Push("30")
	stack.Push("40")
	stack.Show()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Show()
	//  resposta: 10
}

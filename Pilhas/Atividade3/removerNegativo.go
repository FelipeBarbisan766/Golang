package main

import (
	"Pilhas/domains"
)

func main() {
	var stack domains.Stack
	stack.Push("30")
	stack.Push("-15")
	stack.Push("20")
	stack.Push("-25")
	RemoveNegative(&stack)
	stack.Show()
}

func RemoveNegative(originalStack *domains.Stack) {
	var tempStack domains.Stack
	for !originalStack.IsEmpty() {
		node := originalStack.Pop()
		if node.Value[0] != '-' {
			tempStack.Push(node.Value)
		}
	}
	for !tempStack.IsEmpty() {
		originalStack.Push(tempStack.Pop().Value)
	}
}
	
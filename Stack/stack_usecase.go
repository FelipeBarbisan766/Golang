package main

import (
	"ronaldo.prof/domains"
)

func main() {
	var stack domains.Stack
	stack.Push("Livro 1")
	stack.Push("Livro 2")
	stack.Push("livro 3")
	stack.Show()
	stack.Pop()
	stack.Show()
	stack.Pop()
	stack.Show()
	stack.Pop()
}

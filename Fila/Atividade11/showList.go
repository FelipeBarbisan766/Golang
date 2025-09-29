package main

import (
	"Fila/domains"
)

func main() {
	var lista domains.Queue
	lista.Enqueue("1")
	lista.Enqueue("2")
	lista.Enqueue("3")
	lista.Enqueue("4")
	lista.Enqueue("5")
	lista.Enqueue("6")
	lista.Enqueue("7")
	lista.Enqueue("8")
	lista.Enqueue("9")
	lista.Enqueue("0")

	lista.Show()
}

package main

import(
	"Lista/domain"
)

func main() {
	var list domains.LinkedList
	list.Add("Maria",0)
	list.Add("João",1)
	list.Add("Pedro",0)
	list.Add("Ana",2)
	list.Show()
}

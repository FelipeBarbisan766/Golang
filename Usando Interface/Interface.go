//! Interface
//* ->Coleção de assinaturas de metodo. Estabelecem "contrato" a ser cumprido pelos metodos
//* ->Structs são rigidamente tipadas e servem para agrupar dados
//* ->Interfaces definem um congunto de metodos que precisam ser implementados mas o codigo que implementa o metodo pode ser feito de forma diferente o que torna a inteface flexivel (adaptavel)

package main

import (
	"fmt"
)


	//interface
type Geometry interface{
	area() float32
}
type Rectangle struct{
	length,width float32
}
func(r Rectangle) area() float32{
	return r.length * r.width
}

type Triangle struct{
  l1, l2, base float32
}

func(t Triangle) area() float32{
	return (t.base * t.l1) / 2
}

func CalculaMedidas(g Geometry) {
	fmt.Println("Area: ", g.area())
}
func main() {
	rect := Rectangle{10,20}
	CalculaMedidas(rect)
	tri := Triangle{10,20,30}
	CalculaMedidas(tri)
}


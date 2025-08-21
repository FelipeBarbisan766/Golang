package main

import (
	// "fmt"
	// "sort"
)

func main() {
	// EX 1
	// a := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(a[0])
	// fmt.Println(a[1])
	// fmt.Println(a[2])
	// fmt.Println(a[3])
	// fmt.Println(a[4])

	// EX 2
	// b := make([]int,0, 5)
	// b = append(b, 1,2,3,4,5,6,7,8,9,0)
	// fmt.Println(b)

	// EX 3
	// c := [8]int{}
	// fmt.Println("Tamanho do array", len(c))

	// EX 4
	// d := [3]int{1,2,3}
	// e := make([]int,5)
	// e = append(e,d[0],d[1],d[2])
	// fmt.Println(e)

	// EX 5
	// f:= [5]int{1,3,5,5,2}
	// g:= [5]int{3,2,4,5,2}
	// for i:=0; i < len(f); i++ {
	// 	if f[i] == g[i] {
	// 		fmt.Println("Igual")
	// 	} else {
	// 		fmt.Println("Diferente")
	// 	}
	// }

	// EX 6
	// h:= [3]int{1,3,5}
	// i:= [2]int{7,4}
	// j:= make([]int,0,5)
	// j = append(j,h[0],h[1],h[2],i[0],i[1])
	// fmt.Println(h)
	// fmt.Println(i)
	// fmt.Println(j)

	// EX 7
	// k := [3]int{5,7,8}
	// l := make([]int,0,5)
	// for i:=2; i >= 0; i-- {
	// 	l = append(l,k[i])
	// }
	// fmt.Println(l)

	// EX 8
	// m:= [4]int{7,3,10,10}
	// var n int
	// for i:=0; i < len(m); i++ {
	// 	n = n + m[i]
	// }
	// fmt.Println(n)

	// EX  9
	// o := [5]int{1,2,3,4,5}
	// for i:= 0; i < len(o); i++{
	// 	if o[i] % 2 == 0 {
	// 		fmt.Println(o[i])
	// 	}
	// }

	// EX 10
	// fmt.Println("Digite o tamanho do array")
	// var tamanho int
	// fmt.Scanf("%d", &tamanho)
	// p := make([]int,tamanho)
	// fmt.Println(p)

	// EX 11
	// q := [5]int{1,2,3,4,5}
	// fmt.Println("Digite um valor para ser excloido")
	// var r int
	// fmt.Scanf("%d", &r)
	// for i:=0; i < len(q); i++ {
	// 	if r == q[i]{
	// 		q[i] = 0
	// 	}
	// }
	// fmt.Println(q)

	// EX 12
	// s := [5]int{1,2,3,4,5}
	// fmt.Println("Digite um valor na qual procura")
	// var t int
	// fmt.Scanf("%d", &t)
	// for i:=0; i < len(s); i++ {
	// 	if t == s[i]{
	// 		fmt.Printf("O valor procurado %v está no indice %v",s[i],i)
	// 	}
	// }

	// 13
	// u:= [4]int{5,0,7,1}
	// v:= [4]int{0,4,2,4}
	// w:= [4]int{}
	// for i:=0; i < len(u); i++ {
	// 	w[i] = u[i] + v[i]
	// }
	// fmt.Println(u)
	// fmt.Println(v)
	// fmt.Println(w)

	// EX 14
	// x:= [5]int{1,2,3,3,2}
	// var y int
	// var z int
	// fmt.Println("Digite o valor a ser procurado")
	// fmt.Scanf("%d", &y)
	// for i:=0; i<len(x); i++ {
	// 	if y == x[i]{
	// 		z++
	// 	}
	// }
	// fmt.Println(z)

	// EX 15
	// var multi [3][3]int
	// for i:=0; i < 3; i++ {
	// 	for l:=0; l < 3; l++ {
	// 		fmt.Println("Digite os valores da linha",i)
	// 		fmt.Println("Digite os valores da coluna",l)
	// 		fmt.Scanf("%d\n", &multi[i][l])
	// 	}
	// }
	// fmt.Print(multi)

	// EX 16
	// var multi [3][3]int
	// multi = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	// aa := [3]int{}
	// for i:=0; i < 3; i++ {
	// 	aa[i] = multi[i][0]+ multi[i][1]+ multi[i][2]
	// }
	// fmt.Print(aa)

	// EX 17
	// var multi [2][3]int
	// multi = [2][3]int{{1,2,3},{4,5,6}}
	// ab := [3]int{}
	// for i:=0; i < 3; i++ {
	// 	ab[i] = multi[0][i]+ multi[1][i]
	// }
	// fmt.Print(ab)

	// EX 18
	// var multi [3][3]int
	// var diagonal [3]int
	// for i:=0; i < 3; i++ {
	// 	for l:=0; l < 3; l++ {
	// 		fmt.Println("Digite os valores da linha",i)
	// 		fmt.Println("Digite os valores da coluna",l)
	// 		fmt.Scanf("%d\n", &multi[i][l])
	// 	}
	// }
	// for i:=0; i < len(diagonal); i++{
	// 	diagonal[i] = multi[i][i]
	// }
	// fmt.Println("Matriz")
	// fmt.Println(multi)
	// fmt.Println("Diagonal Direita")
	// fmt.Println(diagonal)

	// EX 19
	// var multi [3][3]int
	// var diagonal [3]int
	// for i := 0; i < 3; i++ {
	// 	for l := 0; l < 3; l++ {
	// 		fmt.Println("Digite os valores da linha", i)
	// 		fmt.Println("Digite os valores da coluna", l)
	// 		fmt.Scanf("%d\n", &multi[i][l])
	// 	}
	// }
	// var l int = 2
	// for i := 0; i < 3; i++ {
	// 	diagonal[i] = multi[i][l]
	// 	l --
	// }
	// fmt.Println("Matriz")
	// fmt.Println(multi)
	// fmt.Println("Diagonal Esquerda")
	// fmt.Println(diagonal)

	// EX 20
	// var multi [3][3]int
	// var soma int
	// for i := 0; i < 3; i++ {
	// 	for l := 0; l < 3; l++ {
	// 		fmt.Println("Digite os valores da linha", i)
	// 		fmt.Println("Digite os valores da coluna", l)
	// 		fmt.Scanf("%d\n", &multi[i][l])
	// 	}
	// }
	// for i := 0; i < 3; i++ {
	// 	for l := 0; l < 3; l++ {
	// 		soma = soma + multi[i][l]
	// 	}
	// }
	// fmt.Println("Matriz")
	// fmt.Println(multi)
	// fmt.Println("Soma dos valores")
	// fmt.Println(soma)

	// EX 21
	// slice := make([]int, 2)
	// var iten int
	// for i := 0; i < 2; i++ {
	// 	fmt.Println("Digite um valor")
	// 	fmt.Scanf("%d\n", &iten)
	// 	slice = append(slice, iten)
	// }
	// fmt.Println(slice)

	// EX 23
	// fmt.Println("Array")
	// array := [5]int{1,2,3}
	// fmt.Println(array)
	// fmt.Println("Adicionar valor em Array")
	// fmt.Scanf("%d\n", &array[3])
	// fmt.Println("Excluir um valor em Array")
	// var item int
	// fmt.Scanf("%d\n", &item)
	// for i:=0; i < len(array); i++ {
	// 	if item == array[i]{	
	// 		array[i] = 0
	// 	}
	// }
	// fmt.Println("Ordenagem valor em Array")
	// sort.Ints(array[:])
	// fmt.Println(array)

	// fmt.Println("------------------------")
	// fmt.Println("Slice")
	// slice := make([]int,5)
	// slice = append(slice, 1,3,2,4)
	// fmt.Println(slice)
	// fmt.Println("Adicionar valor em Array")
	// var valor int
	// fmt.Scanf("%d\n", &valor)
	// slice = append(slice, valor)
	// fmt.Println("Excluir um valor em Array")
	// fmt.Scanf("%d\n", &item)
	// for i:=0; i < len(slice); i++ {
	// 	if item == slice[i]{	
	// 		slice[i] = 0
	// 	}
	// }
	// fmt.Println("Ordenagem valor em Slice")
	// sort.Ints(slice[:])
	// fmt.Println(slice)

	// fmt.Println("------------------------")
	// fmt.Println("Map")
	// maping := map[string]string{"item":"alguma coisa", "outro item":"outra coisa mais"}
	// fmt.Println(maping)
	// fmt.Println("Adicionar valor em Map")
	// var palavra string
	// fmt.Scanf("%d\n", &palavra)
	// maping["novo item"] = palavra
	// fmt.Println("Excluir um valor em Map")
	// var procura string
	// fmt.Scanf("%d\n", &procura)
	// for i:=0; i < len(maping); i++ {
	// 	if procura == maping[procura]{	
	// 		maping[procura] = "-"
	// 	}
	// }
}

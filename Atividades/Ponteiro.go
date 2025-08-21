package main

func mainPoint() {
	//! 1
	// a := 10
	// v := &a
	// fmt.Println("Conteúdo da variável: ", *v)
	// fmt.Println("Endereço da variável: ", &v)

	//! 2
	// 	f := 10.5
	// 	alteraPorValor(f)
	// 	fmt.Println("----------------------------")
	// 	p := &f
	// 	alteraPorReferencia(p)
	// 	fmt.Println("----------------------------")
	// 	testaFloat := 125.93
	// 	alteraPorValor(testaFloat)
	// 	fmt.Println("Valor fora da função:", testaFloat)
	// 	fmt.Println("----------------------------")
	// 	ptest := &testaFloat
	// 	alteraPorReferencia(ptest)
	// 	fmt.Println("Valor fora da função:", testaFloat)
	// 	fmt.Println("----------------------------")
	// }
	// func alteraPorValor(f float64) {
	// 	f = 522.65
	// 	fmt.Println("Valor dentro da função:", f)
	// }
	// func alteraPorReferencia(p *float64) {
	//#	*p = 331.33
	// 	fmt.Println("Valor dentro da função:", *p)

	//! EX 3
	// Vantagens:
	// 1. Acesso rápido a memória:
	// Ponteiros permitem acesso direto aos endereços de memória, o que pode resultar em maior velocidade de execução em comparação com acessos indiretos via índices de array.
	// 2. Manipulação de dados complexos:
	// Ponteiros são fundamentais para implementação de estruturas de dados como listas ligadas, árvores e outras estruturas que exigem manipulação dinâmica da memória.
	// Desvantagens:
	// 1. Risco de erros de memória:
	// Erros ao manipular ponteiros, como apontar para memória inválida, podem causar travamentos, erros de execução ou vazamentos de memória.
	// 2. Complexidade do código:
	// A manipulação de ponteiros pode tornar o código mais difícil de entender e manter, especialmente para programadores menos experientes.

	//! EX 4
	// Falso, o uso de ponteiros pode até melhorar o desempenho em algumas partes, mas ele por si só não,
	// a melhora do desempenho depende de outros fatores além desse.

	//!EX 05
	//  a, b, c, d, e := 10, 20, 30, 40, 50
	//  arr := [5]*int{&a, &b, &c, &d, &e}
	//  valoresArray(arr)
	//  fmt.Println("Valores após chamada da função:")
	//  fmt.Println("a =", a)
	//  fmt.Println("b =", b)
	//  fmt.Println("c =", c)
	//  fmt.Println("d =", d)
	//  fmt.Println("e =", e)
	// }
	// func valoresArray(arr [5]*int) {
	// 	valores := [5]int{125, 0, 240, 12, 39}
	// 	for i := 0; i < 5; i++ {
	//#	*arr[i] = valores[i]
	// 	}
	//quando você passa um ponteiro para uma função e modifica o valor apontado por ele você está alterando diretamente o valor original na memória. Por isso, as variáveis no main() foram atualizadas.
}

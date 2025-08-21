// Alunos:
// Eduardo Vizicato
// Felipe Barbisan
// Felipe Galati
// Vitor Hugo
// Curso: 4º AMS (ADS)

package main
import (
	"fmt"
)
// EX 1

// type Geometry interface{
// 	area() float32
// }

// type Trapezio struct{
// 	baseinf, basesup, altura float32
// }

// func (t Trapezio) area() float32{
// 	return ((t.baseinf+t.basesup))/2
// }

// //uso da interface
// func CalculaMedidas(g Geometry){
// 	fmt.Println("Área:",g.area())
// }

// func main() {
// 	trap := Trapezio{3,4,3}
// 	CalculaMedidas(trap)
	
// }
// EX 2
// type Geometry interface{
// 	perimetro() float32
// }
// type Rectangle struct{
// 	length,width float32
// }
// func(r Rectangle) perimetro() float32{
// 	return 2 * (r.length + r.width)
	
// }
// type trapezio struct{
// 	baseMaior, baseMenor, altura float32
// }
// func(t trapezio) perimetro() float32{
// 	return t.baseMaior + t.baseMenor + 2 * t.altura
	
// }
// func CalculaMedidas(g Geometry) {
// 	fmt.Println("Perimetro: ", g.perimetro())
// }
// func main(){
// 	//retangulo
// 	r := Rectangle{10,20}
// 	//trapezio
// 	t := trapezio{10,20,30}
	
// 	CalculaMedidas(r)
// 	CalculaMedidas(t)
// }
// EX 3


type ContaBancaria interface {
	AbrirConta(nome string, cpf string, numeroConta string, banco string, saldoInicial float64)
	ApresentarSaldo()
	Retirar(valor float64)
	Depositar(valor float64)
}

type Conta struct {
	nome        string
	cpf         string
	numeroConta string
	banco       string
	saldo       float64
}


func (c *Conta) AbrirConta(nome string, cpf string, numeroConta string, banco string, saldoInicial float64) {
	c.nome = nome
	c.cpf = cpf
	c.numeroConta = numeroConta
	c.banco = banco
	c.saldo = saldoInicial
	fmt.Println("Conta criada com sucesso!")
}

func (c *Conta) ApresentarSaldo() {
	fmt.Printf("Cliente: %s\nCPF: %s\nBanco: %s\nConta: %s\nSaldo Atual: R$ %.2f\n\n", c.nome, c.cpf, c.banco, c.numeroConta, c.saldo)
}

func (c *Conta) Retirar(valor float64) {
	if valor > c.saldo {
		fmt.Println("Saldo insuficiente para retirada.")
	} else {
		c.saldo -= valor
		fmt.Printf("Retirada de R$ %.2f realizada com sucesso.\n", valor)
	}
	c.ApresentarSaldo()
}

func (c *Conta) Depositar(valor float64) {
	c.saldo += valor
	fmt.Printf("Depósito de R$ %.2f realizado com sucesso.\n", valor)
	c.ApresentarSaldo()
}

func mainInterface() {
	var minhaConta ContaBancaria = &Conta{}

	minhaConta.AbrirConta("Eduardo Vizicato", "123.456.789-00", "000123", "Banco XPTO", 1000.00)

	minhaConta.ApresentarSaldo()

	minhaConta.Depositar(500.00)

	minhaConta.Retirar(200.00)
}
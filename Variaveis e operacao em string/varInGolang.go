// ! Variáveis em Golang
// ?-> Escopo fora das funções = Uso da palavra reservada "var"
// ? -> Escopo de funções = Short Declaration - name := "ronaldo"
// * Declaração de variável sem indicação de valor fará com que sejá atribuido automaticamente o menor valor permitido conforme o tipo de dado que a variável representa Ex: var nome String
// ? => Constantes: podem ser definidas com o uso da palavra reservada "const", Ex: const valorPi = 3.14638
// -----------------------------------------------------------------
// ! Principais tipos Basicos Golang
// ? -> Inteiro: int, int8, int16, int32, rune, uint ...
// ? -> Decimais: float32, float64
// ? -> Byte: byte
// ? -> Conjunto de Caracteres: String
// ? -> Verdadeiro ou Falso: bool
// -----------------------------------------------------------------
//! Operações String
//? len(name) => Quantidade de Caracteres
//? name[0] => "J"
//? name[0:2] => "Jo"
//? name[:] => "John"
//? string.ToUpper() => Letras Maiúsculas
// -----------------------------------------------------------------

package main
import (
	"fmt"
	"strings"
)
func main() {
	var nome string = "John Deere"
	fmt.Println("Nome", nome)
	fmt.Println("Quantidade de Letras:", len(nome))
	fmt.Println("A Primeira Letra do nome é: ", string(nome[0]))
	fmt.Println("A ultima Letra do nome é:", string(nome[len(nome)-1]))
	fmt.Println("Todas as letras do nome:", string(nome[:]))
	fmt.Println("Todas as letras do nome em Maiúsculo:", strings.ToUpper(nome))
	// -----------------------------------------------------------------
	fmt.Println(fmt.Sprintf("O nome é %s em maiúsculo fica %s", nome, strings.ToUpper(nome)))
	var area float32 = 2.5 / 2
	fmt.Println("Area", area)
	fmt.Println(fmt.Sprintf("A Parte Inteira da Area é : %f", area))
	// -----------------------------------------------------------------
	var myarray = [3]string{"First", "Second", "Third"}
	fmt.Println(myarray)
	fmt.Println(myarray[0])
	fmt.Println(myarray[0:2])
}
// -----------------------------------------------------------------

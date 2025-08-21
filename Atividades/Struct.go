// Aluno: Felipe Adrian Lourenço Barbisan
// Curso: 4º AMS (ADS) Fatec

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
// EX 01 - Estudante

// func main(){
// 	 alu := Estudante{
// 	 	nome:       "Ronaldo",
// 	 	disciplina: "Estrutura de dados",
// 	 }
// }
// 	alu.setNotas(7.4, 8.2, 6.5, 9.0)

// 	fmt.Printf("Estudante: %+v\n", alu)
// 	fmt.Printf("Média das notas: %.2f\n", alu.mediaNotas())
// }

// type Estudante struct{
// 	nome string
// 	notas []float32
// 	disciplina string
// }

// func (e *Estudante) setNotas(novasNotas ...float32) {
// 	e.notas = novasNotas
// }

// func (e *Estudante) mediaNotas() float32 {
// 	var soma float32
// 	for _, nota := range e.notas {
// 		soma += nota
// 	}
// 	if len(e.notas) == 0 {
// 		return 0
// 	}
// 	return soma / float32(len(e.notas))
// }

// Ex 02 - Carro
// 		c1 := Car{}
// 		Registercar(&c1)
// 		c2 := Car{}
// 		Registercar(&c2)
// 		c3 := Car{}
// 		Registercar(&c3)
// 		c4 := Car{}
// 		Registercar(&c4)
// 		c5 := Car{}
// 		Registercar(&c5)
// 		fmt.Println("Carro 1: ", c1)
// 		fmt.Println("Carro 2: ", c2)
// 		fmt.Println("Carro 3: ", c3)
// 		fmt.Println("Carro 4: ", c4)
// 		fmt.Println("Carro 5: ", c5)
// 		ComsuptionCalculation(c1.Consumption);
// 		ComsuptionCalculation(c2.Consumption);
// 		ComsuptionCalculation(c3.Consumption);
// 		ComsuptionCalculation(c4.Consumption);
// 		ComsuptionCalculation(c5.Consumption);

// }

// func ComsuptionCalculation(c float32) {
// 	fmt.Println("Digite o Km da viagem Pretendida:")
// 	var km float32
// 	fmt.Scanf("%s\n", &km)
// 	fmt.Println("Digite o Preço do Combustivel:")
// 	var fuel float32
// 	fmt.Scanf("%s\n", &fuel)
// 	calc := (km / c) * fuel
// 	fmt.Println("O Gasto dessa viagem será de : R$", calc)
// }
// func Registercar(c *Car) {

// 	fmt.Println("Digite o nome da marca do carro:")
// 	fmt.Scanf("%s\n", &c.Brand)
// 	fmt.Println("Digite o modelo do carro:")
// 	fmt.Scanf("%s\n", &c.Model)
// 	fmt.Println("Digite o ano do carro:")
// 	fmt.Scanf("%d\n", &c.Year)
// 	fmt.Println("Digite a identificação do carro:")
// 	fmt.Scanf("%s\n", &c.Indetify)
// 	fmt.Println("Digite o consumo do carro:")
// 	fmt.Scanf("%f\n", &c.Consumption)
// }
// type Car struct{
// 	Brand string
// 	Model string
// 	Year int
// 	Indetify string
// 	Consumption float32
// }

// Ex 03 - Funcionário
// 	e1 := Employee{}
// 	RegisterEmployee(&e1)
// 	e2 := Employee{}
// 	RegisterEmployee(&e2)
// 	e3 := Employee{}
// 	RegisterEmployee(&e3)
// 	e4 := Employee{}
// 	RegisterEmployee(&e4)
// 	e5 := Employee{}
// 	RegisterEmployee(&e5)
// 	e6 := Employee{}
// 	RegisterEmployee(&e6)
// 	e7 := Employee{}
// 	RegisterEmployee(&e7)
// 	e8 := Employee{}
// 	RegisterEmployee(&e8)
// 	e9 := Employee{}
// 	RegisterEmployee(&e9)
// 	e10 := Employee{}
// 	RegisterEmployee(&e10)
// 	ShowEmployee(e1, e2, e3, e4, e5, e6, e7, e8, e9, e10)
// 	wageAverage(e1, e2, e3, e4, e5, e6, e7, e8, e9, e10)
// }

// func wageAverage(e1, e2, e3, e4, e5, e6, e7, e8, e9, e10 Employee) {
// 	average := (e1.wage + e2.wage + e3.wage + e4.wage + e5.wage + e6.wage + e7.wage + e8.wage + e9.wage + e10.wage) / 10
// 	fmt.Println("A média salarial é: ", average)
// }

// func ShowEmployee(e1, e2, e3, e4, e5, e6, e7, e8, e9, e10 Employee) {

// 	fmt.Println("Funcionário 1: ", e1)
// 	fmt.Println("Salario Liquido:", e1.wage+e1.benefits-e1.discounts)
// 	fmt.Println("Funcionário 2: ", e2)
// 	fmt.Println("Salario Liquido:", e2.wage+e2.benefits-e2.discounts)
// 	fmt.Println("Funcionário 3: ", e3)
// 	fmt.Println("Salario Liquido:", e3.wage+e3.benefits-e3.discounts)
// 	fmt.Println("Funcionário 4: ", e4)
// 	fmt.Println("Salario Liquido:", e4.wage+e4.benefits-e4.discounts)
// 	fmt.Println("Funcionário 5: ", e5)
// 	fmt.Println("Salario Liquido:", e5.wage+e5.benefits-e5.discounts)
// 	fmt.Println("Funcionário 6: ", e6)
// 	fmt.Println("Salario Liquido:", e6.wage+e6.benefits-e6.discounts)
// 	fmt.Println("Funcionário 7: ", e7)
// 	fmt.Println("Salario Liquido:", e7.wage+e7.benefits-e7.discounts)
// 	fmt.Println("Funcionário 8: ", e8)
// 	fmt.Println("Salario Liquido:", e8.wage+e8.benefits-e8.discounts)
// 	fmt.Println("Funcionário 9: ", e9)
// 	fmt.Println("Salario Liquido:", e9.wage+e9.benefits-e9.discounts)
// 	fmt.Println("Funcionário 10: ", e10)
// 	fmt.Println("Salario Liquido:", e10.wage+e10.benefits-e10.discounts)

// }

// func RegisterEmployee(employee *Employee) {
// 	fmt.Println("Digite o nome do funcionário:")
// 	fmt.Scanf("%s\n", &employee.name)
// 	fmt.Println("Digite o cargo do funcionário:")
// 	fmt.Scanf("%s\n", &employee.position)
// 	fmt.Println("Digite o salário do funcionário:")
// 	fmt.Scanf("%f\n", &employee.wage)
// 	fmt.Println("Digite os benefícios do funcionário:")
// 	fmt.Scanf("%f\n", &employee.benefits)
// 	fmt.Println("Digite os descontos do funcionário:")
// 	fmt.Scanf("%f\n", &employee.discounts)
// }

// type Employee struct {
// 	name      string
// 	position  string
// 	wage      float32
// 	benefits  float32
// 	discounts float32
// }

// Ex 04 - Livro

//  	var livros []livro
// 	l1 := livro{
// 		titulo: "banana",
// 		autor: "J.R.R. Tolkien",
// 		ano: 1954,
// 		status: true,
// 	}

// 	l1.cadastrarLivro(&livros,l1)
// 	l1.mostrarLivro(&livros)
// 	fmt.Println("Digite o título do livro:")
// 	var titulo string
// 	fmt.Scanf("%v\n", &titulo)
// 	pesquisarLivro(titulo, livros)
// 	l1.mudarStatus(&livros, l1)
// 	l1.mostrarLivro(&livros)

// }
// type livro struct {
// 	titulo string
// 	autor string
// 	ano int
// 	status bool
// }
// func (l *livro) cadastrarLivro(livros *[]livro, livro livro) {
// 	*livros = append(*livros, livro)
// }
// func pesquisarLivro(t string, livros []livro) {
// 	for _, livro := range livros {
// 		if livro.titulo == t {
// 			fmt.Printf("Livro encontrado: %s, Autor: %s, Ano: %d, Status: %t\n", livro.titulo, livro.autor, livro.ano, livro.status)
// 			return
// 		}
// 	}
// 	fmt.Println("Livro não encontrado.")

// }
// func (l *livro)mostrarLivro(livros *[]livro){
// 	for i, livro := range *livros{
// 		fmt.Printf("Livro %d: %s, Autor: %s, Ano: %d, Status: %t\n", i+1, livro.titulo, livro.autor, livro.ano, livro.status)
// 	}
// }
// func (l *livro) mudarStatus(livros *[]livro, livro livro) {
// 	for i, l := range *livros {
// 		if l.titulo == livro.titulo {
// 			(*livros)[i].status = !(*livros)[i].status
// 			fmt.Printf("Status do livro %s alterado para %t\n", livro.titulo, (*livros)[i].status)
// 			return
// 		}
// 	}
// 	fmt.Println("Livro não encontrado.")

// }
// EX 5
// func main(){

// 	aeroportos := []Aeroporto{
// 		{"GRU", "Aeroporto de Guarulhos", "SaoPaulo"},
// 		{"SDU", "Santos Dumont", "RiodeJaneiro"},
// 		{"CNF", "Confins", "BeloHorizonte"},
// 		{"POA", "Salgado Filho", "PortoAlegre"},
// 		{"BSB", "Juscelino Kubitschek", "Brasília"},
// 	}

// 	voos := []Voo{
// 		{"V101", "GRU", "SDU", "08:30"},
// 		{"V102", "SDU", "GRU", "09:45"},
// 		{"V103", "CNF", "POA", "12:15"},
// 		{"V104", "POA", "CNF", "13:50"},
// 		{"V105", "BSB", "GRU", "15:00"},
// 		{"V106", "GRU", "BSB", "16:30"},
// 		{"V107", "SDU", "CNF", "17:20"},
// 		{"V108", "CNF", "SDU", "18:00"},
// 		{"V109", "BSB", "POA", "19:10"},
// 		{"V110", "POA", "BSB", "20:45"},
// 	}

// 	var origem, destino string

// 	fmt.Println("Digite a origem do voo:")
// 	fmt.Scan(&origem)

// 	fmt.Println("Digite a destino do voo:")
// 	fmt.Scan(&destino)

// 	codigosOrigem := buscarCodigosPorCidade(aeroportos, origem)
// 	codigosDestino := buscarCodigosPorCidade(aeroportos, destino)

// 	if len(codigosOrigem) == 0 || len(codigosDestino) == 0 {
// 		fmt.Println("Cidade de origem ou destino não encontrada.")
// 		return
// 	}

// 	fmt.Printf("\nVoos de %s para %s:\n", origem, destino)
// 	encontrado := false

// 	for _, voo := range voos {
// 		if contem(codigosOrigem, voo.codAeroOrigem) && contem(codigosDestino, voo.codAeroDestino) {
// 			nomeOrigem := buscarNomeAeroporto(aeroportos, voo.codAeroOrigem)
// 			nomeDestino := buscarNomeAeroporto(aeroportos, voo.codAeroDestino)
// 			fmt.Printf("Voo %s - %s -> %s - Horário: %s\n",
// 				voo.codVoo, nomeOrigem, nomeDestino, voo.horario)
// 			encontrado = true
// 		}
// 	}

// 	if !encontrado {
// 		fmt.Println("Nenhum voo encontrado entre as cidades informadas.")
// 	}
// }

// type Voo struct{
// 	codVoo string
// 	codAeroOrigem string
// 	codAeroDestino string
// 	horario string
// }

// type Aeroporto struct{
// 	codigo string
// 	nome string
// 	cidade string
// }

// func buscarCodigosPorCidade(aeroportos []Aeroporto, cidade string) []string {
// 	var codigos []string
// 	for _, a := range aeroportos {
// 		if strings.EqualFold(a.cidade, cidade) {
// 			codigos = append(codigos, a.codigo)
// 		}
// 	}
// 	return codigos
// }

// func buscarNomeAeroporto(aeroportos []Aeroporto, codigo string) string {
// 	for _, a := range aeroportos {
// 		if a.codigo == codigo {
// 			return a.nome
// 		}
// 	}
// 	return "Desconhecido"
// }

// func contem(slice []string, valor string) bool {
// 	for _, item := range slice {
// 		if item == valor {
// 			return true
// 		}
// 	}
// 	return false
// }

// EX6
// type Receita struct{
// 	codReceita int
// 	nome string
// 	ingredientes []Ingrediente
// }

// type Ingrediente struct{
// 	quant float32
// 	uniMedida string
// 	nome string
// }

// func main(){
// 	var maxReceitas = 5
// 	var maxIngredientes = 10

// 	var receitas []Receita
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		fmt.Println("\n--- MENU ---")
// 		fmt.Println("1. Inserir receita")
// 		fmt.Println("2. Listar receitas")
// 		fmt.Println("3. Buscar receita e calcular ingredientes")
// 		fmt.Println("4. Sair")
// 		fmt.Print("Escolha uma opção: ")

// 		scanner.Scan()
// 		opcao := scanner.Text()

// 		switch opcao {
// 		case "1":
// 			if len(receitas) >= maxReceitas {
// 				fmt.Println("Limite de 5 receitas atingido.")
// 				continue
// 			}
// 			var r Receita
// 			fmt.Print("Código da receita: ")
// 			scanner.Scan()
// 			r.codReceita, _ = strconv.Atoi(scanner.Text())

// 			fmt.Print("Nome da receita: ")
// 			scanner.Scan()
// 			r.nome = scanner.Text()

// 			fmt.Print("Quantos ingredientes? (máx 10): ")
// 			scanner.Scan()
// 			qtdIng, _ := strconv.Atoi(scanner.Text())

// 			if qtdIng > maxIngredientes {
// 				fmt.Println("Número de ingredientes excede o limite.")
// 				continue
// 			}

// 			for i := 0; i < qtdIng; i++ {
// 				var ing Ingrediente
// 				fmt.Printf("\nIngrediente %d:\n", i+1)

// 				fmt.Print("Nome do ingrediente: ")
// 				scanner.Scan()
// 				ing.nome = scanner.Text()

// 				fmt.Print("Unidade de medida (g, kg, ml, xícara...): ")
// 				scanner.Scan()
// 				ing.uniMedida = scanner.Text()

// 				fmt.Print("Quantidade: ")
// 				scanner.Scan()
// 				qtd, _ := strconv.ParseFloat(scanner.Text(), 32)
// 				ing.quant = float32(qtd)

// 				r.ingredientes = append(r.ingredientes, ing)
// 			}
// 			receitas = append(receitas, r)
// 			fmt.Println("Receita adicionada com sucesso!")

// 		case "2":
// 			if len(receitas) == 0 {
// 				fmt.Println("Nenhuma receita cadastrada.")
// 				continue
// 			}
// 			fmt.Println("\n--- Lista de Receitas ---")
// 			for _, r := range receitas {
// 				fmt.Printf("Receita: %s\n", r.nome)
// 				for _, ing := range r.ingredientes {
// 					fmt.Printf("- %.2f %s de %s\n", ing.quant, ing.uniMedida, ing.nome)
// 				}
// 				fmt.Println()
// 			}

// 		case "3":
// 			fmt.Print("Digite o nome da receita que deseja buscar: ")
// 			scanner.Scan()
// 			nomeBusca := strings.ToLower(scanner.Text())

// 			var receitaEncontrada *Receita
// 			for i := range receitas {
// 				if strings.ToLower(receitas[i].nome) == nomeBusca {
// 					receitaEncontrada = &receitas[i]
// 					break
// 				}
// 			}

// 			if receitaEncontrada == nil {
// 				fmt.Println("Receita não encontrada.")
// 				continue
// 			}

// 			fmt.Print("Quantas vezes deseja repetir a receita? ")
// 			scanner.Scan()
// 			repeticoes, _ := strconv.Atoi(scanner.Text())

// 			fmt.Printf("\nIngredientes para %d vez(es) da receita '%s':\n", repeticoes, receitaEncontrada.nome)
// 			for _, ing := range receitaEncontrada.ingredientes {
// 				total := ing.quant * float32(repeticoes)
// 				fmt.Printf("- %.2f %s de %s\n", total, ing.uniMedida, ing.nome)
// 			}

// 		case "4":
// 			fmt.Println("Encerrando o programa.")
// 			return

// 		default:
// 			fmt.Println("Opção inválida. Tente novamente.")
// 		}
// 	}
// }

// EX 7

type Produto struct {
	codProd int
	nome    string
	preco   float64
	quant   int
}

type Carrinho struct {
	produto   Produto
	quantiCar int
}

func mainStruct() {

	produtosLoja := []Produto{
		{codProd: 1, nome: "Arroz", preco: 5.90, quant: 10},
		{codProd: 2, nome: "Feijão", preco: 6.50, quant: 8},
		{codProd: 3, nome: "Macarrão", preco: 3.75, quant: 15},
		{codProd: 4, nome: "Leite", preco: 4.20, quant: 12},
		{codProd: 5, nome: "Açúcar", preco: 2.80, quant: 20},
	}

	var carrinho []Carrinho
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Lista de Produtos ---")
		for _, p := range produtosLoja {
			fmt.Printf("Código: %d | Nome: %s | Preço: R$%.2f | Estoque: %d\n",
				p.codProd, p.nome, p.preco, p.quant)
		}

		fmt.Print("\nDigite o código do produto que deseja comprar (pressione 0 para finalizar): ")
		scanner.Scan()
		codStr := scanner.Text()
		cod, _ := strconv.Atoi(codStr)

		if cod == 0 {
			break
		}

		produto, encontrado := buscarProdutoPorCodigo(produtosLoja, cod)
		if !encontrado {
			fmt.Println("Produto não encontrado.")
			continue
		}

		fmt.Printf("Quantidade desejada de %s: ", produto.nome)
		scanner.Scan()
		qtdStr := scanner.Text()
		qtd, _ := strconv.Atoi(qtdStr)

		if qtd <= 0 {
			fmt.Println("Quantidade inválida.")
			continue
		}

		if qtd > produto.quant {
			fmt.Printf("Estoque insuficiente. Disponível: %d unidades\n", produto.quant)
			continue
		}

		// Adiciona ao carrinho
		carrinho = append(carrinho, Carrinho{produto: produto, quantiCar: qtd})

		// Atualiza o estoque
		for i := range produtosLoja {
			if produtosLoja[i].codProd == cod {
				produtosLoja[i].quant -= qtd
				break
			}
		}
		fmt.Printf("Adicionado ao carrinho: %d x %s\n", qtd, produto.nome)
	}

	// Finaliza a compra
	if len(carrinho) == 0 {
		fmt.Println("Nenhum item comprado.")
		return
	}

	fmt.Println("\n--- Itens no carrinho ---")
	var total float64
	for _, item := range carrinho {
		subtotal := float64(item.quantiCar) * item.produto.preco
		fmt.Printf("%d x %s - R$%.2f\n", item.quantiCar, item.produto.nome, subtotal)
		total += subtotal
	}
	fmt.Printf("Total da compra: R$%.2f\n", total)
}

func buscarProdutoPorCodigo(estoque []Produto, codigo int) (Produto, bool) {
	for _, p := range estoque {
		if p.codProd == codigo {
			return p, true
		}
	}
	return Produto{}, false
}

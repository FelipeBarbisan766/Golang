package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Pilhas/domains"
)

func main() {
	var historico domains.Stack

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Navegador Web (Protótipo) ---")
		fmt.Println("1 - Visitar novo site")
		fmt.Println("2 - Voltar")
		fmt.Println("3 - Mostrar histórico")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		opcao, _ := reader.ReadString('\n')
		opcao = strings.TrimSpace(opcao)

		switch opcao {
		case "1":
			fmt.Print("Digite o endereço do site: ")
			site, _ := reader.ReadString('\n')
			site = strings.TrimSpace(site)
			historico.Push(site)
			fmt.Println("Você está agora em:", site)

		case "2":
			if historico.Size() <= 1 {
				fmt.Println("Não há site anterior para voltar!")
			} else {
				siteAtual := historico.Pop()
				fmt.Println("Saindo de:", siteAtual.Value)
				siteAnterior := historico.Top.Value
				fmt.Println("Voltando para:", siteAnterior)
			}

		case "3":
			fmt.Println("Histórico de navegação:")
			historico.Show()

		case "0":
			fmt.Println("Encerrando navegador...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}

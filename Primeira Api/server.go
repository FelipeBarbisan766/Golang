package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/menu", menu)
	http.HandleFunc("/estatico", arquivoEstatico)
	http.HandleFunc("/dinamico", arquivoDinamico)
	http.ListenAndServe(":8080", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, `<h1>oi, estou funcionando!<h1>`)
}
func menu(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, `
	<a href='./'> <h3>Inicio</h3></a>
	<a href='./menu'> <h3>Menu</h3></a>
	<a href='./estatico'> <h3>Estatico</h3></a>
	<a href='./dinamico'> <h3>dinamico</h3></a>
	`)
}
func arquivoEstatico(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, "estatico.html")
}
func arquivoDinamico(resp http.ResponseWriter, req *http.Request) {
	modelo, err := template.ParseFiles("dinamico.html")
	if err != nil {
		fmt.Println("erro ao utilizar a pagina de template html", err)
		resp.Header().Set("Content-Type", "text/html;charset=UTF-8")
		fmt.Fprintf(resp, `<h2> Erro ao utilizar a pagina template Html: %s</h2>
		<a href="./menu"> <h3> Ir para Menu</h3></a>
		`,err)
		return
	}
	dados := map[string]string{
		"c":"conteudo Dinamico",
		"Message": "exibindo conteudo dinamico com golang",
	}
	resp.WriteHeader(http.StatusOK)
	modelo.Execute(resp,dados)
}

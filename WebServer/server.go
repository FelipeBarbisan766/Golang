package main

import (
	"fmt"
	"html/template"
	"net/http"
	"WebServer/csvmanager"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/showCSV", showCSV)
	http.HandleFunc("/addOrganization", addOrganization)
	http.ListenAndServe(":8080", nil)
}
func index(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, `<h1>oi, estou funcionando!<h1>`)
}

func showCSV(resp http.ResponseWriter, req *http.Request) {
	manager := csvmanager.NewCSVManager()
	Orgs, err := manager.ReadAll()
	if err != nil {
		fmt.Println("Erro:", err)
	}
	
	modelo, err := template.ParseFiles("organization.html")
	if err != nil {
		fmt.Println("Erro ao utilizar a página de template HTML", err)
		resp.Header().Set("Content-Type", "text/html;charset=UTF-8")
		fmt.Fprintf(resp,
			`<h2> Erro ao utilizar página de template HTML: %s </h2>
             <a href="./menu">     <h3> Ir para Menu  </h3></a>
        `, err)
		return
	}

	resp.WriteHeader(http.StatusOK)
	modelo.Execute(resp, Orgs)

}


func addOrganization(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(resp, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Lê os dados do formulário
	err := req.ParseForm()
	if err != nil {
		http.Error(resp, "Erro ao ler dados do formulário", http.StatusBadRequest)
		return
	}

	org := csvmanager.Org{
		Index:           req.FormValue("Index"),
		OrganizationId:  req.FormValue("OrganizationId"),
		Name:            req.FormValue("name"),
		Website:         req.FormValue("website"),
		Country:         req.FormValue("country"),
		Description:     req.FormValue("description"),
		Founded:         req.FormValue("founded"),
		Industry:       req.FormValue("industry"),
		NumberEmployees: req.FormValue("numberEmployees"),
	}

	manager := csvmanager.NewCSVManager()
	err = manager.Write(org)
	if err != nil {
		http.Error(resp, "Erro ao adicionar organização", http.StatusInternalServerError)
		return
	}

	http.Redirect(resp, req, "/showCSV", http.StatusSeeOther)
}
package main

import (
	"WebServer/csvmanager"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/showCSV", showCSV)
	http.HandleFunc("/addOrganization", addOrganization)
	http.HandleFunc("/searchOrganization", searchOrganization)
	http.ListenAndServe(":8080", nil)
}
func index(resp http.ResponseWriter, req *http.Request) {
	modelo, err := template.ParseFiles("index.html")
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
	modelo.Execute(resp, nil)
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

	err := req.ParseForm()
	if err != nil {
		http.Error(resp, "Erro ao ler dados do formulário", http.StatusBadRequest)
		return
	}
	manager := csvmanager.NewCSVManager()
	orgs, err := manager.ReadAll()
	if err != nil {
		http.Error(resp, "Erro ao ler organizações", http.StatusInternalServerError)
		return
	}

	maxIndex := 0
	for _, o := range orgs {
		idx, _ := strconv.Atoi(o.Index)
		if idx > maxIndex {
			maxIndex = idx
		}
	}

	nextID := maxIndex + 1

	org := csvmanager.Org{
		Index:           strconv.Itoa(nextID),
		OrganizationId:  strconv.Itoa(nextID),
		Name:            req.FormValue("name"),
		Website:         req.FormValue("website"),
		Country:         req.FormValue("country"),
		Description:     req.FormValue("description"),
		Founded:         req.FormValue("founded"),
		Industry:        req.FormValue("industry"),
		NumberEmployees: req.FormValue("numberEmployees"),
	}

	err = manager.Write(org)
	if err != nil {
		http.Error(resp, "Erro ao adicionar organização", http.StatusInternalServerError)
		return
	}

	http.Redirect(resp, req, "/showCSV", http.StatusSeeOther)
}
func searchOrganization(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		if err.Error() == "nenhum resultado encontrado" {
			http.Error(resp, "Nenhum resultado encontrado", http.StatusNotFound)
			return
		}
		http.Error(resp, "Erro ao buscar organizações", http.StatusInternalServerError)
		return
	}

	searchTerm := req.FormValue("search")
	manager := csvmanager.NewCSVManager()
	results, err := manager.SearchByField(searchTerm)
	if err != nil {
		http.Error(resp, "Erro ao buscar organizações", http.StatusInternalServerError)
		return
	}

	modelo, err := template.ParseFiles("organization.html")
	if err != nil {
		http.Error(resp, "Erro ao utilizar template", http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	modelo.Execute(resp, results)
}

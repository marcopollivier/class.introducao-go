package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/clientes", getClientes)

	http.ListenAndServe(":8080", nil)
}

func getClientes(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	clientes := Clientes{
		Cliente{Name: "Marco"},
		Cliente{Name: "Julio"},
	}

	json.NewEncoder(w).Encode(clientes)
}

type Cliente struct {
	Name string
}

type Clientes []Cliente

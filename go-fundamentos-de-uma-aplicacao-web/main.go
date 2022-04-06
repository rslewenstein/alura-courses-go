package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camisa", Descricao: "Azul, social", Preco: 39, Quantidade: 15},
		{"Tenis", "Nike air", 150, 6},
		{"Fone", "Multilaser", 70, 45},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}

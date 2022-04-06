package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func conectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/alura_loja")
	if err != nil {
		panic(err)
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
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

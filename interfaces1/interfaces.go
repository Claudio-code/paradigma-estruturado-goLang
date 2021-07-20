package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

/// ----- >>>>> AS INTERFACES SÃO IMPLMENTAS AUTOMATICAMENTE NO GO

func (p pessoa) toString() string {
	return fmt.Sprintf("%s = %s", p.nome, p.sobrenome)
}

func (p produto) toString() string {
	return fmt.Sprintf("%s = %2.f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = produto{
		nome:  "kraudeo",
		preco: 2.22,
	}

	imprimir(coisa)

	coisa = pessoa{
		nome:      "antonia",
		sobrenome: "maria",
	}

	imprimir(coisa)
}

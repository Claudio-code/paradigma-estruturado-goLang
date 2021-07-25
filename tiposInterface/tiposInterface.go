package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisas interface{}
	fmt.Println(coisas)

	coisas = 2
	fmt.Println(coisas)

	coisas = "nvoas"
	fmt.Println(coisas)

	type dinamico interface{}

	var coisa2 dinamico = "casas"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)
}

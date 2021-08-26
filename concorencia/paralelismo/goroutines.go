package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	go fale("maria", "Oi...dwq", 10)
	go fale("joao", "passa de faze ai", 10)

	time.Sleep(time.Second * 10)
	fmt.Println("isso é tudo pessoal")
}

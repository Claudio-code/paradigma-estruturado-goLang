package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar novos metodos
}

type bwm7 struct{}

func (b bwm7) ligarTurbo() {
	fmt.Println("Ligando o turbo...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("fazendo a baliza :)")
}

func ligarCarro(c esportivoLuxuoso) {
	c.fazerBaliza()
	c.ligarTurbo()
}

func main() {
	bwm7 := bwm7{}

	ligarCarro(bwm7)
}

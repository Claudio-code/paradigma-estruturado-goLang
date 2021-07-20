package main

import "fmt"

type esportivo interface {
	ligarturbo()
}

type ferrari struct {
	modelo      string
	turboStatus bool
	velocidade  int
}

func (f *ferrari) ligarturbo() {
	fmt.Println(f)
	f.turboStatus = true
}

func main() {
	carro1 := ferrari{
		modelo:      "novo 1",
		turboStatus: false,
		velocidade:  22,
	}
	var carro2 esportivo = &ferrari{
		modelo:      "novo 2",
		turboStatus: false,
		velocidade:  22,
	}

	carro1.ligarturbo()
	fmt.Println(carro1)

	fmt.Println("-- ----- -- - -- - -")

	carro2.ligarturbo()
	fmt.Println(carro2)
}

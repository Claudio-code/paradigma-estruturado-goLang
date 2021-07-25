package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{
		ID:    1,
		Nome:  "faca",
		Preco: 121.2,
		Tags: []string{
			"Promoção",
			"Casa",
		},
	}

	// convertendo o struct para json
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	//////////////////////////////////////////////////////////////////////////////

	var p2 produto
	jsonString := `{"id":2,"nome":"caneta","preco":2.22,"tags":["papelaria", "importados"]}`

	// convertendo o json para struct
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
}

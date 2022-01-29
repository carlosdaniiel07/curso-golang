package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome string
	sobrenome string
}

func (p Pessoa) getNomeCompleto() string {
	return fmt.Sprintf("%s %s", p.nome, p.sobrenome)
}

// Quando o método produz alguma alteração na própria variável é necessarío
// definir o receiver para receber um ponteiro!
func (p *Pessoa) setNomeCompleto(valor string) {
	partes := strings.Split(valor, " ")

	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	me := Pessoa {
		nome: "Fulano",
		sobrenome: "da Silva",
	}

	fmt.Println(me.getNomeCompleto())

	me.setNomeCompleto("Ciclano Fulano")

	fmt.Println(me.getNomeCompleto())
}

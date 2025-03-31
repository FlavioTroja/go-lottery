package main

import "fmt"

type Studente struct {
	Nome string
	Età int 
	Punteggio int
}

func (s Studente) Saluta() string { 
	return "Ciao ciao " + s.Nome
}

func main() {

	s := Studente {
		Nome: "Flavio Troia",
		Età: 47,
		Punteggio: 5,
	}

	saluto := s.Saluta()

	fmt.Println(saluto)
}
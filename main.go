package main

import (
	"errors"
	"fmt"
)

type Studente struct {
	Nome string
	Età int 
	Punteggio int
}

func (s Studente) Saluta() string { 
	return "Ciao ciao " + s.Nome
}

func somma(a int, b int) int {
	return a + b
}

func divisione (a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisione per zero")
	}
	return a/b, nil
}

func main() {

	s := Studente {
		Nome: "Flavio Troia",
		Età: 47,
		Punteggio: 5,
	}

	saluto := s.Saluta()

	fmt.Println(saluto.Saluta("Flavio"))

	fmt.Println(saluto)
	fmt.Println("Età:", s.Età)
    fmt.Println("Punteggio:", s.Punteggio)	

	fmt.Println("Risultato somma: ", somma( 5, 3 ))

	div, err := divisione(4, 2)
	if err != nil {
		fmt.Println("Errore divisione: ", err)
	} else {
		fmt.Println("Risultato divisione: ", div)
	}

	_, err = divisione(4, 0)
	if err != nil {
		fmt.Println("Errore divisione: ", err)
	} 
}
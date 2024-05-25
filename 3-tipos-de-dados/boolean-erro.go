package main

import ( 
	"fmt"
	"errors"
)


func main () {
	var booleano1 bool = true 
	fmt.Println(booleano1)

	var booleano2 bool = false 
	fmt.Println(booleano2)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}
package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int64 = -100000000000
	fmt.Println(numero)


	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias
	// int 32 == rune 

	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)


	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	
	var numeroReal2 float64 = 1233242342.4343
	fmt.Println(numeroReal2)

	numeroReal3 := 23132.43
	fmt.Println(numeroReal3)

	var str string = "AlALALLALA"
	fmt.Println(str)
	

	str2 := "Texto2"
	fmt.Println(str2)

	// Não existe char em go , aspas simples declarado na forma abaixo mostra qual seria o número desse caracter na tabela ASCII
	char := 'B'
	fmt.Println(char)


	var texto int
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)
	
	
	
	//error é um tipo , é um tipo de dado "tipo com valor 0"
	var erro error	
	fmt.Println(erro)

	var erro2 error = errors.New("Error interno")	
	fmt.Println(erro2)

}

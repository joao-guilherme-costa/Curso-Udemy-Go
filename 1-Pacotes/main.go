package main

import (
	"fmt"
	"modulo/auxiliar"
)

//go mod tidy -> tira todas as dependencias que não estão sendo usadas


func main() {

	fmt.Println("Escrevendo do pacote main")
	auxiliar.Escrever() // Chamou uma função que estava em outro arquivo pois foi importado o pacote e a função estava com letra maíscula
}

package main

import "fmt"

func main() {
	var nota1, nota2 float32 
	

	fmt.Print("Bem-Vindo ao Programa de Soma de DUAS notas.")
	fmt.Println("\n")
	fmt.Print(" NOTA - 1:  ")
	fmt.Scanln(&nota1)
	fmt.Print(" NOTA - 2:  ")
	fmt.Scanln(&nota2)
	fmt.Println(" ")


	fmt.Printf("SOMA DAS NOTAS: %.2f \n", nota1 + nota2)
}
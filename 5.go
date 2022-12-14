package main

import "fmt"

func main() {
	var nota1, nota2 float32 
	

	fmt.Print("Bem-Vindo ao Programa de MÉDIA de notas.")
	fmt.Println("\n")
	fmt.Print(" NOTA - 1:  ")
	fmt.Scanln(&nota1)
	fmt.Print(" NOTA - 2:  ")
	fmt.Scanln(&nota2)
	fmt.Println(" ")
	var soma float32 = nota1 + nota2
	var media float32 = soma / 2


	fmt.Printf("a MÉDIA entre as nota é de: %.2f \n", media)
}
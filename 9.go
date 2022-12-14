package main

import "fmt"

func main() {
	var  num1, num2, num3 float32


	fmt.Print("Bem-Vindo ao programa de multiplicação de 3 numeros ")
	fmt.Println("\n")	


	fmt.Print("Numero 1: ")
	fmt.Scanln(&num1)
	fmt.Print("Numero 2: ")
	fmt.Scanln(&num2)
	fmt.Print("Numero 3: ")
	fmt.Scanln(&num3)
	fmt.Println(" ")	

	
	fmt.Printf("o total é de: %.2f \n", num1 * num2 * num3)
}
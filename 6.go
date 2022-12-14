package main

import "fmt"

func main() {
	var num1 float32
	var num2 float32

	fmt.Print(" NUM - 1:  ")
	fmt.Scanln(&num1)
	fmt.Print(" NUM - 2:  ")
	fmt.Scanln(&num2)

	var soma float32 = num1 + num2
	var subtracao float32 = num1 - num2
	var multiplicacao float32 = num1 * num2
	var divisao float32 = num1 / num2

	fmt.Printf("a SOMA entre os numeros é de: %.2f \n", soma)
	fmt.Printf("a SUBTRAÇÃO entre os numeros é de: %.2f \n", subtracao)
	fmt.Printf("a MULTIPLICAÇÃO entre os numeros é de: %.2f \n", multiplicacao)
	fmt.Printf("a DIVISÃO entre os numeros é de: %.2f \n", divisao)

	
	
	
}
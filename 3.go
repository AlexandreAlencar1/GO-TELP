package main

import "fmt"

func main() {

	var  ano, dia, hora, minuto, segundo int
	
	fmt.Printf("Bem-Vindo Ao conversor de ANOS para MINUTOS E SEGUNDOS: ")
	fmt.Printf("\n")
	fmt.Printf("ENTRE com os ANOS que deseja: ")
	fmt.Scanf("%d", &ano)

	dia = ano * 365
	hora = dia * 24
	minuto = hora * 60
	segundo = minuto * 60

	fmt.Printf("\n \n")
	fmt.Printf("\n\n%d ANOS para \n%d MINUTOS E \n%d SEGUNDOS.\n",ano, minuto, segundo,  )
}
package main

import "fmt"

func main() {

	var segundo,minuto,segundo2,restante int
	

	fmt.Printf("entre com os segundos: ")
	fmt.Scanf("%d", &segundo)

	restante = segundo % 3600
	minuto = segundo / 60
	segundo2 = restante % 60
	fmt.Printf("%d digitado pelo usuario: || conversão igual: %d minuto e %d segundo\n", segundo, minuto, segundo2)
}
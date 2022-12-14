package main

import "fmt"

func looping(media, num_maior int) (int, int) {
	var num int
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		if num > num_maior {
			num_maior = num
		}
		media += num
	}
	media = media / 10
	return media, num_maior
}

func main() {
	var media, num_maior int
	fmt.Println("Entre com 10 numeros para saber a media e qual o maior deles:")
	media, num_maior = looping(media, num_maior)
	fmt.Println("a média:", media)
	fmt.Println("o maior:", num_maior)
}
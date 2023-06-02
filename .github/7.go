package main

import "fmt"

//7. Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que
//informe um número inteiro.
//Adicione esse número ao Slice apenas se ele não estiver presente.
//Imprima o Slice resultante.

func main() {
	var slice = []int{2, 3}
	var num int
	presente := false

	fmt.Print("Informe um número inteiro: ")
	fmt.Scan(&num)
	for _, n := range slice {
		if num == n {
			presente = true
		}
	}
	if presente == false {
		slice = append(slice, num)
	}
	fmt.Print(slice)
}

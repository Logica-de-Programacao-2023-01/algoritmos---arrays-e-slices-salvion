package main

import "fmt"

//4. Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
//Em seguida, imprima o Slice e a soma dos valores armazenados.

func main() {
	soma := 0
	var tam int
	var elemento int
	var slice []int

	fmt.Println("Tamanho do slice a ser criado:")
	fmt.Scan(&tam)

	for i := 0; i < tam; i++ {
		fmt.Println("elemento", fmt.Sprint(i+1), ":")
		fmt.Scan(&elemento)
		slice = append(slice, elemento)
		soma += elemento
	}

	fmt.Printf("O slice criado é: %d, e a soma de seus valores é igual a %d", slice, soma)
}

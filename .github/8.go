package main

import "fmt"

//8. Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor.
//Remova todas as ocorrências desse valor no Slice e imprima o resultado.

func main() {
	var numeros = []int{1, 2, 3, 2, 5, 2, 7, 8}
	var num int
	fmt.Print("Informe um valor: ")
	fmt.Scan(&num)

	for i := 0; i < len(numeros); i++ {
		if numeros[i] == num {
			numeros = append(numeros[:i], numeros[i+1:]...)
		}
	}
	fmt.Print(numeros)
}

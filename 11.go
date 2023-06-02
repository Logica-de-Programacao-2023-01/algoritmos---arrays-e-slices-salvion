package main

import "fmt"

//11. Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas.
//Em seguida, solicite ao usuário que informe um índice de linha e
//outro de coluna e imprima o valor armazenado nessa posição da matriz.

func main() {
	array := [2][3]int{{2, 3, 4}, {1, 5, 6}}
	var li, col int
	fmt.Print("Linha: ")
	fmt.Scanln(&li)
	fmt.Print("Coluna: ")
	fmt.Scanln(&col)
	fmt.Print(array[li][col])
}

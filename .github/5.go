package main

import "fmt"

//5. Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas.
//Solicite ao usuário que informe os valores de cada elemento da matriz.
//Em seguida, imprima a matriz resultante.

func main() {
	const linhas int = 3
	const colunas int = 2

	var matriz [linhas][colunas]int
	var val int
	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			fmt.Printf("Valor da linha %d, coluna %d: ", i, j)
			fmt.Scan(&val)
			matriz[i][j] = val
		}
	}
	fmt.Print(matriz)
}

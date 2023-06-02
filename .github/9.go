package main

import "fmt"

//9. Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um
//número que será adicionado em todas as posições do Array. Imprima o Array resultante.

func main() {
	var lista [6]float64
	var n float64

	fmt.Print("número de substituição: ")
	fmt.Scan(&n)
	for i := 0; i < 6; i++ {
		lista[i] = n
	}
	fmt.Print(lista)
}

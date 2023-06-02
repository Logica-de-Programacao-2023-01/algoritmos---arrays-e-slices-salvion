package main

import "fmt"

//10. Crie um Slice de inteiros
//com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.

func main() {
	valores := []int{5, 78, 3, 2, 1, 80, 3, 0, 3, -1, 3, 23}
	min := valores[0]
	max := valores[0]
	for _, n := range valores {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Print(min, max)
}

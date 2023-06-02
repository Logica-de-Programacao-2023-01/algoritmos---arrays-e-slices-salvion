package main

import "fmt"

//12. Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que
//contenha apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.

func main() {
	array := [5]int{1, 2, 3, 0, 5}
	impares := []int{}
	for _, n := range array {
		if n%2 != 0 {
			impares = append(impares, n)
		}
	}
	fmt.Print(impares)
}

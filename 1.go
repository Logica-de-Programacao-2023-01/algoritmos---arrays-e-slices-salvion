package main

import "fmt"

//1. Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.

func main() {
	var array = [3]int{1, 2, 3}
	soma := 0

	for _, n := range array {
		soma += n
	}
	fmt.Print(soma)
}

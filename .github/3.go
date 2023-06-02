package main

import "fmt"

//3. Crie um Array de floats com 4 elementos e
//calcule o produto dos valores armazenados no Array.

func main() {
	var prod float32
	var array = [4]float32{2, 3, 1, 2}
	prod = array[0]
	for _, val := range array {
		prod = prod * val
	}
	prod = prod / array[0]
	fmt.Print(prod)
}

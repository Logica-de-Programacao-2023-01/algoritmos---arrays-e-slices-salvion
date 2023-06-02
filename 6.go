package main

import "fmt"

//6. Crie um Array de inteiros com 10 elementos.
//Em seguida, solicite ao usuário que informe um valor e verifique
//se esse valor está presente no Array. Imprima o resultado da busca.

func main() {
	var array = [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	var val int
	achou := false

	fmt.Print("Informe um valor: ")
	fmt.Scan(&val)
	for _, n := range array {
		if val == n {
			achou = true
		}
	}
	if achou == true {
		fmt.Print("Encontrado.")
	} else {
		fmt.Print("Não encontrado.")
	}
}

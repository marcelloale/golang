// A ramificação com `if` e `else` em Go é direta.

package main

import "fmt"

func main() {

	// Aqui está um exemplo básico.
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é ímpar")
	}

	// Você pode ter uma condição `if` sem um else.
	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	}

	// Uma condição pode preceder condições; qualquer variável
	// declarada em uma condição está disponível em todas
	// as ramificações.
	if num := 9; num < 0 {
		fmt.Println(num, "é negativo")
	} else if num < 10 {
		fmt.Println(num, "tem 1 dígito")
	} else {
		fmt.Println(num, "tem vários dígitos")
	}
}

// Note que você não precisa de parênteses nas condições
// em Go, mas os colchetes são necessários.

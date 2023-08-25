// `for` é o único tipo de loop do Go.
// Veja esses três tipos básicos de loops `for`.

package main

import "fmt"

func main() {

	// O tipo mais básico, com uma única condição.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println() //só para pular uma linha
	// Um inicialização clássica como estamos mais acostumados.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println() //só para pular uma linha
	// `for` sem uma condição irá fazer o loop infinitamente
	// até um `break` ou um `return`.
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println()
	// Você também pode `continuar` para a próxima iteração
	// do loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

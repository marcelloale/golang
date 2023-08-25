// No Go, uma _matriz_ é uma seqüência de elementos de um
// tamanho específico.

package main

import "fmt"

func main() {

	// Aqui nós criamos uma matriz `a` que irá manter
	// exatamente 5 `int`eiros. O tipo de elementos e tamanho
	// são ambos parte de um tipo de matriz. Por padrão um
	// array é de valor-zero, onde cada `int` significa `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// Nós podemos definir um valor em um índice usando a
	// sintaxe `matriz[index] = valor`, e pegar o valor
	// com `matriz[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// O `len` padrão retorna o tamanho de uma matriz.
	fmt.Println("len:", len(a))

	// Use esta sintaxe para declarar e inicializar uma
	// matriz em uma linha.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Tipos de matrizes são unidimensionais, mas você
	// pode compor tipos para construir estruturas de dados
	// multidimensionais.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Arrays (o comprimento especifica o número de elementos a serem armazenados na matriz. Em Go, os arrays têm um comprimento fixo.)
	arrdef := [5]int{1, 2, 3, 4} //  comprimentos definido
	fmt.Println("Array definido:", arrdef)
	fmt.Println("Array definido capacity:", cap(arrdef))
	arrdefvaz := [5]int{} //  comprimentos definido vazio
	fmt.Println("Array definido vazio:", arrdefvaz)
	fmt.Println("Array definido vazio capacity:", cap(arrdefvaz))
	arrinf := [...]int{4, 5, 6, 7, 8} //  comprimentos inferido
	fmt.Println("Array inferido:", arrinf)
	fmt.Println("Array inferido capacity:", cap(arrinf))
	arrinfvaz := [...]int{} //  comprimentos inferido vazio
	fmt.Println("Array inferido vazio:", arrinfvaz)
	fmt.Println("Array inferido vazio capacity:", cap(arrinfvaz))
}

// _Slices_ são tipos de dados fundamentais no Go, fornecendo uma
// interface mais poderosa para seqüências de matrizes.

package main

import "fmt"

func main() {

	// Ao contrário das matrizes, os slices são digitados apenas
	// pelos elementos que eles contêm (e não o número de elementos).
	// Para criar um slice vazio com tamanho diferente de zero, use
	// o `make` padrão. Aqui nós fizemos um slice de uma `string` de
	// tamanho `3` (inicialmente com valor zero).

	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Podemos definir e obter assim como as matrizes.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// O `len` retorna o tamanho do slice, como esperado.
	fmt.Println("len:", len(s))

	// Além das operações básicas, os slices suportam
	// muito mais, o que os tornam mais ricos do que as
	// matrizes. Uma é o padrão `append`, na qual
	// retorna um slice contendo um ou mais novos valores.
	// Note que precisamos aceitar um valor de retorno do
	// append assim como podemos obter um novo valor do slice.

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices podem ser também copiados. Aqui nós criamos um
	// slice vazio `c` do mesmo tamanho de `s` e copiamos
	// o `s` no `c`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices suportam um operador "slice" com a sintaxe

	// Slices (pode aumentar e diminuir conforme você desejar)
	slice := []int{10, 20, 30}
	fmt.Println("Slice:", slice)
	fmt.Println("Slice capacity:", cap(slice))

}

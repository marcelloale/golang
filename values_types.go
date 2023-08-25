// variables_types.go
package main

import "fmt"

func main() {
	//Concatena as duas strings com o sinal de +
	fmt.Println("go" + "lang") //Saída: golang
	frase := "Olá, mundo!"
	fmt.Println("Frase: ", frase+" tamanho frase: ", len(frase))
	fmt.Println("Primeiro e ultimo char:", string(frase[0]), string(frase[(len(frase)-1)]))
	fmt.Println("a = ", 'a') //Saída: ascii

	// Operadores Aritméticos
	fmt.Println("1 + 1 = ", 1+1) // Saída: 1 + 1 = 2
	fmt.Println("2 - 1 = ", 2-1)
	fmt.Println("2 * 3 = ", 2*3)
	fmt.Println("7 / 3 =", 7/3)          //Divisão de inteiros Saída: 7 / 3 = 2
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0) //Saída: 7.0 / 3.0 = 2.3333333333333335
	fmt.Println("7 % 3 =", 7%3)          //Operação de resto/mod Saída: 7 % 3 = 1
	x := 10
	fmt.Println("x = ", x)
	x++
	fmt.Println("x++ = ", x)
	x--
	fmt.Println("x-- = ", x)

	// Operadores de Comparação
	fmt.Println("1 == 1:", 1 == 1)
	fmt.Println("3 < 2:", 3 < 2)
	fmt.Println("4 >= 4:", 4 >= 4)
	// Operadores Lógicos
	fmt.Println(true && false) //AND Saída: false
	fmt.Println(true || false) //OR  Saída: true
	fmt.Println(!true)         //NOT Saída: false

	// Ponteiro
	var ptr *int
	numero := 7
	ptr = &numero
	fmt.Println("Ponteiro:", *ptr)

	// Maps (A ordem dos elementos do mapa definida no código é diferente da forma como eles são armazenados. Os dados são armazenados de forma a ter uma recuperação de dados eficiente do mapa.)
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	// Os tipos de chave inválidos são: Slices, Maps, Functions. Os valores do mapa podem ser de qualquer tipo.

	// Structs
	type Pessoa struct {
		Nome  string
		Idade int
	}
	pessoa := Pessoa{Nome: "Alice", Idade: 30}
	fmt.Println("Struct:", pessoa)

}

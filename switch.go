// _instruções Switch_ expressam condições através de
// vários ramos.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Aqui está um `switch` básico.
	i := 2
	fmt.Print("escreva ", i, " como ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	// Você pode usar vírgulas para separar várias expressões
	// na mesma instrução `case`. Nós usamos também, o case opcional
	// `default` nesse exemplo.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("é final de semana")
	default:
		fmt.Println("é um dia da semana")
	}

	// `switch` sem uma expressão é uma forma alternativa
	// para expressar a lógica if/else. Aqui nós também podemos
	// mostrar como as expressões `case` podem ser não-constantes.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("é antes do meio dia")
	default:
		fmt.Println("é depois do meio dia")
	}

	//Uma declaração switch de tipos compara tipos em vez de valores.
	//Você pode usar isso para descobrir o tipo de um valor de interface.
	//Neste exemplo, a variável t terá o tipo correspondente à sua cláusula.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// todo: tipos de switches

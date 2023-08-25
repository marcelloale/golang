/*
O Hipermercado Tabajara está com uma promoção de carnes que é imperdível. Confira:
Para atender a todos os clientes, cada cliente poderá levar apenas um dos tipos de carne da promoção, porém não há limites para a quantidade de carne por cliente.
Se compra for feita no cartão Tabajara o cliente receberá ainda um desconto de 5% sobre o total da compra.
Escreva um programa que peça o tipo e a quantidade de carne comprada pelo usuário e gere um cupom fiscal,
contendo as informações da compra: tipo e quantidade de carne, preço total, tipo de pagamento, valor do desconto e valor a pagar.
                      Até 5 Kg           Acima de 5 Kg
File Duplo      R$ 4,90 por Kg          R$ 5,80 por Kg
Alcatra         R$ 5,90 por Kg          R$ 6,80 por Kg
Picanha         R$ 6,90 por Kg          R$ 7,80 por Kg
*/

package main

import "fmt"

func main() {
	var tipoCarne uint16
	var qtdCarne float32
	var cartaoTaba uint16
	var kgCarne float32
	var valorCompra float32

	fmt.Println("Digite o tipo de carne:")
	fmt.Scan(&tipoCarne)

	fmt.Println("Digite a quantidade:")
	fmt.Scan(&qtdCarne)

	fmt.Println("Pagamento com o cartão Tabajara?: 1=Sim 2=Nao")
	fmt.Scan(&cartaoTaba)

	switch tipoCarne {
	case 1:
		fmt.Println("Tipo de carne escolhido: File Duplo")
		if qtdCarne > 5 {
			kgCarne = 4.90
		} else {
			kgCarne = 5.80
		}
	case 2:
		fmt.Println("Tipo de carne escolhido: Alcatra")
		if qtdCarne > 5 {
			kgCarne = 5.90
		} else {
			kgCarne = 6.80
		}
	case 3:
		fmt.Println("Tipo de carne escolhido: Picanha")
		if qtdCarne > 5 {
			kgCarne = 6.90
		} else {
			kgCarne = 7.80
		}
	}

	if cartaoTaba == 1 {
		valorCompra = (kgCarne * qtdCarne) - ((kgCarne * qtdCarne) * 0.05)
	} else {
		valorCompra = kgCarne * qtdCarne
	}

	fmt.Println("Valor total da Compra: ", valorCompra)
}

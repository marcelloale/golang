// Faça um programa que peça o tamanho de um arquivo para download (em MB) e a velocidade de um link de Internet (em Mbps), calcule e informe o tempo aproximado de download do arquivo usando este link (em minutos)

package main

import "fmt"

func main() {
	var tamanhoArquivo float64
	var velocidadeLink float64

	fmt.Println("Digite o tamanho do arquivo em MB")
	fmt.Scan(&tamanhoArquivo)

	fmt.Println("Digite a velocidade do link de Internet em Mbps")
	fmt.Scan(&velocidadeLink)

	velocidadeDownload := (tamanhoArquivo / (velocidadeLink / 8)) / 60
	fmt.Printf("O Tempo aproximado de download do arquivo usando este link: %.2f minutos.\n", velocidadeDownload)

}

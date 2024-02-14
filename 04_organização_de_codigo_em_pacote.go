// Em Go (Golang), a organização de código em pacotes é uma prática fundamental para estruturar e modularizar projetos de software. 
// Os pacotes em Go são coleções de arquivos com código fonte relacionado que podem ser reutilizados em diferentes partes de um programa ou em diferentes projetos.
// Aqui está um exemplo simples de como organizar o código em pacotes em Go:

// Suponha que você tenha um projeto de calculadora com as seguintes funcionalidades: Adição / Subtração / Multiplicação / Divisão
// Você pode organizar seu código em pacotes da seguinte forma:

	// pacote main: Ponto de entrada do programa
	package main

	// Importa os pacotes necessários
	import (
		"fmt"
		"minhacalculadora/calculos"
	)

	// Função principal
	func main() {
		// Executa operações de calculadora
		soma := calculos.Somar(10, 5)
		fmt.Println("10 + 5 =", soma)

		subtracao := calculos.Subtrair(10, 5)
		fmt.Println("10 - 5 =", subtracao)

		multiplicacao := calculos.Multiplicar(10, 5)
		fmt.Println("10 * 5 =", multiplicacao)

		divisao := calculos.Dividir(10, 5)
		fmt.Println("10 / 5 =", divisao)
	}

// Agora, vamos organizar as funcionalidades da calculadora em um pacote chamado "calculos":

	// pacote calculos: Contém funções para cálculos matemáticos
	package calculos

	// Função para adição
	func Somar(a, b int) int {
		return a + b
	}

	// Função para subtração
	func Subtrair(a, b int) int {
		return a - b
	}

	// Função para multiplicação
	func Multiplicar(a, b int) int {
		return a * b
	}

	// Função para divisão
	func Dividir(a, b int) float64 {
		if b == 0 {
			return 0
		}
		return float64(a) / float64(b)
	}

// Neste exemplo:
// O código principal está no pacote main. Este pacote é especial e contém a função main(), que é o ponto de entrada do programa.
// As funções de cálculo (Somar, Subtrair, Multiplicar, Dividir) são organizadas no pacote calculos. Este pacote fornece funcionalidades reutilizáveis que podem ser importadas e usadas em outros lugares.
// Usamos o mecanismo de importação do Go para acessar as funções do pacote calculos no pacote main.
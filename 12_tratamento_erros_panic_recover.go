// Em Go, panic e recover são mecanismos utilizados para lidar com erros excepcionais, semelhantes às exceções em outras linguagens de programação. 
// Aqui está uma breve explicação de cada um:

// 1- panic: É uma função interna em Go que é usada para interromper abruptamente a execução de um programa. 
// Quando uma função panic é chamada, ela causa o encerramento imediato da execução da função atual e, subsequentemente, das funções que a chamaram, até que o programa seja encerrado. 
// panic é comumente usado para indicar um erro fatal que não pode ser tratado naquele momento, como uma divisão por zero ou uma tentativa de acessar um índice inválido em um slice.

// 2- recover: É uma função que permite ao programa se recuperar de um panic. 
// recover só é útil quando chamado dentro de uma função diferida (defer), caso contrário, não terá efeito. 
// Quando uma chamada recover é feita dentro de uma função diferida que está sendo executada em uma função onde ocorreu um panic, o recover interrompe o fluxo de pânico e retorna o valor passado para panic. 
// Isso permite que o programa continue a execução em vez de ser terminado abruptamente.

// Aqui está um exemplo básico de como panic e recover podem ser usados juntos:

	package main

	import "fmt"

	func main() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()

		divideByZero()
		fmt.Println("This line will never be reached.")
	}

	func divideByZero() {
		numerator := 10
		denominator := 0
		result := numerator / denominator // This will cause a panic
		fmt.Println(result)
	}

// Neste exemplo, a função divideByZero tenta dividir numerator por denominator, que é zero. Isso causará um panic de divisão por zero. 
// No entanto, o defer no main vai chamar a função anônima após o panic, que por sua vez chama recover, recuperando assim o controle e permitindo que o programa imprima uma mensagem de erro em vez de terminar abruptamente.
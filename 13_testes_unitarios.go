// Os testes unitários são focados em testar unidades individuais de código, como funções ou métodos, isoladamente do restante do sistema.
// Geralmente, esses testes são escritos pelo desenvolvedor para verificar se cada unidade de código funciona conforme o esperado. 
// Em Go, os testes unitários são escritos no mesmo pacote que o código a ser testado, no arquivo _test.go.

//Suponha que você tenha uma função em um pacote chamado "math" que calcula a soma de dois números:
	
	package math

	func Sum(a, b int) int {
		return a + b
	}

// Agora, você pode escrever testes unitários para essa função em um arquivo chamado "math_test.go":

	package math

	import "testing"

	func TestSum(t *testing.T) {
		// Test case 1
		result := Sum(2, 3)
		expected := 5
		if result != expected {
			t.Errorf("Sum(2, 3) = %d; expected %d", result, expected)
		}

		// Test case 2
		result = Sum(-1, 1)
		expected = 0
		if result != expected {
			t.Errorf("Sum(-1, 1) = %d; expected %d", result, expected)
		}
	}

// Para executar esses testes, você pode usar o comando go test na linha de comando na pasta do pacote:

	go test
// Isso executará todos os testes no pacote atual e fornecerá a saída correspondente.
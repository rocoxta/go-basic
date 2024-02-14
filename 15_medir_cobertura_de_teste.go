// Medir a cobertura de testes em uma linguagem como Go (também conhecida como Golang) é uma prática comum 
// para avaliar o quanto do código-fonte do seu programa está sendo exercitado pelos testes automatizados. 
// A cobertura de teste é geralmente expressa como uma porcentagem do código que foi executado durante a execução dos testes.

// Para medir a cobertura de testes em Go, você pode usar a ferramenta go test junto com a flag -cover. 
// Aqui está um exemplo de como você pode medir a cobertura de testes para um pacote em Go:
// Suponha que você tenha um pacote chamado math com o seguinte código:

	package math

	func Add(a, b int) int {
		return a + b
	}

	func Subtract(a, b int) int {
		return a - b
	}

// E um conjunto de testes para este pacote:

	package math

	import "testing"

	func TestAdd(t *testing.T) {
		result := Add(2, 3)
		if result != 5 {
			t.Errorf("Expected 5 but got %d", result)
		}
	}

	func TestSubtract(t *testing.T) {
		result := Subtract(5, 3)
		if result != 2 {
			t.Errorf("Expected 2 but got %d", result)
		}
	}

// Para medir a cobertura de testes para este pacote, você pode executar o seguinte comando no terminal:

	go test -cover

// Isso exibirá a porcentagem de cobertura de testes para o pacote math. 
// Se você deseja uma cobertura de teste mais detalhada, pode usar a flag -coverprofile para gerar um perfil de 
// cobertura em um arquivo e, em seguida, usar a ferramenta go tool cover para visualizar o perfil de cobertura em HTML:

	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

// Isso abrirá uma página HTML em seu navegador, mostrando quais partes do código estão cobertas pelos testes e quais não estão. 
// Isso pode ajudar a identificar áreas do código que precisam de testes adicionais.

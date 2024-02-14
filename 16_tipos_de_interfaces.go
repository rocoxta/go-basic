package main
// A linguagem de programação Go (ou Golang) oferece vários tipos de interfaces que permitem uma programação mais flexível e genérica. 
// Aqui estão os principais tipos de interfaces em Go, juntamente com exemplos simples:
import (
    "fmt"
    "math"
)

// Empty Interface: Uma interface vazia em Go pode representar qualquer tipo de valor.

	func main() {
		var any interface{}
		any = 42
		fmt.Println(any) // Output: 42
		any = "hello"
		fmt.Println(any) // Output: hello
	}

// Named Interface: Interfaces definidas pelo usuário que especificam um conjunto de métodos que um tipo deve implementar para satisfazer a interface.

	type Shape interface {
		Area() float64
	}

	type Circle struct {
		Radius float64
	}

	func (c Circle) Area() float64 {
		return math.Pi * c.Radius * c.Radius
	}

// Embedded Interfaces: Interfaces podem ser embutidas em outras interfaces para compor funcionalidades.

	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	type Closer interface {
		Close() error
	}

	type ReadCloser interface {
		Reader
		Closer
	}

// Interfaces with Methods: Interfaces vazias podem ter métodos definidos, o que permite restrições adicionais sobre os tipos que podem satisfazer a interface.

	type Stringer interface {
		String() string
	}

	type Person struct {
		Name string
	}

	func (p Person) String() string {
		return "Person: " + p.Name
	}

// Type Interfaces: Interfaces que especificam um conjunto de métodos, mas não exigem que o tipo explicitamente declare que está implementando a interface.
	func PrintString(s Stringer) {
		fmt.Println(s.String())
	}

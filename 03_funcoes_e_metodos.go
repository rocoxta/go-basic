package main
	// Funções e métodos: fuc.
import "fmt"

	// Função básica: É usado para definir uma função em Go.
func soma(a, b int) int {
    return a + b
}

	// Definição de estrutura: É usado para definir uma função em Go.
type Retangulo struct {
    altura, largura float64
}

	// Método em estrutura: É possível definir métodos em estruturas em Go.
func (r Retangulo) Area() float64 {
    return r.altura * r.largura
}

	// Método com ponteiros: Em Go, é comum definir métodos que operam em ponteiros para estruturas.
func (r *Retangulo) Escala(fator float64) {
    r.altura *= fator
    r.largura *= fator
}

	// Função variática: São funções que aceitam um número variável de argumentos.
func somaVariatica(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

	// Função anônima e closure: Em Go, você pode definir funções sem nome (funções anônimas) e usá-las como closures.
func funcaoAnonima() {
    add := func(a, b int) int {
        return a + b
    }
    result := add(3, 4) // result será 7
    fmt.Println(result)
}

	// Função com defer: A declaração defer adia a execução de uma função até o final do escopo atual.
func minhaFuncao() {
    defer fmt.Println("Execução adiada até o final da função")
    fmt.Println("Esta será impressa primeiro")
}

	// Função recursiva: Go suporta recursão, permitindo que funções chamem a si mesmas.
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    // Exemplo de uso das funções e métodos explicados: 
    fmt.Println("Soma:", soma(3, 4))

    r := Retangulo{altura: 3, largura: 4}
    fmt.Println("Área do retângulo:", r.Area())

    r.Escala(2)
    fmt.Println("Nova altura do retângulo:", r.altura)
    fmt.Println("Nova largura do retângulo:", r.largura)

    fmt.Println("Soma variática:", somaVariatica(1, 2, 3, 4, 5))

    funcaoAnonima()

    minhaFuncao()

    fmt.Println("Fibonacci de 5:", fibonacci(5))
}

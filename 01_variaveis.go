package main
	// Variáveis: int, float, string, bool, arrays, slices, maps, structs, ponteiros e função.
import "fmt"
	
func main() {
    // int: A definição do int puro, pode variar de acordo com a plataforma onde o código é compilado. Em alguns sistemas, int pode ter 32 bits de tamanho (usualmente chamado de int32), enquanto em outros pode ter 64 bits (int64).
	// Se você precisa de um tipo de inteiro com um tamanho específico e precisa garantir a portabilidade do seu código entre diferentes arquiteturas, é recomendável usar tipos com tamanhos definidos, como:
	// int8 (-128 a 127), int16 (-32768 a 32767), int32 (-2147483648 a 2147483647), int64 (-9223372036854775808 a 9223372036854775807).
    var idade int = 25
    
    // float: Tipos para representar números decimais, como float32 e float64.
    var altura float64 = 1.75
    
    // string: Tipo para representar sequências de caracteres.
    var nome string = "João"
    
    // bool: Tipo para representar valores booleanos, verdadeiro ou falso.
    var ativo bool = true
    
    // arrays: Estruturas de dados que representam coleções de elementos de tamanho fixo.
    var numeros [5]int = [5]int{1, 2, 3, 4, 5}
    
    // slices: Estruturas de dados semelhantes a arrays, mas de tamanho dinâmico.
    numerosSlice := []int{1, 2, 3, 4, 5}
    
    // maps: Estruturas de dados para mapear chaves a valores.
    idadePorNome := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
    
    // structs: Tipos de dados que permitem agrupar diferentes tipos de variáveis sob um único nome.
    type Pessoa struct {
        Nome  string
        Idade int
    }
    pessoa1 := Pessoa{Nome: "Alice", Idade: 30}
    
    // ponteiros: Variáveis que armazenam o endereço de memória de outra variável.
    var x int = 10
    var ponteiroParaX *int = &x
    
    // função: Em Go, as funções são tipos de primeira classe, o que significa que podem ser tratadas como valores.
    resultadoSoma := soma(3, 4)
    
    // Imprimir os valores das variáveis
    fmt.Println("Idade:", idade)
    fmt.Println("Altura:", altura)
    fmt.Println("Nome:", nome)
    fmt.Println("Ativo:", ativo)
    fmt.Println("Array de números:", numeros)
    fmt.Println("Slice de números:", numerosSlice)
    fmt.Println("Mapa de idade por nome:", idadePorNome)
    fmt.Println("Pessoa 1:", pessoa1)
    fmt.Println("Valor de x:", x)
    fmt.Println("Ponteiro para x:", ponteiroParaX)
    fmt.Println("Resultado da soma:", resultadoSoma)
}

// Função de soma
func soma(a, b int) int {
    return a + b
}

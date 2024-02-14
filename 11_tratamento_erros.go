package main
 // Em Go, o tratamento de erros é uma prática fundamental para lidar com situações inesperadas ou excepcionais que podem ocorrer durante a execução de um programa. 
 // O idioma oferece um mecanismo simples e eficaz para lidar com erros usando valores de erro explícitos. 
import (
    "errors"
    "fmt"
)

// Função que retorna um erro caso o número seja negativo
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("não é possível dividir por zero")
    }
    if a < 0 || b < 0 {
        return 0, errors.New("números negativos não são suportados")
    }
    return a / b, nil
}

func main() {
    // Exemplo de uso da função divide
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println("Resultado da divisão:", result)
    }

    // Exemplo de tratamento de erro
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println("Resultado da divisão:", result)
    }
}

// Neste exemplo, a função divide recebe dois números inteiros e retorna o resultado da divisão e um erro, caso ocorra algum problema. 
// Na função main, chamamos a função divide duas vezes - uma vez com argumentos válidos e outra vez com um divisor zero para demonstrar o tratamento de erros.
// Quando um erro ocorre, a função divide retorna 0 como resultado da divisão e um valor de erro. Na função main, verificamos se o erro é nulo (nil). 
// Se não for nulo, isso indica que um erro ocorreu e tratamos o erro de acordo, imprimindo uma mensagem de erro.
package main
 // Manipulacao de dados: ponteiros.
import "fmt"

func main() {
    // Declarando uma variável e atribuindo um valor
    x := 10

    // Declarando um ponteiro para a variável x
    var ponteiro *int
    ponteiro = &x // atribui o endereço de memória de x ao ponteiro

    // Alterando o valor de x indiretamente através do ponteiro
    *ponteiro = 20

    // Imprimindo o novo valor de x
    fmt.Println("O novo valor de x é:", x) // Saída: O novo valor de x é: 20
}

  // Neste exemplo, ponteiro armazena o endereço de memória da variável x. 
  // Através do operador de desreferência *, podemos acessar o valor armazenado no endereço de memória apontado por ponteiro.
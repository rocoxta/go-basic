package main
	// Estruturas de controle: if/else, switch, for e goto.
import "fmt"

func main() {
    // if/else: A estrutura condicional if é usada para executar um bloco de código se uma condição especificada for verdadeira. A cláusula else é opcional e é executada se a condição do if for falsa.
    x := 10
    if x > 5 {
        fmt.Println("x é maior que 5")
    } else {
        fmt.Println("x é menor ou igual a 5")
    }

    fmt.Println("")

    // switch: O switch permite a execução condicional de diferentes blocos de código com base no valor de uma expressão.
    dia := "quarta"
    switch dia {
    case "segunda", "terça", "quarta", "quinta", "sexta":
        fmt.Println("Dia de trabalho")
    case "sábado", "domingo":
        fmt.Println("Fim de semana")
    default:
        fmt.Println("Dia inválido")
    }

    fmt.Println("")

    // for: O for é a principal estrutura de repetição em Go. Ele pode ser usado de várias maneiras, incluindo o loop infinito, o loop com condição e o loop como um iterador de coleção.
    // Loop com condição
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    fmt.Println("")

    // Loop como iterador de coleção
    lista := []string{"a", "b", "c"}
    for index, value := range lista {
        fmt.Printf("Índice: %d, Valor: %s\n", index, value)
    }

    fmt.Println("")

    // goto: O goto é uma declaração de salto incondicional que transfere o controle do programa para um rótulo especificado.
    i := 0
Here:
    fmt.Println(i)
    i++
    if i < 3 {
        goto Here
    }
}

	// É importante usar o goto com cautela, pois pode tornar o código menos legível e mais difícil de manter. 
	// O uso excessivo de goto pode levar a código confuso e propenso a erros. Por isso, é recomendável evitá-lo sempre que possível e preferir as estruturas de controle de fluxo mais comuns, como if, else, switch e for.
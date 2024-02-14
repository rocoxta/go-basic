package main
	// Operações de entrada e saída: 
	// Em Go, as operações de entrada e saída básicas são realizadas principalmente 
	// utilizando as funcionalidades fornecidas pelo pacote fmt para entrada e saída formatadas e pelo pacote os para operações de entrada e saída de arquivos. 
import (
    "fmt"
    "os"
)

func main() {
    // Saída básica:
	// Neste exemplo, fmt.Println() é usado para imprimir "Olá, mundo!" na saída padrão.
    fmt.Println("Olá, mundo!")

    // Ler da entrada padrão:
	// Neste exemplo, fmt.Scanln() é usado para ler uma linha da entrada padrão (normalmente o teclado) e armazená-la na variável nome.
    var nome string
    fmt.Print("Digite seu nome: ")
    fmt.Scanln(&nome)
    fmt.Println("Olá,", nome)

    // Operações de arquivo:
	// Neste exemplo, um arquivo chamado exemplo.txt é criado e o texto é escrito nele usando a função arquivo.Write(). 
	// A função os.Create() é usada para criar o arquivo.
    texto := []byte("Exemplo de texto para escrever em um arquivo.")
    
    arquivo, err := os.Create("exemplo.txt")
    if err != nil {
        fmt.Println("Erro ao criar arquivo:", err)
        return
    }
    defer arquivo.Close()
    
    _, err = arquivo.Write(texto)
    if err != nil {
        fmt.Println("Erro ao escrever no arquivo:", err)
        return
    }
    
    fmt.Println("Texto escrito com sucesso no arquivo.")
}

// Em Go (Golang), a alocação de memória é gerenciada automaticamente pelo coletor de lixo (garbage collector), 
// permitindo que os programadores se concentrem na lógica do programa sem se preocupar explicitamente com a alocação ou desalocação de memória. 
// No entanto, é útil entender como funciona esse processo.

// 1- Alocação de memória estática: Variáveis ​​que são declaradas em nível de pacote ou como variáveis ​​locais são alocadas na memória estática. 
// Essas variáveis ​​são alocadas quando o programa começa a ser executado e permanecem até o final da execução do programa.

	package main

	import "fmt"

	var globalVariable int // Alocada na memória estática

	func main() {
		localVariable := 10 // Alocada na memória estática
		fmt.Println(globalVariable, localVariable)
	}

// 2- Alocação de memória dinâmica: Em Go, alocamos memória dinamicamente usando a palavra-chave new ou usando a 
// sintaxe de inicialização de compostos (slices, maps, channels) que alocam e inicializam dinamicamente a memória necessária.

	package main

	import "fmt"

	func main() {
		// Alocando um novo inteiro na memória dinâmica
		ptr := new(int)
		*ptr = 10
		fmt.Println(*ptr)
	}

// Exemplo usando slices:

	package main

	import "fmt"

	func main() {
		// Alocando um slice de inteiros na memória dinâmica
		slice := make([]int, 5)
		slice[0] = 1
		fmt.Println(slice)
	}

// Exemplo usando maps:

	package main

	import "fmt"

	func main() {
		// Alocando um map de string para int na memória dinâmica
		m := make(map[string]int)
		m["one"] = 1
		fmt.Println(m)
	}

// Exemplo usando channels:

	package main

	import "fmt"

	func main() {
		// Alocando um channel de inteiros na memória dinâmica
		ch := make(chan int)
		go func() {
			ch <- 42
		}()
		fmt.Println(<-ch)
	}

// 3- Coleta de lixo: O coletor de lixo é responsável por desalocar a memória que não é mais referenciada pelo programa. 
// Isso permite que o programador não precise se preocupar com a liberação de memória manualmente.

	package main

	import "time"

	func main() {
		for {
			// Aloca memória dinamicamente em cada iteração
			_ = make([]byte, 1024)
			time.Sleep(time.Second)
		}
	}

// Neste exemplo, o coletor de lixo do Go irá detectar que os slices alocados em cada iteração não são mais acessíveis e irá desalocá-los automaticamente quando necessário.
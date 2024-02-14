// Em Go (ou Golang), concorrência é a capacidade de executar várias tarefas simultaneamente. 
// Isso é conseguido principalmente por meio de duas primitivas principais: goroutines e canais (goroutines and channels).

// 1- Goroutines: Goroutines são como "threads" leves que Go gerencia para você. 
// Elas permitem que você execute funções de forma assíncrona, sem bloquear o fluxo principal do programa. Por exemplo:

	package main

	import (
		"fmt"
		"time"
	)

	func minhaFuncao() {
		for i := 0; i < 5; i++ {
			fmt.Println("Minha Goroutine")
			time.Sleep(1 * time.Second)
		}
	}

	func main() {
		go minhaFuncao() // Iniciando uma goroutine
		time.Sleep(5 * time.Second)
		fmt.Println("Fim do programa")
	}

// Neste exemplo, minhaFuncao() é executada como uma goroutine, permitindo que o programa principal continue sua execução enquanto a função está sendo executada em segundo plano.

// 2- Canais (channels): Canais são uma forma de comunicação entre goroutines. Eles permitem que as goroutines enviem e recebam valores uns dos outros, facilitando a sincronização e a troca de dados. Por exemplo:

	package main

	import (
		"fmt"
	)

	func produtor(c chan int) {
		for i := 0; i < 5; i++ {
			c <- i // Enviando valores para o canal
		}
		close(c)
	}

	func consumidor(c chan int) {
		for num := range c {
			fmt.Println("Recebido:", num) // Recebendo valores do canal
		}
	}

	func main() {
		canal := make(chan int) // Criando um canal

		go produtor(canal) // Iniciando uma goroutine para o produtor
		consumidor(canal)
	}

//	Neste exemplo, a goroutine do produtor envia números para o canal, e a goroutine do consumidor recebe esses números e os imprime.
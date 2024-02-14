// Em Go (Golang), a concorrência é a capacidade de executar várias partes de um programa simultaneamente, aproveitando ao máximo os recursos disponíveis, como CPUs múltiplas ou núcleos. 
// Isso é alcançado principalmente usando goroutines e canais para coordenar a execução concorrente.
// Vamos dar uma olhada em dois exemplos simples de como usar wait groups e mutexes para sincronizar goroutines:

// 1- Utilização de Wait Groups:
// Um sync.WaitGroup é uma estrutura de sincronização que permite que uma ou mais goroutines esperem até que um grupo de goroutines tenha concluído a execução antes de continuar. Aqui está um exemplo básico:

	package main

	import (
		"fmt"
		"sync"
		"time"
	)

	func worker(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}

	func main() {
		var wg sync.WaitGroup

		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go worker(i, &wg)
		}

		wg.Wait()
		fmt.Println("All workers have finished")
	}

// Neste exemplo, a função worker simula algum trabalho e espera um segundo antes de terminar. No main, cinco goroutines são iniciadas, cada uma representando um trabalhador. 
// O programa espera até que todas as goroutines tenham terminado antes de imprimir "All workers have finished".

// 2- Utilização de Mutexes:
// Um mutex (mutual exclusion) é usado para controlar o acesso a recursos compartilhados por múltiplas goroutines. 
// Isso evita que várias goroutines acessem o recurso simultaneamente, o que pode levar a condições de corrida e resultados indeterminados. Aqui está um exemplo:

	package main

	import (
		"fmt"
		"sync"
	)

	var counter = 0
	var mutex = &sync.Mutex{}

	func increment(wg *sync.WaitGroup) {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		counter++
	}

	func main() {
		var wg sync.WaitGroup

		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go increment(&wg)
		}

		wg.Wait()
		fmt.Println("Counter:", counter)
	}

// Neste exemplo, múltiplas goroutines estão incrementando uma variável counter compartilhada. 
// O acesso a essa variável é protegido por um mutex para garantir que apenas uma goroutine possa modificá-la de cada vez. 
// O programa espera até que todas as goroutines tenham terminado antes de imprimir o valor final de counter.
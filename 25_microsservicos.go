// Exemplo básico de como você poderia implementar microsserviços em Go: Servidor principal e 2 endpoints.
// main.go - Servidor principal

	package main

	import (
		"fmt"
		"log"
		"net/http"
		"github.com/gorilla/mux"
	)

	func main() {
		router := mux.NewRouter()
		router.HandleFunc("/api/servico1", Servico1Handler)
		router.HandleFunc("/api/servico2", Servico2Handler)

		fmt.Println("Servidor rodando na porta 8080...")
		log.Fatal(http.ListenAndServe(":8080", router))
	}

	func Servico1Handler(w http.ResponseWriter, r *http.Request) {
		// Implemente o código do serviço 1 aqui
		fmt.Fprintf(w, "Serviço 1")
	}

	func Servico2Handler(w http.ResponseWriter, r *http.Request) {
		// Implemente o código do serviço 2 aqui
		fmt.Fprintf(w, "Serviço 2")
	}

// Este é o arquivo principal main.go que configura o servidor HTTP usando o pacote net/http e o roteador gorilla/mux. 
// Este servidor possui dois endpoints, /api/servico1 e /api/servico2, que podem ser considerados como microsserviços.
// Aqui está um exemplo de implementação de um microsserviço específico em Go:

// servico1.go - Serviço 1

	package main

	import (
		"fmt"
		"log"
		"net/http"
	)

	func main() {
		http.HandleFunc("/api/servico1", Servico1Handler)

		fmt.Println("Servidor Servico1 rodando na porta 8081...")
		log.Fatal(http.ListenAndServe(":8081", nil))
	}

	func Servico1Handler(w http.ResponseWriter, r *http.Request) {
		// Implemente o código do serviço 1 aqui
		fmt.Fprintf(w, "Serviço 1")
	}

// Este é um exemplo de microsserviço separado, servico1.go, que fornece o mesmo endpoint /api/servico1 que o servidor principal. 
// Isso poderia ser uma implementação separada e independente que serve a um propósito específico.
// Para executar esses exemplos, você pode compilá-los e executá-los separadamente. Por exemplo, você pode executar go run main.go para o servidor principal e go run servico1.go para o serviço 1. 
// Certifique-se de que não haja conflito de portas ao executar microsserviços simultaneamente.
package main
// Exemplo de API RESTful em Go (Golang) usando o pacote padrão net/http:
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// Struct para representar um item
type Item struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}

// Slice para armazenar os itens
var inventory []Item

// Função para lidar com requisições GET em /items
func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(inventory)
}

// Função para lidar com requisições POST em /items
func addItem(w http.ResponseWriter, r *http.Request) {
    var newItem Item
    json.NewDecoder(r.Body).Decode(&newItem)
    inventory = append(inventory, newItem)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newItem)
}

func main() {
    // Adiciona alguns itens à lista de inventário para fins de demonstração
    inventory = append(inventory, Item{ID: "1", Name: "Item 1", Price: 10})
    inventory = append(inventory, Item{ID: "2", Name: "Item 2", Price: 20})

    // Configuração das rotas
    http.HandleFunc("/items", getItems) // Rota para listar itens
    http.HandleFunc("/items/add", addItem) // Rota para adicionar um item

    // Inicia o servidor na porta 8080
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


// Este exemplo cria uma API RESTful simples em Go que tem duas rotas:

//1- items: Para listar todos os itens.
//2- items/add: Para adicionar um novo item.

// Você pode testar essa API usando ferramentas como cURL, Postman ou Insomnia. 
// Por exemplo, para adicionar um novo item, você pode fazer uma solicitação POST para http://localhost:8080/items/add com um corpo JSON contendo os detalhes do item que você deseja adicionar.
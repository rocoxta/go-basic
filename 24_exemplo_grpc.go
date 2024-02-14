// Primeiro, você precisa definir o arquivo .proto, que é um arquivo de IDL (Interface Description Language) 
// usado para definir os serviços e as mensagens que serão trocadas entre os clientes e os servidores.

	syntax = "proto3";

	package example;

	service Greeter {
		rpc SayHello (HelloRequest) returns (HelloResponse);
	}

	message HelloRequest {
		string name = 1;
	}

	message HelloResponse {
		string message = 1;
	}

// Aqui, estamos definindo um serviço chamado Greeter, que possui um método RPC SayHello. Este método leva uma mensagem HelloRequest e retorna uma mensagem HelloResponse.
// Em seguida, você precisa compilar o arquivo .proto para gerar o código Go correspondente. Para fazer isso, você precisa instalar o compilador protobuf para Go:

	go get -u github.com/golang/protobuf/protoc-gen-go

// Depois de instalar o compilador, você pode gerar o código Go a partir do arquivo .proto usando o seguinte comando:

	protoc --go_out=. example.proto

// Agora você pode implementar os servidores e clientes em Go usando o código gerado.
// Aqui está um exemplo simples de como implementar um servidor e um cliente:
// server.go:

	package main

	import (
		"context"
		"fmt"
		"log"
		"net"

		"example"
		"google.golang.org/grpc"
	)

	type server struct{}

	func (s *server) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
		return &example.HelloResponse{
			Message: fmt.Sprintf("Hello, %s!", req.GetName()),
		}, nil
	}

	func main() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		example.RegisterGreeterServer(s, &server{})
		log.Println("Server started...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}

// client.go:

	package main

	import (
		"context"
		"log"
		"os"
		"time"

		"example"
		"google.golang.org/grpc"
	)

	func main() {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := example.NewGreeterClient(conn)

		name := "Alice"
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &example.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}

// Você pode executar o servidor usando go run server.go e o cliente usando go run client.go. Certifique-se de executar o servidor antes do cliente.
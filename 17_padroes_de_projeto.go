// Na linguagem Go (ou Golang), alguns padrões de projeto comuns são: Singleton, Builder, Factory e Strategy.

// 1- Singleton (Singleton): O Singleton é um padrão de design que restringe a instância de uma classe a um único objeto. 
// Em Go, pode ser implementado usando a inicialização preguiçosa e uma variável global.

	package singleton

	import "sync"

	type Singleton struct {
		// Campos
	}

	var instance *Singleton
	var once sync.Once

	func GetInstance() *Singleton {
		once.Do(func() {
			instance = &Singleton{}
		})
		return instance
	}

// 2- Builder (Builder): O padrão Builder é usado para construir objetos complexos passo a passo. 
// Em Go, isso pode ser implementado usando uma estrutura que expõe métodos para configurar o objeto.

	package builder

	type Product struct {
		// Campos
	}

	type ProductBuilder struct {
		product Product
	}

	func NewProductBuilder() *ProductBuilder {
		return &ProductBuilder{}
	}

	func (pb *ProductBuilder) SetField1(value int) *ProductBuilder {
		pb.product.Field1 = value
		return pb
	}

	func (pb *ProductBuilder) SetField2(value string) *ProductBuilder {
		pb.product.Field2 = value
		return pb
	}

	func (pb *ProductBuilder) Build() *Product {
		return &pb.product
	}

// Factory (Factory): O padrão Factory é usado para criar objetos sem especificar a classe exata do objeto que será criado. 
// Em Go, isso pode ser implementado usando funções que retornam interfaces.

	package factory

	type Product interface {
		Method1() string
	}

	type ConcreteProduct1 struct{}

	func (p *ConcreteProduct1) Method1() string {
		return "Product 1 method"
	}

	type ConcreteProduct2 struct{}

	func (p *ConcreteProduct2) Method1() string {
		return "Product 2 method"
	}

	func CreateProduct(productType string) Product {
		switch productType {
		case "Product1":
			return &ConcreteProduct1{}
		case "Product2":
			return &ConcreteProduct2{}
		default:
			return nil
		}
	}

// Strategy (Strategy): O padrão Strategy define uma família de algoritmos, encapsula cada um deles e os torna intercambiáveis. 
// Em Go, isso pode ser implementado usando interfaces.

	package strategy

	type Strategy interface {
		ExecuteAlgorithm() string
	}

	type ConcreteStrategy1 struct{}

	func (s *ConcreteStrategy1) ExecuteAlgorithm() string {
		return "Strategy 1"
	}

	type ConcreteStrategy2 struct{}

	func (s *ConcreteStrategy2) ExecuteAlgorithm() string {
		return "Strategy 2"
	}

	type Context struct {
		strategy Strategy
	}

	func (c *Context) SetStrategy(strategy Strategy) {
		c.strategy = strategy
	}

	func (c *Context) ExecuteAlgorithm() string {
		return c.strategy.ExecuteAlgorithm()
	}

// Esses são apenas alguns exemplos dos padrões de design comuns em Go. 
// A linguagem Go oferece uma sintaxe simples e direta que torna a implementação desses padrões relativamente simples e eficiente.
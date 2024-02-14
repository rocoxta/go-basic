// Gin, Echo, Beego, Gorilla e Revel.

// Os frameworks web em Go (Golang) são conjuntos de bibliotecas e ferramentas que facilitam o desenvolvimento de aplicativos da web em Go, fornecendo estruturas e 
// funcionalidades comuns para lidar com rotas, middleware, manipulação de solicitações HTTP, autenticação, renderização de modelos, entre outras tarefas relacionadas ao desenvolvimento web.
// Aqui estão exemplos de alguns frameworks web populares em Go:

// 1-Gin: Gin é um framework web minimalista e de alto desempenho para Go. 
// Ele fornece uma sintaxe simples e rápida para definir rotas, além de suportar middleware, renderização de templates e integração com bibliotecas de autenticação e validação de formulários.
	
	package main

	import "github.com/gin-gonic/gin"

	func main() {
		router := gin.Default()

		router.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, Gin!",
			})
		})

		router.Run(":8080")
	}

// 2- Echo: Echo é outro framework web leve e de alto desempenho para Go. 
// Ele é conhecido por sua simplicidade e velocidade, oferecendo uma API intuitiva para definir rotas, middleware e manipulação de solicitações HTTP. Echo é muito popular para criar APIs RESTful.

	package main

	import "github.com/labstack/echo/v4"

	func main() {
		e := echo.New()

		e.GET("/hello", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"message": "Hello, Echo!",
			})
		})

		e.Start(":8080")
	}

// 3- Beego: Beego é um framework web completo e de alto nível para Go. 
// Ele segue o padrão MVC (Model-View-Controller) e oferece uma ampla gama de recursos integrados, como ORM (Object-Relational Mapping), validação de formulários, sessões, cache e internacionalização.

	package main

	import "github.com/astaxie/beego"

	type MainController struct {
		beego.Controller
	}

	func (c *MainController) Get() {
		c.Data["json"] = map[string]string{"message": "Hello, Beego!"}
		c.ServeJSON()
	}

	func main() {
		beego.Router("/hello", &MainController{})
		beego.Run(":8080")
	}

// 4- Gorilla: Gorilla não é exatamente um framework completo, mas uma coleção de pacotes úteis para desenvolvimento web em Go. 
// Ele inclui pacotes para manipulação de roteamento, middleware, sessões, autenticação, WebSocket e outros.

package main

	import (
		"net/http"

		"github.com/gorilla/mux"
	)

	func helloHandler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello, Gorilla!"}`))
	}

	func main() {
		r := mux.NewRouter()
		r.HandleFunc("/hello", helloHandler).Methods("GET")

		http.ListenAndServe(":8080", r)
	}

// 5- Revel: Revel é um framework web de opinião completa para Go, inspirado em frameworks web populares de outras linguagens, como Ruby on Rails. 
// Ele é projetado para facilitar o desenvolvimento rápido de aplicativos web escaláveis, seguindo as melhores práticas de desenvolvimento.

	package controllers

	import "github.com/revel/revel"

	type App struct {
		*revel.Controller
	}

	func (c App) Index() revel.Result {
		return c.RenderJSON(map[string]string{"message": "Hello, Revel!"})
	}

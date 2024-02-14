// Exemplo básico de como prevenir ataques comuns, como injeção de SQL e XSS (Cross-Site Scripting) em um aplicativo Go:
// 1- Prevenção de Injeção de SQL: Para prevenir a injeção de SQL, é altamente recomendável usar consultas parametrizadas ou preparadas ao invés de concatenar diretamente strings SQL.

	package main

	import (
		"database/sql"
		"fmt"
		"log"
	)

	func main() {
		// Suponha que db seja seu objeto de conexão com o banco de dados
		db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Suponha que userName seja uma entrada do usuário
		userName := "usuario'; DROP TABLE usuarios; --"
		
		// Consulta parametrizada
		rows, err := db.Query("SELECT * FROM usuarios WHERE nome = ?", userName)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		
		// Iterar sobre os resultados
		for rows.Next() {
			var id int
			var nome string
			if err := rows.Scan(&id, &nome); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("ID: %d, Nome: %s\n", id, nome)
		}
	}

// Neste exemplo, usamos uma consulta parametrizada SELECT * FROM usuarios WHERE nome = ? em vez de concatenar a string diretamente com a entrada do usuário.
// Prevenção de XSS: Para prevenir ataques XSS, é importante escapar as entradas do usuário antes de exibi-las em páginas da web. 

	package main

	import (
		"fmt"
		"html/template"
		"net/http"
	)

	func main() {
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			// Suponha que userName seja uma entrada do usuário
			userName := r.FormValue("name")
			
			// Escapar a entrada do usuário antes de exibi-la
			safeUserName := template.HTMLEscapeString(userName)
			
			fmt.Fprintf(w, "Olá, %s!", safeUserName)
		})

		http.ListenAndServe(":8080", nil)
	}


// Neste exemplo, usamos template.HTMLEscapeString para escapar a entrada do usuário antes de exibi-la na página da web. 
// Isso evitará que o navegador interprete o conteúdo como HTML ou JavaScript malicioso.
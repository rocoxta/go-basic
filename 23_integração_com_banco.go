// Exemplo simples de como você pode se conectar a um banco de dados MySQL usando a linguagem Go (Golang). 
// Certifique-se de que você tenha o driver MySQL instalado. Você pode instalá-lo usando o seguinte comando:

package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // DSN (Data Source Name) que contém as informações de conexão
    dsn := "username:password@tcp(127.0.0.1:3306)/database_name"

    // Conectando ao banco de dados MySQL
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Verificando se a conexão com o banco de dados está disponível
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Conexão com o banco de dados estabelecida!")

    // Executando uma consulta simples
    rows, err := db.Query("SELECT * FROM tabela")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    // Iterando sobre os resultados
    for rows.Next() {
        var id int
        var nome string
        if err := rows.Scan(&id, &nome); err != nil {
            panic(err.Error())
        }
        fmt.Printf("ID: %d, Nome: %s\n", id, nome)
    }

    // Verificando se houve algum erro durante o processo
    if err := rows.Err(); err != nil {
        panic(err.Error())
    }
}

// Lembre-se de substituir "username", "password", "127.0.0.1:3306" e "database_name" pelas suas próprias credenciais de acesso ao banco de dados MySQL. 
// Em Go, a utilização de módulos e dependências é uma prática fundamental para gerenciar as bibliotecas e pacotes que um projeto utiliza. 
// Os módulos são uma maneira de organizar e versionar o código, permitindo que diferentes projetos possam compartilhar código de forma eficiente e segura.
// Para começar a usar módulos em um projeto Go, você precisa iniciar um novo módulo. Isso é feito executando o comando go mod init no diretório raiz do seu projeto. 
// Por exemplo:

$ go mod init github.com/seu-usuario/seu-projeto

// Isso criará um arquivo go.mod no diretório raiz do seu projeto. 
// Esse arquivo é onde as informações sobre o módulo, suas dependências e as versões específicas das dependências são registradas.

// Para adicionar uma dependência a um projeto Go, você pode usar o comando go get. Por exemplo:

$ go get github.com/pacote/exemplo

// Isso adicionará o pacote github.com/pacote/exemplo como uma dependência ao seu projeto e atualizará o arquivo go.mod para incluir essa dependência.
// Aqui está um exemplo de um go.mod básico após adicionar algumas dependências:

module github.com/seu-usuario/seu-projeto

go 1.17

require (
    github.com/pacote/exemplo v1.0.0
    github.com/outra-dependencia/pacote v2.3.1
)

// Este arquivo go.mod especifica o módulo atual (github.com/seu-usuario/seu-projeto), a versão mínima do Go necessária (go 1.17), e as dependências requeridas, junto com suas versões específicas.
// Em Go (Golang), os perfis de CPU e memória são ferramentas de diagnóstico que ajudam os desenvolvedores a entender e otimizar o desempenho de seus programas. 
// Esses perfis são particularmente úteis para identificar gargalos de desempenho, vazamentos de memória e áreas de código que podem ser otimizadas.

// 1- Perfil de CPU: Este perfil fornece informações sobre onde a CPU está gastando seu tempo durante a execução do programa. 
// Ele mostra a quantidade de tempo que cada função ou linha de código leva para ser executada, permitindo que os desenvolvedores 
// identifiquem partes do código que podem estar consumindo muitos recursos de CPU. Um exemplo de como gerar um perfil de CPU em Go é usando a ferramenta de linha de comando go tool pprof. Por exemplo:

	go test -cpuprofile=cpu.prof

// Esse comando irá gerar um perfil de CPU para um teste em Go. Depois de executar o teste, você pode usar o comando go tool pprof para analisar e visualizar o perfil:
	
	go tool pprof -http=:8080 cpu.prof

// Isso abrirá uma interface web onde você pode analisar o perfil de CPU.

// 2- Perfil de Memória: Este perfil fornece informações sobre o uso de memória do programa durante a execução. 
// Ele pode ajudar a identificar vazamentos de memória ou áreas do código que estão consumindo uma quantidade excessiva de memória. 
// Assim como o perfil de CPU, o perfil de memória é gerado usando a ferramenta go tool pprof. Por exemplo:

	go test -memprofile=mem.prof

// Isso irá gerar um perfil de memória para o teste em Go. Da mesma forma que com o perfil de CPU, você pode usar o go tool pprof para analisar e visualizar o perfil:
	
	go tool pprof -http=:8080 mem.prof

// Isso abrirá uma interface web onde você pode analisar o perfil de memória.
package main
	// Manipulacao de dados: arrays, slices, maps.
import "fmt"

func main() {
    // Arrays: Arrays em Go são estruturas de dados fixas que armazenam uma coleção de elementos do mesmo tipo.
    var array [5]int
    array[0] = 10
    array[1] = 20
    array[2] = 30
    array[3] = 40
    array[4] = 50

    fmt.Println("Elemento 3:", array[2])
    array[2] = 35
    fmt.Println("Elemento 3 modificado:", array[2])

    // Slices: Slices em Go são estruturas de dados dinâmicas que representam uma parte de um array.
    slice := []int{10, 20, 30, 40, 50}

    fmt.Println("Elemento 3:", slice[2])
    slice[2] = 35
    fmt.Println("Elemento 3 modificado:", slice[2])

    slice = append(slice, 60)
    fmt.Println("Slice após adicionar 60:", slice)

    // Maps: Maps em Go são estruturas de dados que armazenam pares chave-valor, onde cada chave é única.
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    fmt.Println("Valor de 'b':", m["b"])
    m["b"] = 20
    fmt.Println("Valor de 'b' modificado:", m["b"])

    m["d"] = 4
    fmt.Println("Map após adicionar 'd':", m)
}

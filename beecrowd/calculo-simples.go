package main

import "fmt"

func main() {
    var total float64

    for i := 0; i < 2; i++ {
        var codigo int
        var quantidade, preco float64

        fmt.Scan(&codigo, &quantidade, &preco)

        total += quantidade * preco
    }

    fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}

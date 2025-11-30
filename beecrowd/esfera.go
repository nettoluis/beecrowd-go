package main

import "fmt"

const (
    Pi = 3.14159
)

func calculaVolumeEsfera(raio float64) float64 {
    return (4.0 / 3.0) * Pi * raio * raio * raio
}

func main() {
    var raio float64

    fmt.Scan(&raio)

    volume := calculaVolumeEsfera(raio)

    fmt.Printf("VOLUME = %.3f\n", volume)
}

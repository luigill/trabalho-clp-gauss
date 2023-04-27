package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}


func gaussElimination(){
    var norm, linha, col int

    var multiplier float64

    for norm := 0; norm < N - 1; norm++ {
        for linha := norm + 1; linha < N; linha++ {
            multiplier = A[linha][norm] / A[norm][norm]
            for col := norm; col < N; col++ {
                A[linha][col] 
            }
        }   
    }

}
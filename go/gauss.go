package main

import (
    "fmt"  
    "math/rand"
    "time"
    //"os"
 )

const maxN = 2000
var N int

var A [][]int
var B []int
var X []int


func main() {
    fmt.Println("Hello, world!")
    rand.Seed(time.Now().UnixNano())

    //N = os.Args[1:]

    N = 3

    A = make([][]int, N)
    for i := range A {
        A[i] = make([]int, N)
    }
    B = make([]int, N)
    X = make([]int, N)

    initialize_in()
    print_in()
    gaussElimination()
    print_out()
}


func gaussElimination(){
    var norm, linha, col int

    var multiplier int

    for norm = 0; norm < N - 1; norm++ {
        for linha = norm + 1; linha < N; linha++ {
            multiplier = A[linha][norm] / A[norm][norm]
            for col = norm; col < N; col++ {
                A[linha][col] -= A[norm][col] * multiplier
            }
            B[linha] -= B[norm] * multiplier
        }   
    }

    //Back Substituition
    for linha := N - 1; linha >= 0; linha-- {
        X[linha] = B[linha]
        for col := N-1; col > linha; col-- {
            X[linha] -= A[linha][col] * X[col]
        }    
        X[linha] /= A[linha][linha]
    }
}

//FEITO
func initialize_in(){
    var linha, col int

    fmt.Printf("Initializing...\n")
    for col = 0; col < N; col++ {
        for linha = 0; linha < N; linha++ {
            A[linha][col] = rand.Intn(100)
        }
        B[col] = rand.Intn(100)
        X[col] = 0
    }
}

//FEITO
func print_in(){
    var linha, col int        
        fmt.Println("A")
        for linha = 0; linha < N; linha++ {
            for col = 0; col < N; col++ {
                fmt.Printf("%d ", A[linha][col])
        }
        fmt.Println();
        }

        fmt.Println("B")
        for col = 0; col < N; col++ {
            fmt.Printf("%d ", B[col])
        }
        fmt.Println();
}


func print_out(){

    var linha int

    if N < 100 {
        fmt.Println("X")
        for linha = 0; linha < N; linha++ {
            fmt.Printf("%d ", X[linha])
        }
    }
}
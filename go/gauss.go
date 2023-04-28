package main

import (
    "fmt"  
    "math/rand"
    "time" )

const maxN = 2000
var N int

var A [][]int
var B []int
var X []int


func main() {
    fmt.Println("Hello, world!")
    rand.Seed(time.Now().UnixNano())
    A := make([][]int, maxN)
    for i := range A {
        A[i] = make([]int, maxN)
    }

    initialize_in()
    print_in()
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

func print_in(){
    var linha, col int

    if N < 10 {       
        fmt.Println("A = \n\t")
        for linha = 0; linha < N; linha++ {
            for col = 0; col < N; col++ {
                fmt.Printf("%d%s", A[linha][col], func() string {
                    if col < N-1 {
                        return "; "
                    }
                    return "]\n"
                }())                 
            }
        }
        fmt.Println("B = [")
        for col := 0; col < N; col++ {
            fmt.Printf("%d%s", B[col], func() string {
                if col < N-1 {
                    return "; "
                }
                return "]\n"
            }())
            
        }
    }
}

func print_out(){
    var linha int

    if N < 100 {
        fmt.Println("X = [")
        for linha = 0; linha < N; linha++ {
            fmt.Printf("%5.2d%s", X[linha], func() string {
                if linha < N-1 {
                    return "; "
                }
                return "]\n"
            }())
            
        }
    }
}
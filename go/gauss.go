package main

import (
    "fmt"  
    "math/rand"
    "time"
    "os"
    "strconv"
 )

var N int
var err error

var A [][]float64
var B []float64
var X []float64


func main() {
    start := time.Now()
    rand.Seed(time.Now().UnixNano())

    if len(os.Args) < 2 {
        fmt.Println("Uso: go run gauss.go <int>")
        return
    }

    N, err = strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Erro: Argumento Inválido")
        return
    }


    A = make([][]float64, N)
    for i := range A {
        A[i] = make([]float64, N)
    }
    B = make([]float64, N)
    X = make([]float64, N)

    initialize_in()
    //initialize_in_test()
    print_in()
    gaussElimination()
    print_out()

    elapsed := time.Since(start)
    fmt.Println()
    fmt.Printf("Tempo total de execução: %s\n", elapsed)
}


func gaussElimination(){
    var norm, linha, col float64

    var multiplier float64

    for norm = 0; norm < float64(N) - 1; norm++ {
        for linha = norm + 1; linha < float64(N); linha++ {
            multiplier = A[int(linha)][int(norm)] / A[int(norm)][int(norm)]
            for col = norm; col < float64(N); col++ {
                A[int(linha)][int(col)] -= A[int(norm)][int(col)] * multiplier
            }
            B[int(linha)] -= B[int(norm)] * multiplier
        }   
    }

    //Back Substituition
    for linha := N - 1; linha >= 0; linha-- {
        X[int(linha)] = B[int(linha)]
        for col := N-1; col > int(linha); col-- {
            X[int(linha)] -= A[int(linha)][col] * X[col]
        }    
        X[int(linha)] /= A[int(linha)][int(linha)]
    }
}

func initialize_in(){
    var linha, col int

    fmt.Printf("Initializing...\n")
    for col = 0; col < N; col++ {
        for linha = 0; linha < N; linha++ {
            A[linha][col] = rand.Float64() * 2000.0 - 1000.0
        }
        B[col] = rand.Float64() * 2000.0 - 1000.0
        X[col] = 0
    }
}

//Inicializa os inputs com um caso de teste corrigido da aplicação original
func initialize_in_test(){
	A = [][]float64{
       {55062.54, 59745.72, 18204.24},
       {25846.28, 12946.73, 36304.98},
       {51321.19, 21969.16, 31286.69},
    }

	B = []float64{52326.57, 50346.70, 41213.68}

    //Resultados Esperados
    //-0.374256 0.804421 1.366347 
}

func print_in(){
    var linha, col int        
        fmt.Println("A")
        for linha = 0; linha < N; linha++ {
            for col = 0; col < N; col++ {
                fmt.Printf("%f ", A[linha][col])
        }
        fmt.Println();
        }

        fmt.Println("B")
        for col = 0; col < N; col++ {
            fmt.Printf("%f ", B[col])
        }
        fmt.Println();
}

func print_out(){

    var linha int

    if N < 100 {
        fmt.Println("X")
        for linha = 0; linha < N; linha++ {
            fmt.Printf("%f ", X[linha])
        }
    }
}

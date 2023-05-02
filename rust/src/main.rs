use rand::prelude::*;
use std::{env, usize};
use std::time::Instant;

fn main() {
    let start = Instant::now();

    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        println!("Uso: cargo run <int>");
        return;
    }

    let num:usize = match args[1].parse::<usize>() {
        Ok(n) => n,
        Err(_) => {
            println!("Erro: Argumento Inválido");
            return;
        }
    };

    let mut a = vec![vec![0.0; num]; num];
    let mut b = vec![0.0; num];
    let mut x = vec![0.0; num];

    initialize_in(&mut a, &mut b, &mut x, num);
    //initialize_in_test(&mut a, &mut b);
    print_in(&a, &b);
    gauss_elimination(&mut a, &mut b, num, &mut x);
    print_out(&x, num);

    println!();

    let end = Instant::now();
    let duration = end.duration_since(start);
    println!("Tempo total de execução: {:?}", duration);
}

fn gauss_elimination(a: &mut Vec<Vec<f64>>, b: &mut Vec<f64>, n: usize, x: &mut Vec<f64> ) {
    for norm in 0..(n - 1) {
        for linha in (norm + 1)..n {
            let multiplier = a[linha][norm] / a[norm][norm];
            for col in norm..n {
                a[linha][col] -= a[norm][col] * multiplier;
            }
            b[linha] -= b[norm] * multiplier;
        }
    }

    // Back Substituition
    for linha in (0..n).rev() {
        x[linha] = b[linha];
        for col in (linha + 1)..n {
            x[linha] -= a[linha][col] * x[col];
        }
        x[linha] /= a[linha][linha];
    }
}

fn initialize_in(a: &mut Vec<Vec<f64>>, b: &mut Vec<f64>, x:  &mut Vec<f64>, n: usize) {
    let mut rng = rand::thread_rng();
 
    println!("Initializing...");
    for col in 0..n {
        for linha in 0..n {
            a[linha][col] = rng.gen_range(-1000.0..1000.0);
        }
        b[col] = rng.gen_range(-1000.0..1000.0);
        x[col] = 0.0;
    }
}

//Inicializa os inputs com um caso de teste corrigido da aplicação original
fn initialize_in_test(a: &mut Vec<Vec<f64>>, b: &mut Vec<f64>) {
    *a = vec![
        vec![55062.54, 59745.72, 18204.24],
        vec![25846.28, 12946.73, 36304.98],
        vec![51321.19, 21969.16, 31286.69],
    ];
    *b = vec![52326.57, 50346.70, 41213.68];

    //Resultados Esperados
    //-0.374256 0.804421 1.366347 
}

fn print_in(a: &Vec<Vec<f64>>, b: &Vec<f64>) {
    println!("A");
    for linha in 0..a.len() {
        for col in 0..a[linha].len() {
            print!("{} ", a[linha][col]);
        }
        println!();
    }

    println!("B");
    for col in 0..b.len() {
        print!("{} ", b[col]);
    }
    println!();
}

fn print_out(x: &Vec<f64>, n: usize) {
    if n < 100 {
        println!("X");
        for linha in 0..n {
            print!("{} ", x[linha]);
        }
    }
}

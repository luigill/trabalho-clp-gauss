use rand::prelude::*;
use std::env;

const MAX_N: usize = 2000;

fn main() {
    let mut rng = thread_rng();

    let args: Vec<String> = env::args().collect();
    let n: usize = if args.len() > 1 {
        args[1].parse().unwrap()
    } else {
        3
    };

    let mut a = vec![vec![0.0; n]; n];
    let mut b = vec![0.0; n];
    let mut x = vec![0.0; n];

    initialize_in_test(&mut a, &mut b);

    print_in(&a, &b);

    gauss_elimination(&mut a, &mut b, n, &mut x);

    print_out(&x, n);
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

fn initialize_in_test(a: &mut Vec<Vec<f64>>, b: &mut Vec<f64>) {
    *a = vec![
        vec![55062.54, 59745.72, 18204.24],
        vec![25846.28, 12946.73, 36304.98],
        vec![51321.19, 21969.16, 31286.69],
    ];
    *b = vec![52326.57, 50346.70, 41213.68];
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

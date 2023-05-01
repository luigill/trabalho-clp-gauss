use rand::Rng;

//static N = 3;

fn main() {
    let mut a: [[f64; 3]; 3] = [[0.0; 3]; 3];
    let mut b: Vec<f64> = Vec::new();
    let mut x: Vec<f64> = Vec::new();
    
    let mut rng = rand::thread_rng();

    //Gerando A, B e X
    println!("Initializing...");
    for col in 0..3 {
        for line in 0..3 {
            a[line][col] = rng.gen_range(1000.0..7000.0);
        }
        b[col] = rng.gen_range(1000.0..7000.0);
        x[col] = 0.0;
    }

    //print_in(a, &b);

    println!("The random matrix is:");
    for row in &a {
        println!("{:?}", row);
    }

    gauss_elimination();
    print_out();
}


fn gauss_elimination(){
    print!("Gauss Elimination");
}

fn print_in(a: [[f64; 3]; 3], b: &Vec<f64>){
    // Print the matrix
    println!("A");
    for row in 0..3 {
        for col  in 0..3 {
            print!("{} ", a[row][col]);
        }
    println!("");

    //println!("B");
    //for col in 0..3 {   
    //        print!("{} ", b[col]);
    //println!("");
    //    }
    }
}

fn print_out(){
    
}


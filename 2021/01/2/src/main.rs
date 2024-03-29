use circular_queue::CircularQueue;
use std::error::Error;
use std::fs::File;
use std::io;
use std::io::BufRead;

fn main() -> Result<(), Box<dyn Error>> {
    let mut window = CircularQueue::with_capacity(3);
    let mut increased = 0;
    let mut first = true;
    let mut prev = 0;
    let f = io::BufReader::new(File::open("input.txt")?);

    for l in f.lines() {
        let l = l?;
        let n = l.parse::<i32>()?;
        window.push(n);
        if window.len() < 3 {
            continue;
        }
        let sum = window.iter().sum();
        if first {
            prev = sum;
            first = false;
            continue;
        }
        if sum > prev {
            increased += 1;
        }
        prev = sum;
    }

    println!("{}", increased);

    Ok(())
}

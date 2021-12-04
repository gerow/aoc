use std::error::Error;
use std::io;
use std::io::BufRead;

fn main() -> Result<(), Box<dyn Error>> {
    let mut increased = 0;
    let mut first = true;
    let mut prev = 0;

    for l in io::stdin().lock().lines() {
        let l = l?;
        let n = l.parse::<i32>()?;
        if first {
            prev = n;
            first = false;
            continue;
        }

        if n > prev {
            increased += 1;
        }
        prev = n;
    }

    println!("{}", increased);

    Ok(())
}

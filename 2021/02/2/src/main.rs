use std::error::Error;
use std::fmt;
use std::fs::File;
use std::io;
use std::io::BufRead;
use std::str::FromStr;
use text_io::scan;

struct Vec2(i32, i32);

impl Vec2 {
    fn scale(&self, s: i32) -> Vec2 {
        Vec2(self.0 * s, self.1 * s)
    }

    fn add(&self, other: &Vec2) -> Vec2 {
        Vec2(self.0 + other.0, self.1 + other.1)
    }
}

fn main() -> Result<(), Box<dyn Error>> {
    let f = io::BufReader::new(File::open("input.txt")?);

    let mut pos = Vec2(0, 0);
    let mut aim = Vec2(1, 0);
    for l in f.lines() {
        let l = l?;
        let action: String;
        let arg: i32;
        scan!(l.bytes() => "{} {}", action, arg);

        match action.as_ref() {
            "down" => aim = Vec2(aim.0, aim.1 - arg),
            "up" => aim = Vec2(aim.0, aim.1 + arg),
            "forward" => pos = pos.add(&aim.scale(arg)),
            _ => panic!("bad action `{}`", action),
        }
    }
    println!("{}", pos.0 * -pos.1);

    Ok(())
}

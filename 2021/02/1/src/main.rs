use std::error::Error;
use std::fmt;
use std::fs::File;
use std::io;
use std::io::BufRead;
use std::str::FromStr;
use text_io::scan;

struct Vec2(i32, i32);

impl Vec2 {
    const FORWARD: Vec2 = Vec2(1, 0);
    const DOWN: Vec2 = Vec2(0, -1);
    const UP: Vec2 = Vec2(0, 1);

    fn scale(&self, s: i32) -> Vec2 {
        Vec2(self.0 * s, self.1 * s)
    }

    fn add(&self, other: &Vec2) -> Vec2 {
        Vec2(self.0 + other.0, self.1 + other.1)
    }
}

#[derive(Debug)]
struct ParseVec2Error;

impl Error for ParseVec2Error {}

impl fmt::Display for ParseVec2Error {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "didn't work, yo")
    }
}

impl FromStr for Vec2 {
    type Err = ParseVec2Error;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let dir_s: String;
        let scale: i32;
        scan!(s.bytes() => "{} {}", dir_s, scale);
        let dir = match dir_s.as_ref() {
            "forward" => Vec2::FORWARD,
            "down" => Vec2::DOWN,
            "up" => Vec2::UP,
            _ => return Err(ParseVec2Error),
        };
        Ok(dir.scale(scale))
    }
}

fn main() -> Result<(), Box<dyn Error>> {
    let f = io::BufReader::new(File::open("input.txt")?);

    let mut pos = Vec2(0, 0);
    for l in f.lines() {
        let l = l?;
        let mov = Vec2::from_str(&l)?;
        pos = pos.add(&mov);
    }
    println!("{}", pos.0 * -pos.1);

    Ok(())
}

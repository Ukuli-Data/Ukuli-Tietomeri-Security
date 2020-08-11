use std::fs::File;
use std::io::prelude::*;
use std::path::Path;
use std::env;

fn main() -> bool {
    let jail: bool = true;
    let args: Vec<String> = env::args().collect();
    let mut hacker_ip = &args[1];
    let linechange = "\n";
    hacker_ip.push_str(&linechange);

    let path = Path::new("hackers.txt");
    let display = path.display();

    let mut file = match File::create(&path) {
        Err(why) => panic!("couldn't create {}: {}", display, why),
        Ok(file) => file,
    };

    match file.write_all(hacker_ip.as_bytes()) {
        Err(why) => panic!("couldn't write to {}: {}", display, why),
        Ok(_) => println!("successfully wrote to {}", display),
    }

    return(jail);
}

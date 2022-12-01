use std::env;
use std::fs;

pub fn get_file_contents() -> String {
    let filename = get_filename();
    fs::read_to_string(filename).expect("Error reading input")
}

fn get_filename() -> String {
    let mut args: Vec<String> = env::args().collect();
    assert!(args.len() > 1, "Must provide a filename!");
    args.remove(1)
}

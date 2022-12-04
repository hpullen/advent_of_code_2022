mod util;

fn main() {
    let data = util::get_file_contents_vec();
    part_one(&data);
}

fn part_one(data: &Vec<String>) {
    let mut total = 0;
    for line in data {
        let compartment_size = line.len() / 2;
        let c1 = &line[..compartment_size];
        let c2 = &line[compartment_size..];
        for c in c1.chars() {
            if c2.contains(c) {
                total += get_priority(c);
                break;
            }
        }
    }
    println!("Part 1: {}", total);
}

fn get_priority(c: char) -> u32 {
    if c.is_uppercase() {
        let a_digit = 'A'.to_digit(10).unwrap();
        let diff = char_digit - a_digit;
        return diff + 27;
    } else {
        let a_digit = 'a'.to_digit(10).unwrap();
        let diff = char_digit - a_digit;
        return diff + 1;
    }
}

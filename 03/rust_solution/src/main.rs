mod util;

fn main() {
    let data = util::get_file_contents_vec();
    part_one(&data);
    part_two(&data);
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

fn part_two(data: &Vec<String>) {
    let mut total = 0;
    let group_size = 3;
    let n = data.len() / group_size;
    for i in 0..n {
        let first_elf = &data[i * group_size];
        for c in first_elf.chars() {
            let mut is_answer = true;
            for j in 1..group_size {
                let elf = &data[i * group_size + j];
                is_answer = is_answer && elf.contains(c);
            }
            if is_answer {
                total += get_priority(c);
                break;
            }
        }
    }
    println!("Part 2: {}", total);
}

fn get_priority(c: char) -> u32 {
    let c_bytes = c.to_string().as_bytes()[0];
    if c.is_uppercase() {
        let a_bytes = 'A'.to_string().as_bytes()[0];
        let diff = (c_bytes - a_bytes) as u32;
        diff + 27
    } else {
        let a_bytes = 'a'.to_string().as_bytes()[0];
        let diff = (c_bytes - a_bytes) as u32;
        diff + 1
    }
}

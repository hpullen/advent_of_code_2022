mod util;

fn main() {
    let data = util::get_file_contents();
    let cals = calories_per_elf(data);

    let res1 = cals[0];
    println!("Part 1: {}", res1);

    let mut res2 = 0;
    for cal in cals[0..3].iter() {
        res2 += cal;
    }
    println!("Part 2: {}", res2);
}

fn calories_per_elf(data: String) -> Vec<i32> {
    let mut result = Vec::new();
    let mut current_cals = 0;
    for cal in data.split('\n') {
        if !cal.is_empty() {
            let cal_val: i32 = cal.parse().unwrap();
            current_cals += cal_val;
        } else {
            result.push(current_cals);
            current_cals = 0;
        }
    }
    result.push(current_cals);
    result.sort();
    result.reverse();
    result
}

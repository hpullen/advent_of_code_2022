mod monkey;
mod util;

fn main() {
    let data = util::get_file_contents_vec();
    part_two(&data);
}

fn part_two(data: &Vec<String>) {
    let mut monkeys = monkey::Monkeys::new(&data);
    for i in 1..=10000 {
        monkeys.process_round();
        if i % 1000 == 0 {
            println!("{:?}", monkeys.get_inspections());
        }
    }
    let mut inspections = monkeys.get_inspections();
    inspections.sort();
    let answer = inspections[inspections.len() - 1] * inspections[inspections.len() - 2];
    println!("Part 2: {}", answer);
}

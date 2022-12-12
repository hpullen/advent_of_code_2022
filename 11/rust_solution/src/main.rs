mod monkey;
mod util;

fn main() {
    let data = util::get_file_contents_vec();
    part_one(&data);
    part_two(&data);
}

fn part_one(data: &Vec<String>) {
    let mut monkeys = monkey::load_monkeys(&data);
    let sum = count_inspections(&mut monkeys, 20, 3);
    println!("Part 1: {}", sum);
}

fn part_two(data: &Vec<String>) {
    let mut monkeys = monkey::load_monkeys(&data);
    let sum = count_inspections(&mut monkeys, 10000, 1);
    println!("Part 2: {}", sum);
}

fn count_inspections(monkeys: &mut Vec<monkey::Monkey>, n_rounds: u128, worry_div: u128) -> usize {
    let mut inspections = vec![0; monkeys.len()];
    for r in 0..n_rounds {
        for j in 0..monkeys.len() {
            inspections[j] += monkeys[j].items.len();
            let to_throw = monkeys[j].take_turn(worry_div);
            for (n, items) in to_throw {
                for item in items {
                    monkeys[n].add_item(item);
                }
            }
        }
        if r % 1000 == 0 {
            println!("Round {}: {:?}", r, inspections);
        }
    }
    inspections.sort();
    inspections[inspections.len() - 1] * inspections[inspections.len() - 2]
}

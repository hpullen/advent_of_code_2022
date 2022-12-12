use eval::eval;
use std::collections::HashMap;

#[derive(Debug)]
pub struct Monkey {
    pub items: Vec<u128>,
    operation: String,
    divisor: u128,
    monkey_true: usize,
    monkey_false: usize,
}

impl Monkey {
    pub fn take_turn(&mut self, worry_div: u128) -> HashMap<usize, Vec<u128>> {
        let mut to_throw = HashMap::new();
        while self.items.len() > 0 {
            self.process_item(worry_div, &mut to_throw);
        }
        to_throw
    }

    fn process_item(&mut self, worry_div: u128, to_throw: &mut HashMap<usize, Vec<u128>>) {
        let init = self.items[0];
        self.items = self.items[1..].to_vec();
        let mut item = self.apply_operation(init);
        item /= worry_div;
        let throw_to = if item % self.divisor == 0 {
            self.monkey_true
        } else {
            self.monkey_false
        };
        to_throw.entry(throw_to).or_insert(Vec::new()).push(item);
    }

    pub fn add_item(&mut self, val: u128) {
        self.items.push(val);
    }

    fn apply_operation(&self, val: u128) -> u128 {
        let to_eval = self.operation.replace("old", &val.to_string());
        println!("{}", to_eval);
        let result = eval(&to_eval).unwrap().as_u64().unwrap();
        result as u128
    }
}

pub fn load_monkeys(data: &Vec<String>) -> Vec<Monkey> {
    let mut idx = 0;
    let mut monkeys = Vec::new();
    while idx <= data.len() {
        let items = load_items(&data[idx + 1]);
        let operation = load_operation(&data[idx + 2]);
        let divisor = get_num_from_line(&data[idx + 3]);
        let monkey_true = get_num_from_line(&data[idx + 4]);
        let monkey_false = get_num_from_line(&data[idx + 5]);
        monkeys.push(Monkey {
            items,
            operation,
            divisor,
            monkey_true: monkey_true as usize,
            monkey_false: monkey_false as usize,
        });
        idx += 7;
    }
    monkeys
}

fn load_items(line: &str) -> Vec<u128> {
    let line_no_comma = line.replace(",", "");
    let line_split: Vec<&str> = line_no_comma.split_ascii_whitespace().collect();
    let mut items = Vec::new();
    for item in line_split[2..].iter() {
        let item_val = item.parse::<u128>().unwrap();
        items.push(item_val);
    }
    items
}

fn load_operation(line: &str) -> String {
    let split_line: Vec<&str> = line.split_ascii_whitespace().collect();
    split_line[3..].join(" ")
}

fn get_num_from_line(line: &str) -> u128 {
    let split_line: Vec<&str> = line.split_ascii_whitespace().collect();
    let divisor = split_line[split_line.len() - 1];
    divisor.parse::<u128>().unwrap()
}

use eval::eval;

#[derive(Debug)]
pub struct Monkeys {
    monkeys: Vec<Monkey>,
}

#[derive(Debug)]
struct Monkey {
    pub items: Vec<Item>,
    operation: String,
    divisor: u32,
    monkey_true: usize,
    monkey_false: usize,
    inspections: u128,
}

#[derive(Debug, Clone)]
struct Item {
    vals: Vec<u32>,
}

impl Item {
    fn new(monkeys: &Vec<Monkey>, init_val: &u32) -> Item {
        let mut item = Item { vals: Vec::new() };
        for _ in 0..monkeys.len() {
            item.vals.push(init_val.to_owned());
        }
        item.reduce(monkeys);
        item
    }

    fn reduce(&mut self, monkeys: &Vec<Monkey>) {
        for i in 0..monkeys.len() {
            self.vals[i] = self.vals[i] % monkeys[i].divisor;
        }
    }

    fn apply_operation(&mut self, monkeys: &Vec<Monkey>, operation: &str) {
        for i in 0..monkeys.len() {
            let to_eval = operation.replace("old", &self.vals[i].to_string());
            let result = eval(&to_eval).unwrap().as_u64().unwrap();
            self.vals[i] = result as u32;
        }
        self.reduce(monkeys);
    }
}

impl Monkeys {
    pub fn new(data: &Vec<String>) -> Monkeys {
        let mut idx = 0;
        let mut monkeys = Vec::new();
        let mut items = Vec::new();

        // Load each monkey
        while idx <= data.len() {
            items.push(load_items(&data[idx + 1]));
            let operation = load_operation(&data[idx + 2]);
            let divisor = get_num_from_line(&data[idx + 3]);
            let monkey_true = get_num_from_line(&data[idx + 4]);
            let monkey_false = get_num_from_line(&data[idx + 5]);
            monkeys.push(Monkey {
                items: Vec::new(),
                operation,
                divisor,
                monkey_true: monkey_true as usize,
                monkey_false: monkey_false as usize,
                inspections: 0,
            });
            idx += 7;
        }

        // Create Item objects and attach to the monkey that currently has the item
        for i in 0..monkeys.len() {
            for item_val in items[i].iter() {
                let item = Item::new(&monkeys, item_val);
                monkeys[i].items.push(item);
            }
        }

        Monkeys { monkeys }
    }

    pub fn process_round(&mut self) {
        for i in 0..self.monkeys.len() {
            while self.monkeys[i].items.len() > 0 {
                let mut item = self.monkeys[i].pop_item();
                item.apply_operation(&self.monkeys, &self.monkeys[i].operation);
                let throw_to: usize;
                if item.vals[i] == 0 {
                    throw_to = self.monkeys[i].monkey_true;
                } else {
                    throw_to = self.monkeys[i].monkey_false;
                };
                self.monkeys[throw_to].add_item(item);
            }
        }
    }

    pub fn get_inspections(&self) -> Vec<u128> {
        let mut inspections = Vec::new();
        for monkey in self.monkeys.iter() {
            inspections.push(monkey.inspections);
        }
        inspections
    }
}

impl Monkey {
    fn pop_item(&mut self) -> Item {
        let item = self.items[0].to_owned();
        self.items = self.items[1..].to_vec();
        self.inspections += 1;
        item
    }

    fn add_item(&mut self, item: Item) {
        self.items.push(item);
    }
}

fn load_items(line: &str) -> Vec<u32> {
    let line_no_comma = line.replace(",", "");
    let line_split: Vec<&str> = line_no_comma.split_ascii_whitespace().collect();
    let mut items = Vec::new();
    for item in line_split[2..].iter() {
        let item_val = item.parse::<u32>().unwrap();
        items.push(item_val);
    }
    items
}

fn load_operation(line: &str) -> String {
    let split_line: Vec<&str> = line.split_ascii_whitespace().collect();
    split_line[3..].join(" ")
}

fn get_num_from_line(line: &str) -> u32 {
    let split_line: Vec<&str> = line.split_ascii_whitespace().collect();
    let divisor = split_line[split_line.len() - 1];
    divisor.parse::<u32>().unwrap()
}

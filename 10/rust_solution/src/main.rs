mod util;

fn main() {
    let cmds = util::get_file_contents_vec();
    let reg = run_program(&cmds);
    part_one(&reg);
    part_two(&reg);
}

fn part_one(reg: &Register) {
    let mut idx = 20;
    let mut sum = 0;
    while idx <= 220 {
        let val = reg.get(idx - 1);
        let sig = val * (idx as i32);
        idx += 40;
        sum += sig;
    }
    println!("Part 1: {}", sum);
}

fn part_two(reg: &Register) {
    let mut current_row = "".to_owned();
    let mut pos = 0;
    for cycle in &reg.entries {
        if (cycle - pos).abs() < 2 {
            current_row = format!("{}#", current_row);
        } else {
            current_row = format!("{}.", current_row);
        }
        pos += 1;
        if pos == 40 {
            pos = 0;
            println!("{}", current_row);
            current_row = "".to_string();
        }
    }
}

fn run_program(cmds: &Vec<String>) -> Register {
    let mut reg = Register::new();
    let mut x: i32 = 1;
    reg.add(x);
    for cmd in cmds {
        let cmd_split: Vec<&str> = cmd.split_ascii_whitespace().collect();
        reg.add(x);
        if cmd_split[0] == "addx" {
            let to_add = cmd_split[1].parse::<i32>();
            x += to_add.unwrap();
            reg.add(x);
        }
    }
    reg
}

struct Register {
    entries: Vec<i32>,
}

impl Register {
    fn new() -> Register {
        Register {
            entries: Vec::new(),
        }
    }

    fn add(&mut self, val: i32) {
        self.entries.push(val);
    }

    fn get(&self, idx: usize) -> i32 {
        self.entries[idx]
    }
}

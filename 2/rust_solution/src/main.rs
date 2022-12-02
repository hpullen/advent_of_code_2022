mod game_rules;
mod util;

fn main() {
    let data = split_moves(util::get_file_contents_vec());
    part_one(&data);
    part_two(&data);
}

fn split_moves(data: Vec<String>) -> Vec<Vec<String>> {
    let mut output = Vec::new();
    for line in data {
        output.push(line.split_whitespace().map(str::to_string).collect());
    }
    output
}

fn get_their_move(their_move: &str) -> game_rules::Move {
    match their_move {
        "A" => game_rules::Move::Rock,
        "B" => game_rules::Move::Paper,
        "C" => game_rules::Move::Scissors,
        _ => panic!("Unrecognised code for their move: {}", their_move),
    }
}

fn get_your_move(your_move: &str) -> game_rules::Move {
    match your_move {
        "X" => game_rules::Move::Rock,
        "Y" => game_rules::Move::Paper,
        "Z" => game_rules::Move::Scissors,
        _ => panic!("Unrecognised code for your move: {}", your_move),
    }
}

fn part_one(data: &Vec<Vec<String>>) {
    let mut score = 0;
    for line in data {
        let their_move = get_their_move(&line[0]);
        let your_move = get_your_move(&line[1]);
        score += game_rules::get_score_from_moves(their_move, your_move);
    }
    println!("Part 1: {}", score);
}

fn get_outcome(your_move: &str) -> game_rules::Outcome {
    match your_move {
        "X" => game_rules::Outcome::Loss,
        "Y" => game_rules::Outcome::Draw,
        "Z" => game_rules::Outcome::Win,
        _ => panic!("Unrecognised code for outcome: {}", your_move),
    }
}

fn part_two(data: &Vec<Vec<String>>) {
    let mut score = 0;
    for line in data {
        let their_move = get_their_move(&line[0]);
        let outcome = get_outcome(&line[1]);
        score += game_rules::get_score_from_outcome(their_move, outcome);
    }
    println!("Part 2: {}", score);
}

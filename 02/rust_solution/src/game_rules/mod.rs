#[derive(Debug, PartialEq, Eq, Clone, Copy)]
pub enum Move {
    Rock,
    Paper,
    Scissors,
}

pub enum Outcome {
    Loss,
    Draw,
    Win,
}

fn wins_against(your_move: &Move) -> Move {
    match your_move {
        Move::Rock => Move::Scissors,
        Move::Paper => Move::Rock,
        Move::Scissors => Move::Paper,
    }
}

fn loses_against(your_move: &Move) -> Move {
    match your_move {
        Move::Rock => Move::Paper,
        Move::Paper => Move::Scissors,
        Move::Scissors => Move::Rock,
    }
}

fn get_move_score(your_move: &Move) -> u32 {
    match your_move {
        Move::Rock => 1,
        Move::Paper => 2,
        Move::Scissors => 3,
    }
}

pub fn get_score_from_moves(their_move: Move, your_move: Move) -> u32 {
    let mut score = get_move_score(&your_move);
    if wins_against(&your_move) == their_move {
        score += 6;
    } else if your_move == their_move {
        score += 3;
    }
    score
}

pub fn get_score_from_outcome(their_move: Move, outcome: Outcome) -> u32 {
    let your_move = match outcome {
        Outcome::Loss => wins_against(&their_move),
        Outcome::Draw => their_move,
        Outcome::Win => loses_against(&their_move),
    };
    get_score_from_moves(their_move, your_move)
}

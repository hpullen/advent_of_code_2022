mod util;

fn main() {
    let data = util::get_file_contents_vec();
    let tree_map = load_tree_map(data);
    part_one(&tree_map);
    part_two(&tree_map);
}

fn part_one(tree_map: &Vec<Vec<u32>>) {
    let mut count = 0;
    let mut vis_map = Vec::new();
    for i in 0..tree_map.len() {
        let mut vis_line = Vec::new();
        for j in 0..tree_map[0].len() {
            let vis = is_visible(tree_map, i, j);
            vis_line.push(vis);
            if vis {
                count += 1;
            }
        }
        vis_map.push(vis_line);
    }
    println!("Part 1: {}", count);
}

fn part_two(tree_map: &Vec<Vec<u32>>) {
    let mut best_score = 0;
    for i in 0..tree_map.len() {
        for j in 0..tree_map[0].len() {
            let score = get_viewing_score(tree_map, i, j);
            if score > best_score {
                best_score = score;
            }
        }
    }
    println!("Part 2: {}", best_score);
}

fn get_viewing_score(tree_map: &Vec<Vec<u32>>, i: usize, j: usize) -> u32 {
    let mut score = 1;
    let surrounding_trees = get_surrounding_trees(tree_map, i, j);
    for trees in surrounding_trees {
        let mut count = 0;
        for tree in trees {
            count += 1;
            if tree >= tree_map[i][j] {
                break;
            }
        }
        score *= count;
    }
    score
}

fn load_tree_map(data: Vec<String>) -> Vec<Vec<u32>> {
    let mut result = Vec::new();
    for line in data {
        let mut result_line = Vec::new();
        for c in line.chars() {
            result_line.push(c.to_digit(10).unwrap());
        }
        result.push(result_line);
    }
    result
}

fn is_visible(tree_map: &Vec<Vec<u32>>, i: usize, j: usize) -> bool {
    let surrounding_trees = get_surrounding_trees(tree_map, i, j);
    for trees in surrounding_trees {
        let mut taller_than_all = true;
        for tree in trees {
            if tree >= tree_map[i][j] {
                taller_than_all = false;
            }
        }
        if taller_than_all {
            return true;
        }
    }
    false
}

fn get_surrounding_trees(tree_map: &Vec<Vec<u32>>, i: usize, j: usize) -> Vec<Vec<u32>> {
    let mut trees = vec![Vec::new(), Vec::new(), Vec::new(), Vec::new()];
    for ii in 0..tree_map.len() {
        for jj in 0..tree_map[0].len() {
            if ii == i {
                if jj < j {
                    trees[0].push(tree_map[ii][jj]);
                } else if jj > j {
                    trees[1].push(tree_map[ii][jj]);
                }
            } else if jj == j {
                if ii < i {
                    trees[2].push(tree_map[ii][jj]);
                } else if ii > i {
                    trees[3].push(tree_map[ii][jj]);
                }
            }
        }
    }
    trees[0].reverse();
    trees[2].reverse();
    trees
}

fn print<T: std::fmt::Display>(data: &Vec<Vec<T>>) {
    for line in data {
        let mut to_print = "".to_owned();
        for item in line {
            if to_print.len() == 0 {
                to_print = format!("{}", item);
            } else {
                to_print = format!("{} {}", to_print, item);
            }
        }
        println!("{}", to_print);
    }
}

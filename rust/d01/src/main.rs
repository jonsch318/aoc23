fn main() {
    let input: Vec<_> = include_bytes!("../../../input/d01/input.txt")
        .split(|b| b == &b'\n')
        .collect();
    println!("{}", run_p1(input));
}

fn run_p1(lines: Vec<&[u8]>) -> usize {
    lines
        .iter()
        .map(|line| p1_get_first_num(line) * 10 + p1_get_reverse_num(line))
        .sum()
}

fn p1_get_first_num(line: &[u8]) -> usize {
    for b in line {
        if (48..58).contains(b) {
            return (b - 48) as usize;
        }
    }
    0
}

fn p1_get_reverse_num(line: &[u8]) -> usize {
    for b in line.iter().rev() {
        if (48..58).contains(b) {
            return (b - 48) as usize;
        }
    }
    0
}

use std::fs;

fn main() {
    let input = fs::read_to_string("./input.txt").expect("No input.txt file found.");
    let lines = input.split("\r\n");

    let mut first_list: Vec<i32> = Vec::new();
    let mut second_list: Vec<i32> = Vec::new();

    for line in lines {
        let parts: Vec<&str> = line.split("   ").collect();
        first_list.push(parts[0].parse().unwrap());
        second_list.push(parts[1].parse().unwrap());
    }

    first_list.sort();
    second_list.sort();

    let mut part1 = 0;
    for i in 0..first_list.len() {
        part1 += (first_list[i] - second_list[i]).abs();
    }
    println!("Part 1: {}", part1);

    let mut part2 = 0;
    let mut cached_result = 0;
    for first in first_list.iter() {
        if *first == cached_result {
            part2 += cached_result;
            continue;
        }
        let mut count = 0;
        let mut skip = 0;
        for second in second_list.iter().skip(skip) {
            skip+=1;
            if first == second {
                count+=1;
            }
        }
        cached_result = first*count;
        part2 += cached_result;
    }
    println!("Part 2: {}", part2);
}

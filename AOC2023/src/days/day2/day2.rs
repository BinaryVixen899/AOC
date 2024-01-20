use std::{default, str::FromStr};
// static PUZZLE_INPUT: str = "awerawera";
use strum::{Display, EnumString};

pub struct Game {
    turns: Vec<Turn>,
    possible: Option<bool>,
    id: Option<i32>,
    required_red_cubes: Colour,
    required_green_cubes: Colour,
    required_blue_cubes: Colour,
}

impl Game {
    fn determine_possibility(&mut self, mut bag: &mut Bag, id: i32) {
        for turn in &self.turns {
            if turn.red_cubes <= bag.red_cubes
                && turn.green_cubes <= bag.green_cubes
                && turn.blue_cubes <= bag.blue_cubes
            {
            } else {
                println!("Game #{} is not possible!", id);
                println!(
                    "tr:{} br:{}\n\
                    tg:{} bg:{}\n\
                    tb:{} bb:{}
                ",
                    turn.red_cubes,
                    bag.red_cubes,
                    turn.green_cubes,
                    bag.green_cubes,
                    turn.blue_cubes,
                    bag.blue_cubes
                );
                self.possible = Some(false);
                return;
            }
        }
        self.possible = Some(true);
        bag.score += id;
    }

    fn calculate_required_cubes(&mut self) {
        for turn in &self.turns {
            if turn.red_cubes >= self.required_red_cubes {
                self.required_red_cubes = turn.red_cubes
            }
            if turn.green_cubes >= self.required_green_cubes {
                self.required_green_cubes = turn.green_cubes
            }
            if turn.blue_cubes >= self.required_blue_cubes {
                self.required_blue_cubes = turn.blue_cubes
            }
        }
    }
}
impl Default for Game {
    fn default() -> Self {
        return Game {
            turns: vec![Turn::default()],
            possible: None,
            id: None,
            required_red_cubes: Colour::Red(0),
            required_green_cubes: Colour::Green(0),
            required_blue_cubes: Colour::Blue(0),
        };
    }
}

struct Turn {
    red_cubes: Colour,
    green_cubes: Colour,
    blue_cubes: Colour,
}
impl Default for Turn {
    fn default() -> Self {
        return Turn {
            red_cubes: Colour::Red(0),
            green_cubes: Colour::Green(0),
            blue_cubes: Colour::Blue(0),
        };
    }
}

struct Bag {
    red_cubes: Colour,
    green_cubes: Colour,
    blue_cubes: Colour,
    score: i32,
}

pub fn convert_input(input: &str) -> Vec<Game> {
    let lines = input.lines();
    let mut games = vec![];

    for line in lines {
        println!("Current Line: {:?}", line);
        if let Some(l) = line.split_once(": ") {
            // Could simply this further, pass every turn to a function or something... But I think the way I currently have it is legible and I am not sure if making it a one liner would improve performance.
            let turns: Vec<&str> = l.1.split(";").map(|t| t.trim()).collect();
            println!("Turns: {:?}", turns);
            let mut game_turns = vec![];
            for turn in turns {
                println!("Current Turn: {:?}", turn);
                let mut t = Turn::default();
                //TODO: Optimize
                //On the other paw, minimizing this may be well worth it. Unlike turns, we ultimately do not need to keep track of draws beyond what they consist of
                let draws: Vec<&str> = turn.split(",").map(|x| x.trim()).collect();
                println!("Draws: {:?}", draws);
                for draw in draws {
                    println!("Current Draw: {:?}", draw);
                    // hacks
                    if let Some(item) = draw.split_once(" ") {
                        let colour = Colour::from_str(item.1).expect("A valid colour");
                        let count = i32::from_str(item.0).expect("Conversion worked");
                        match colour {
                            Colour::Blue(_) => t.blue_cubes = Colour::Blue(count),
                            Colour::Red(_) => t.red_cubes = Colour::Red(count),
                            Colour::Green(_) => t.green_cubes = Colour::Green(count),
                        }
                    }
                }
                // I wonder if we could eliminate turns as well
                game_turns.push(t);
            }
            let current_game = Game {
                turns: game_turns,
                possible: None,
                ..Default::default()
            };
            games.push(current_game);
            println!()
        }

        // annnnnd this is where we do the regex thing
    }
    return games;
}

#[derive(Display, Debug, PartialEq, EnumString, PartialOrd, Clone, Copy)]
enum Colour {
    #[strum(serialize = "red", serialize = "r")]
    Red(i32),
    #[strum(serialize = "green", serialize = "g")]
    Green(i32),
    #[strum(serialize = "blue", serialize = "b")]
    Blue(i32),
}

impl Into<i32> for Colour {
    fn into(self) -> i32 {
        match self {
            Colour::Red(count) => return count,
            Colour::Green(count) => return count,
            Colour::Blue(count) => return count,
        }
    }
}
pub fn part1() {
    // print!(PUZZLE_INPUT);
    // let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
    // Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
    // Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
    // Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
    // Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";
    let input = "Game 1: 1 green, 1 blue, 1 red; 3 green, 1 blue, 1 red; 4 green, 3 blue, 1 red; 4 green, 2 blue, 1 red; 3 blue, 3 green
    Game 2: 9 blue, 7 red; 5 blue, 6 green, 1 red; 2 blue, 10 red, 9 green; 3 green, 14 red, 5 blue; 8 red, 3 blue, 4 green; 8 green, 14 blue, 10 red
    Game 3: 11 green, 8 blue, 7 red; 3 green, 4 red, 9 blue; 3 red, 4 green, 8 blue; 11 green, 1 red, 16 blue
    Game 4: 3 blue, 20 green, 2 red; 1 green, 3 red, 3 blue; 1 blue, 9 green; 4 red, 17 green; 12 green, 3 red
    Game 5: 2 red, 1 blue, 4 green; 6 blue, 2 green; 2 red, 5 green
    Game 6: 1 green, 7 red; 1 blue, 3 green, 1 red; 1 blue, 2 red, 2 green; 1 blue, 1 green, 2 red; 3 red; 8 red, 1 green, 1 blue
    Game 7: 13 green, 5 blue, 1 red; 9 green, 6 blue; 11 green, 2 blue, 2 red
    Game 8: 2 red, 11 green, 6 blue; 5 green, 3 blue; 3 blue, 3 green, 5 red; 5 blue, 9 green, 1 red
    Game 9: 3 blue, 5 green, 8 red; 4 green, 19 blue, 4 red; 4 red, 17 blue
    Game 10: 2 green, 8 red, 4 blue; 1 green, 5 red, 9 blue; 3 green, 2 red, 4 blue; 2 green, 5 red, 2 blue; 6 green, 4 blue; 10 blue, 8 green, 8 red
    Game 11: 3 green, 1 blue, 9 red; 2 blue, 1 red, 9 green; 1 blue, 9 green, 7 red; 8 red, 6 green
    Game 12: 5 green; 8 green, 7 red, 1 blue; 8 blue, 8 green, 14 red; 6 red, 14 green; 14 green, 3 red, 8 blue
    Game 13: 7 red, 2 green, 4 blue; 5 red, 3 blue, 8 green; 10 green, 1 red, 4 blue; 7 green, 1 red, 13 blue; 11 green, 12 blue, 2 red
    Game 14: 9 green, 4 red; 7 green, 4 blue, 5 red; 2 red, 2 green; 3 green, 2 red, 8 blue; 7 green, 6 red, 8 blue
    Game 15: 19 blue, 1 green; 1 red, 5 blue; 3 green, 8 blue; 1 red, 13 blue, 3 green
    Game 16: 8 red, 11 blue, 3 green; 14 green, 1 red, 12 blue; 6 green, 1 red, 6 blue; 1 red, 6 blue, 17 green; 2 green, 8 blue, 3 red
    Game 17: 3 red, 1 blue; 1 blue, 2 red, 1 green; 1 red; 3 red, 2 green
    Game 18: 8 green, 2 red; 1 blue, 6 red; 8 green, 7 red, 2 blue; 1 red; 6 green, 1 blue, 3 red
    Game 19: 4 blue, 2 green; 4 green, 18 blue, 2 red; 14 green, 4 blue; 1 red, 3 blue, 18 green; 11 blue, 3 red; 14 green, 4 red, 6 blue
    Game 20: 7 blue; 1 blue, 6 green, 1 red; 1 red, 3 blue, 10 green; 7 green, 4 blue, 1 red; 6 green, 6 blue, 1 red; 1 red, 5 blue, 17 green
    Game 21: 11 blue, 9 red; 8 red, 2 blue; 2 red, 11 blue, 2 green
    Game 22: 4 green, 9 blue, 2 red; 2 blue, 8 green; 2 green, 2 red, 6 blue
    Game 23: 7 green, 7 blue; 6 blue, 11 green; 1 red, 14 green; 15 green, 4 blue, 1 red; 1 red, 5 blue, 3 green; 1 red, 1 blue, 13 green
    Game 24: 5 green, 2 red, 2 blue; 1 blue, 3 green, 2 red; 2 blue, 7 green, 3 red
    Game 25: 16 red, 8 green; 2 red, 3 blue; 10 green, 5 red, 4 blue; 9 red, 7 green; 7 red, 6 blue
    Game 26: 3 red, 1 green; 5 red, 1 blue, 10 green; 8 red, 5 green
    Game 27: 3 red, 2 blue; 6 red, 8 green; 5 green, 13 red, 2 blue; 4 red, 1 blue, 8 green; 14 red, 1 blue, 3 green; 2 red, 1 green, 2 blue
    Game 28: 9 red, 10 blue; 9 red, 9 blue; 1 green, 6 red, 4 blue; 12 blue, 3 green, 2 red; 2 green, 12 red, 8 blue
    Game 29: 4 red, 15 blue; 5 blue, 3 green, 6 red; 1 green, 9 blue, 1 red; 5 green, 8 red, 1 blue
    Game 30: 4 green, 2 blue, 10 red; 1 red, 5 green, 6 blue; 15 red, 2 blue; 5 green, 10 red, 8 blue
    Game 31: 6 green, 2 blue, 2 red; 5 green, 6 red, 6 blue; 3 blue, 5 red, 1 green; 4 green, 6 red, 9 blue; 4 red; 3 green, 1 red, 3 blue
    Game 32: 8 green, 17 red, 17 blue; 11 red, 6 green, 13 blue; 14 red, 1 green, 1 blue; 1 green, 17 red, 4 blue; 5 green, 14 red, 15 blue; 15 blue, 8 green
    Game 33: 2 red, 14 blue; 3 blue, 17 red, 4 green; 9 blue, 8 red; 5 red, 2 blue; 4 green, 16 red, 5 blue; 15 blue, 6 green, 17 red
    Game 34: 12 green, 2 red, 1 blue; 3 blue, 9 red, 13 green; 2 blue, 17 green, 6 red; 6 green, 4 red, 3 blue; 2 red, 1 blue; 7 green, 3 blue, 7 red
    Game 35: 4 blue, 5 red, 5 green; 6 green, 12 red, 6 blue; 3 green, 18 red, 2 blue; 13 red, 6 green, 9 blue; 3 green, 10 blue, 17 red; 1 green, 3 blue, 16 red
    Game 36: 4 green, 3 blue, 1 red; 3 green, 3 red, 10 blue; 1 red, 8 green, 8 blue; 3 blue, 1 red; 2 red, 2 blue, 14 green
    Game 37: 16 blue, 1 green; 9 blue; 7 red, 13 blue
    Game 38: 6 green, 8 red, 3 blue; 5 blue, 4 green, 6 red; 5 blue, 4 green; 5 red, 3 green, 1 blue; 6 green, 4 blue, 15 red
    Game 39: 10 blue, 4 green; 1 blue, 7 green, 5 red; 8 red, 2 blue
    Game 40: 2 blue, 2 green, 6 red; 8 green, 4 red, 2 blue; 4 blue, 12 green, 6 red
    Game 41: 4 green, 2 blue, 11 red; 6 red, 11 green; 4 blue, 2 red, 19 green; 3 green, 2 blue, 1 red
    Game 42: 2 blue, 2 green, 4 red; 1 red, 4 blue, 8 green; 5 red, 2 blue, 15 green; 10 green, 5 red, 1 blue; 10 green, 1 blue; 2 blue
    Game 43: 10 red, 19 green, 11 blue; 11 green, 1 red, 2 blue; 13 red, 6 blue, 3 green; 12 red, 10 green; 5 red, 19 green, 8 blue; 2 red, 10 green, 3 blue
    Game 44: 7 blue, 8 red; 1 green, 12 red; 19 red, 3 green, 5 blue
    Game 45: 12 red; 2 green, 5 red, 3 blue; 10 green, 2 blue, 4 red; 10 green, 7 red
    Game 46: 4 blue, 4 red, 2 green; 7 green, 6 blue; 6 blue, 1 red, 4 green
    Game 47: 4 green, 8 red, 1 blue; 4 green, 1 blue, 11 red; 14 red, 3 blue, 10 green; 15 green, 2 blue, 7 red
    Game 48: 6 blue, 3 green, 1 red; 15 blue, 11 red, 3 green; 17 blue, 14 red; 2 green, 13 red; 5 red, 2 green, 4 blue
    Game 49: 7 blue, 1 green; 8 red, 2 blue, 1 green; 1 red, 1 green, 2 blue; 3 red, 2 blue, 1 green
    Game 50: 13 red, 2 green, 2 blue; 13 red, 6 green, 1 blue; 12 red, 8 green; 1 red, 3 green; 13 red; 2 blue, 11 red, 2 green
    Game 51: 7 green, 4 blue, 2 red; 3 red, 7 green, 5 blue; 10 green, 2 blue; 14 green, 3 red, 4 blue; 2 blue, 2 red, 10 green
    Game 52: 7 blue, 10 red; 7 red, 4 blue, 8 green; 12 red, 4 blue, 7 green; 7 green, 7 red; 17 green, 11 blue, 6 red; 8 green, 8 red, 18 blue
    Game 53: 6 green, 2 red; 10 red, 13 green; 2 blue, 3 green, 5 red; 4 red, 4 green; 8 green, 1 red; 13 green, 2 blue, 10 red
    Game 54: 5 red, 9 green, 5 blue; 6 red, 15 green, 10 blue; 8 blue, 3 green, 1 red; 12 blue, 3 red, 13 green
    Game 55: 10 green, 4 red, 2 blue; 2 green, 1 red, 2 blue; 2 blue, 4 red, 8 green; 5 blue, 3 green
    Game 56: 7 green, 9 red, 2 blue; 4 red, 1 blue, 3 green; 3 red, 4 blue, 6 green; 7 green, 2 red; 5 blue, 2 red, 3 green; 7 green, 7 red, 5 blue
    Game 57: 11 blue, 5 green, 6 red; 18 red, 15 green, 10 blue; 5 green, 14 red, 6 blue; 1 green, 11 red, 7 blue; 11 blue, 7 red, 12 green
    Game 58: 9 red, 6 blue, 6 green; 6 blue, 12 red, 3 green; 8 red, 1 blue, 20 green; 3 green, 3 red, 15 blue; 4 blue, 15 green, 3 red
    Game 59: 18 red, 4 blue, 7 green; 11 blue, 19 red, 7 green; 10 red, 9 blue, 1 green; 8 red, 12 green, 4 blue; 6 green, 5 blue, 12 red; 2 blue, 14 green, 2 red
    Game 60: 1 blue, 9 green, 6 red; 1 red, 1 blue, 13 green; 15 green, 4 red; 1 blue, 2 red, 15 green
    Game 61: 9 green, 3 red, 2 blue; 1 green, 5 blue, 10 red; 17 red, 9 green, 5 blue; 10 red, 5 green, 5 blue
    Game 62: 4 red, 2 green; 2 red; 5 red, 2 green, 2 blue; 3 green, 1 blue; 2 blue, 3 red, 3 green
    Game 63: 4 red, 6 blue, 2 green; 3 green, 1 red, 5 blue; 7 blue, 5 green
    Game 64: 8 blue, 12 red; 1 green, 6 red, 14 blue; 12 red, 13 blue
    Game 65: 2 red, 8 green; 1 blue, 2 red, 5 green; 2 blue, 3 green; 1 green, 1 red
    Game 66: 6 red, 8 blue, 2 green; 6 blue, 7 green; 7 green, 11 blue; 5 green, 4 red, 10 blue; 5 blue, 4 green; 6 blue, 6 green, 5 red
    Game 67: 12 green, 4 red; 2 blue, 11 green, 6 red; 9 red, 2 green, 6 blue; 2 red, 8 blue, 18 green; 17 green, 7 blue, 6 red; 12 blue, 9 green
    Game 68: 3 red, 9 blue, 1 green; 3 green, 11 blue; 12 blue, 9 red; 6 red, 13 green, 8 blue
    Game 69: 3 red, 8 green, 3 blue; 6 green, 3 red; 11 green, 3 blue; 4 red, 3 green; 7 green, 4 blue, 6 red; 1 red, 2 blue
    Game 70: 6 green, 4 blue; 7 red, 9 green, 14 blue; 12 red; 9 green, 10 red; 6 green, 16 blue, 8 red
    Game 71: 4 blue, 6 red, 9 green; 6 green, 2 red; 8 green, 4 blue, 2 red; 1 red, 3 blue, 8 green; 9 green, 3 red, 3 blue; 4 red, 6 green
    Game 72: 3 red, 3 green, 3 blue; 4 red, 1 green, 3 blue; 2 red, 2 green, 1 blue
    Game 73: 7 green, 6 red, 7 blue; 2 green, 4 blue; 2 blue, 15 green, 4 red; 1 blue, 4 green, 2 red; 7 blue, 14 green
    Game 74: 4 green, 2 red, 2 blue; 9 blue, 13 green, 4 red; 10 green, 12 blue, 7 red; 4 blue, 8 green, 7 red
    Game 75: 3 red, 3 green; 3 green, 12 red; 18 red, 2 blue; 3 green, 9 red, 1 blue; 14 red, 1 green; 15 red
    Game 76: 1 blue, 7 red, 8 green; 3 blue, 4 red, 1 green; 6 green, 6 red
    Game 77: 2 green, 8 red; 5 green, 11 red, 1 blue; 5 red, 2 blue, 2 green; 6 red, 5 green, 2 blue; 11 red, 2 green, 1 blue; 2 red, 8 green, 2 blue
    Game 78: 1 blue, 4 red, 10 green; 13 green, 4 red, 9 blue; 12 green, 6 blue, 3 red; 5 blue, 8 green, 6 red
    Game 79: 9 red, 1 blue, 17 green; 9 green, 6 red; 15 green, 1 blue, 9 red; 1 blue, 8 red, 12 green; 11 green, 1 blue, 1 red
    Game 80: 3 red, 3 blue, 1 green; 1 blue, 8 green; 10 green, 16 blue
    Game 81: 15 blue, 2 red; 1 red, 8 blue; 2 green, 7 red, 11 blue; 19 blue, 8 red, 2 green; 20 red, 19 blue
    Game 82: 3 red, 17 blue, 9 green; 10 red, 2 green, 17 blue; 13 red, 3 blue, 10 green; 11 red, 10 green, 18 blue; 1 green, 12 blue, 9 red; 3 red, 10 blue, 8 green
    Game 83: 4 green, 2 blue, 14 red; 7 red, 2 blue, 7 green; 16 red, 7 green; 16 red, 2 green; 3 blue, 4 green, 15 red; 6 red, 1 blue, 4 green
    Game 84: 4 blue, 1 green, 2 red; 7 blue, 6 red; 1 blue, 4 red, 3 green
    Game 85: 5 blue, 1 red, 4 green; 14 blue, 7 green; 9 blue, 1 red, 7 green; 15 blue, 9 green; 8 blue, 6 green, 1 red
    Game 86: 12 red, 12 blue, 7 green; 16 red, 11 green, 4 blue; 14 blue, 8 green, 8 red; 2 blue, 15 red, 6 green; 1 green, 6 red, 5 blue; 11 blue, 9 green
    Game 87: 4 blue, 2 green, 6 red; 8 red, 3 green, 5 blue; 10 red, 1 green; 1 green, 3 blue, 1 red
    Game 88: 2 green, 2 red, 4 blue; 4 blue, 4 green, 12 red; 2 green, 3 blue, 4 red; 2 green, 2 blue, 12 red; 4 blue, 8 red, 2 green
    Game 89: 13 blue, 1 red, 5 green; 8 green, 16 blue, 5 red; 12 green, 2 red, 18 blue
    Game 90: 7 red, 11 blue, 1 green; 8 green, 13 blue; 7 green, 16 blue; 5 green, 13 red, 11 blue; 5 blue, 10 green, 3 red
    Game 91: 1 blue; 1 blue, 3 green; 1 green, 2 red
    Game 92: 16 red, 4 blue, 14 green; 15 red, 7 blue, 13 green; 7 green, 13 red, 8 blue; 4 blue, 9 red, 5 green; 6 red, 7 blue, 8 green; 14 green, 7 red, 10 blue
    Game 93: 2 red, 6 blue, 7 green; 8 green, 3 red, 10 blue; 1 green, 4 red; 2 red, 2 green; 8 blue, 7 green
    Game 94: 2 green, 3 blue, 5 red; 10 blue; 1 green, 7 red; 3 blue, 1 green, 12 red
    Game 95: 13 blue, 5 red; 9 blue, 3 red, 7 green; 10 green, 4 red, 12 blue; 14 blue; 7 green, 2 blue, 1 red
    Game 96: 5 red, 2 blue; 4 red; 1 green, 2 blue
    Game 97: 9 red, 10 green, 3 blue; 2 green, 15 red, 1 blue; 2 blue, 16 green, 16 red; 8 green, 14 red; 16 red
    Game 98: 18 green, 16 red, 1 blue; 3 red, 2 blue, 20 green; 1 blue, 20 green, 14 red; 14 red, 2 green
    Game 99: 7 red, 9 green, 5 blue; 6 blue, 1 green; 4 green, 5 blue, 1 red; 8 green, 6 red, 7 blue; 1 blue, 2 red, 9 green
    Game 100: 10 blue, 2 red; 7 green, 20 blue, 9 red; 8 red, 6 green, 2 blue";

    let mut mybag = Bag {
        red_cubes: Colour::Red(12),
        green_cubes: Colour::Green(13),
        blue_cubes: Colour::Blue(14),
        score: 0,
    };
    let mut games = convert_input(input);
    for (id, game) in games.iter_mut().enumerate() {
        game.id = Some(id as i32 + 1);
        game.determine_possibility(&mut mybag, game.id.unwrap());
    }
    println!("The sum of the ids is {}", mybag.score);
    part2(games);
}

pub fn part2(games: Vec<Game>) {
    let mut sum: i32 = 0;
    for mut game in games {
        game.calculate_required_cubes();
        let red_cubes: i32 = game.required_red_cubes.into();
        let green_cubes: i32 = game.required_green_cubes.into();
        let blue_cubes: i32 = game.required_blue_cubes.into();
        let power: i32 = red_cubes * green_cubes * blue_cubes;
        println!(
            "required_red_cubes:{}\n\
            required_green_cubes:{}\n\
            required_blue_cubes:{}
        ",
            red_cubes, green_cubes, blue_cubes
        );
        sum += power
    }
    println!("The power of the cubes is {}", sum)

    // figure out the largest number of each type of cube played per game. This is the minimum
    // then multiply these together, that is the power for the game
    // then add them all together
}

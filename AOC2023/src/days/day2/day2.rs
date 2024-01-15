// static PUZZLE_INPUT: str = "awerawera";

#[derive(Default)]
struct Game {
    turns: Vec<Turn>,
    possible: Option<bool>,
    id: Option<i32>,
}

impl Game {
    fn determine_possibility(&mut self, mut bag: &mut Bag, id: i32) {
        for turn in &self.turns {
            if turn.red <= bag.red && turn.green <= bag.green && turn.blue <= bag.blue {
            } else {
                println!("Game #{} is not possible!", id);
                println!(
                    "tr:{} br:{}
                    tg:{} bg:{}
                    tb:{} bb{}
                ",
                    turn.red, bag.red, turn.green, bag.green, turn.blue, bag.blue
                );
                self.possible = Some(false);
                return;
            }
        }
        self.possible = Some(true);
        bag.score += id;
    }
}

struct Turn {
    red: i32,
    green: i32,
    blue: i32,
}

struct Bag {
    red: i32,
    green: i32,
    blue: i32,
    score: i32,
}

pub fn part1() {
    // print!(PUZZLE_INPUT);
    let mut mybag = Bag {
        red: 12,
        green: 13,
        blue: 14,
        score: 0,
    };
    // Need something that does text comprehension for these, perhaps using the JSON library?
    let game1 = Game {
        turns: {
            vec![
                Turn {
                    red: 4,
                    blue: 3,
                    green: 0,
                },
                Turn {
                    red: 0,
                    blue: 6,
                    green: 2,
                },
                Turn {
                    red: 0,
                    blue: 0,
                    green: 2,
                },
            ]
        },
        possible: None,
        id: None,
    };
    let game2 = Game {
        turns: {
            vec![
                Turn {
                    red: 0,
                    blue: 1,
                    green: 2,
                },
                Turn {
                    red: 1,
                    blue: 4,
                    green: 3,
                },
                Turn {
                    red: 0,
                    blue: 1,
                    green: 1,
                },
            ]
        },
        possible: None,
        id: None,
    };
    let game3 = Game {
        turns: {
            vec![
                Turn {
                    red: 20,
                    blue: 6,
                    green: 8,
                },
                Turn {
                    red: 4,
                    blue: 5,
                    green: 13,
                },
                Turn {
                    red: 1,
                    blue: 0,
                    green: 5,
                },
            ]
        },
        possible: None,
        id: None,
    };
    let game4 = Game {
        turns: {
            vec![
                Turn {
                    red: 3,
                    blue: 6,
                    green: 1,
                },
                Turn {
                    red: 6,
                    blue: 0,
                    green: 3,
                },
                Turn {
                    red: 14,
                    blue: 15,
                    green: 3,
                },
            ]
        },
        possible: None,
        ..Default::default()
    };
    let game5 = Game {
        turns: {
            vec![
                Turn {
                    red: 6,
                    blue: 1,
                    green: 3,
                },
                Turn {
                    red: 1,
                    blue: 2,
                    green: 2,
                },
            ]
        },
        possible: None,
        ..Default::default()
    };

    let mut games = [game1, game2, game3, game4, game5];
    for (id, game) in games.iter_mut().enumerate() {
        game.id = Some(id as i32 + 1);
        game.determine_possibility(&mut mybag, game.id.unwrap());
    }
    println!("The sum of the ids is {}", mybag.score);
}

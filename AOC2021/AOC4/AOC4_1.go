package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type BingoBoard [5][5]int

func (b BingoBoard) Evaluate(called []int) (int, int) {
	// Evaluate takes a list of called numbers and, if it is a winning board, returns the turns to win and the final score

	var lastnumbercalled int

	for turns, callednumber := range called {
		for i, row := range b {
			for j, number := range row {
				if number == callednumber {
					lastnumbercalled = number
					b[j][i] = 'T'
					if b.CheckForBingos() {
						// Alternatively, don't check the first five times but check every time after 5
						score := b.Score(lastnumbercalled)
						return score, turns
						// return score and turns to win here

					}
				}
			}

		}

	}
	return -1, -1
}

func (b BingoBoard) CheckForBingos() bool {
	var bingo bool

	// Both vertical and horizontal checking
	for _, row := range b {
		tcount := 0
		// reset the tcount every iteration
		for _, number := range row {
			if number == 'T' {
				tcount += 1
				if tcount == 5 {
					bingo = true
					return bingo
					// have to hope it doesn't just return to outermost loop
				}
			}

		}
	}
	// after we go through every row, we go through every column

	// OH, PRINTING LEN HERE WILL GIVE THE LENGTH OF THE X

	for i := 0; i <= 4; i++ {
		j := i
		tcount := 0
		for i := 0; i <= 4; i++ {
			if b[j][i] == 'T' {
				tcount += 1
				if tcount == 5 {
					bingo = true
					return bingo
				}

			}
		}
	}

	// so this would get the length of the x values
	// what i want is the length of the y values
	// except I already know it and it's five -1 = 4

	// reset the t count every iteration
	// so when we call row what we call range of b we get the 1st array in that index, the outermost array,
	// so at this point we just need to do the 1st value of every row PERFECT
	// we also need to iterate ourselves here, otherwise we get the numbers

	// Can it really be that easy?
	// Ugh. So inefficient

	// and of course, if there is nothing return, which /should/ return false assuming the zero value of bingo is false, which it is
	return bingo

}

func (b BingoBoard) Score(lastnumbercalled int) int {
	// Calculates the score of a winning board, takes the last number called
	// This looks good to me so i'm not sure what the compiler is complaining about

	var sum int

	for _, row := range b {
		for _, number := range row {
			if number != 'T' {
				sum += number
			}
		}
	}
	sum = sum * lastnumbercalled

	return sum

}

func main() {
	bytes, err := ioutil.ReadFile("bingocards.txt")
	// Read the bingo cards
	if err != nil {
		return
	}
	// Standard error checking
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]
	//extra newline on end
	//fun fact: every file in unix ends in a newline typically

	numbers := strings.Split(split[0], ",")
	// Okay so let's assume this is the first line of numbers
	var called []int
	var boards []BingoBoard
	bm := make(map[int]int)

	for _, n := range numbers {
		i, _ := strconv.Atoi(n)
		called = append(called, i)
		// well this is easy, we're just taking a file where the first line is the called #s and appending it to the called array/slice
	}

	for i := 2; i+5 < len(split); i += 6 {
		var b BingoBoard
		// i := 2 we start 2 in in, then we said "while i + 5 < len(split)".
		//split being the current split we're processing
		//and a split is a bingo board

		for r, line := range split[i : i+5] {
			c := 0
			for pos := 0; pos < len(line); pos += 3 {
				// digit by digit is what we're doing here
				n, _ := strconv.Atoi(strings.Trim(line[pos:pos+2], " "))
				//this is where we get the actual number
				// Which is why we get the actual number here
				b[r][c] = n
				// r and c, row and column
				c++
				// by iterating this at the end we ensure we get the correct column number as well everytime we move onto a new line
			}
		}

		boards = append(boards, b)
		// Okay so this is how we append each board So eventually we will have a list of boards
	}

	for i, b := range boards {
		score, turnsToWin := b.Evaluate(called)
		bm[i] = score
		print(turnsToWin)
		print(score)

		// boardmap = turnsToWin[score]
		// we need to add scores and turnsto win to boardmap here
	}

	// Here is where we go through the values of boardmap
	// figure out which board wins first, what will be your final score if you choose that board

}

// https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.html
// https://www.callicoder.com/golang-arrays/#:~:text=Arrays%20in%20Golang%20are%20value%20types%20unlike%20other%20languages%20like,the%20entire%20array%20is%20copied.

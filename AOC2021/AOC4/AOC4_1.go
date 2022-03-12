package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type BingoBoard [5][5]int 

// I can just create an interface that collects info about bingo boards or hell just a list 

// Oh dang that's easy in Go.
// Literally just creating a type of BingoBoard with a 2d array of int

// Evaluate takes a list of called numbers and gives you the index at which you achieve bingo
func (b BingoBoard) Evaluate(called []int) int {
	// Okay so this is where we have to go ahead and figure it out
	// So theoretically, we could have Evaluate return every single thing we need it to.
	// if the board wins, how many turns it took to win, the sum of the remaining numbers, etc.

	// Time to see if I can declared a number here...
	var lastnumbercalled int

	// okay so what we need to answer is what is a [5]int
	// okay so we've done the rows here
	for turns, callednumber := range called {
		// oh okay we need a range here, got it
		for _, row := range b {
			for _, number := range row {
				if number == callednumber {
					lastnumbercalled = number
					number = 'T'
					if b.CheckForBingos() == true {
						b.Score()
						


					}
					
					// This is where we call that function
					// Alternatively, don't check the first five times but check every time after 5

					// Do we actually assign it to number? Only one way to find out I guess
				}
			}

		}

	}

	// Now to do the columns

	// Now you need to iterate through and do the Evaluating
	// Ugh this is so damn inefficient
	// We really should just do this one at a time
	// Wait no this won't work at /all/ because the bingo could happen at any time
	// So we need a way to loop through columns, or we need to do this twice, or, oh.
	// We don't have to loop through at all for columns

	// I don't see a way to get through this without using a traditional for loop but that may not be a bad thing

	// There's some way to comprehend the columns here as well. I'm sure.
	// we could get the columns with the index I am absolutely sure of it
	// And if we don't, we're just going to be adding time complexity

	// Here we have to return the score and the turnstowin 
	
}

func (b BingoBoard) CheckForBingos() bool {
	// Both vertical and horizontal checking
	var bingo bool

	for _, row  := range b {
		// Could reset the TCount Here
		for _, number := range row {
			if number == 'T'{
				tcount += 1 
				if tcount == 5 {
					bingo = true
					return
					// have to hope it doesn't just return to outermost loop 
				}
			} 
		
	}

	for i, column := range b[i] {
		for _, number := range column {
			if number == 'T' {
				tcount += 1 
				if tcount == 5 {
					bingo = true 
					return
				}
			}
		}
		
	}

	return


}

func (b BingoBoard) Score(lastnumbercalled int) int  {
// The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

var sum int

	for _, row  := range b {
		for _, number := range row {
			if number != 'T' {
				sum += number
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
	boardmap := make(map[string]int)

	for _, n := range numbers {
		i, _ := strconv.Atoi(n)
		called = append(called, i)
		// well this is easy, we're just taking a file where the first line is the called #s and appending it to the called array/slice
	}

	for i := 2; i+5 < len(split); i += 6 {
		var b BingoBoard
		// no idea why we just have var b bingo board here
		//Is what I said when I last edited this file, now I know better. b is now the bingo board, you're welcome
		//But let's dive into what this does for the fun while we get ourselves up to speed
		// i := 2 we start 2 in in, then we said "while i + 5 < len(split)".
		//split being the current split we're processing
		// I think that we are taking the chunk of numbers up to the newline (the blank line after every bingo board)
		// Yes we absolutely are, and when we do that we declare a bingboard object

		for r, line := range split[i : i+5] {
			// so we're only use the first argument of for here, we're just telling it where to start and what to start on
			c := 0
			for pos := 0; pos < len(line); pos += 3 {
				// pos += 3 obviously moves us to the end of the line, so it looks like we're always hitting the first thing
				// Is what I thought a while ago
				// still not sure how we get the last number in the line here  OH
				// okay yes so we are actually going digit by digit here
				n, _ := strconv.Atoi(strings.Trim(line[pos:pos+2], " "))
				//this is where we get the actual number
				// Which is why we get the actual number here
				b[r][c] = n
				//okay so r and c must carry over, same with b
				//right so we are getting OH ITS THE ROW AND COLUMN, OH OH OH
				c++
				//OKAY THIS IS WHY WE IMPLEMENTED C HERE OKAY I GET THAT

			}
		}

		boards = append(boards, b)
		// okay why do we add this to boards?
		// OH okay get this, we're adding each to boards
		// Okay so this is how we parse each board. So eventually we will have a list of boards
	}

	for _, b := range boards {
		turnsToWin := b.Evaluate(called)
		fmt.Println(turnsToWin)
		//okay so what we need to print here as well are like
		// the score of the winning board (aka the sum of all unmarked numbers)
		// multiply that by the number that was just called when the board was won
		//and then find the answer
	}

}

//  chosen number is marked on all boards on which it appears.
// numbers may not appear on all boards
// no diagonols
// sum of all unmarked numbers * number that was just called = final score
// figure out which board wins first, what will be your final score if you choose that board
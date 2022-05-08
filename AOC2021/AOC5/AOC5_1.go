package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"gonum.org/v1/plot/vg"
)

// TODO: Make them all into one big list

func main() {

	pointslist, err := Parser()
	var noslopelist [][]float64

	for i, p := range *pointslist {
		s := make([]float64, 4)

		s[0], err = strconv.ParseFloat(p[0], 64)
		s[1], err = strconv.ParseFloat(p[1], 64)
		s[2], err = strconv.ParseFloat(p[2], 64)
		s[3], err = strconv.ParseFloat(p[3], 64)
		result := slopeChecker(s[0], s[1], s[2], s[3])
		if result == 0 {
			noslopelist = append(noslopelist, s)
			//QUESTION: How come I cannot do this?  *pointslist[i][i]
			// oh well, if I can't remove from the list I'll just make a new list
			// there has to be a way to edit it, perhaps call another function?
		}

	}

	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

	Parser()
}

// func addPoints(sx float64, sy float64, fx float64, fy float64) plotter.XYs {
// 	pts := make(plotter.XYs, 4)
// 	for i := range pts {
// 		if i == 0 || i == 1 {
// 			pts[i].X = sx
// 			pts[i].Y = sy
// 		} else {
// 			pts[i].X = fx
// 			pts[i].Y = fy
// 		}

// 	}
// 	return pts

// }

func slopeChecker(sx float64, sy float64, fx float64, fy float64) int {
	// Want to make sure that zero isn't what folks

	if fy-sy == 0 || fx-sx == 0 {
		return 0

	}

	// zero for horizontal
	// but undefined for y
	// So we might have to use a panic here

	return -1
}

func Parser() (*[][]string, error) {
	var newlist [][]string
	bytes, err := ioutil.ReadFile("puzzleinput.txt")
	if err != nil {
		return &newlist, err
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	for _, v := range split {
		s := make([]string, 4)
		entry := strings.Split(v, "->")
		print(entry[0])
		starting := strings.Split(entry[0], ",")
		finishing := strings.Split(entry[1], ",")
		s[0] = starting[0]
		s[1] = starting[1]
		s[2] = finishing[0]
		s[3] = finishing[1]

		newlist = append(newlist, s)

		// so a list of things like 60,28 -> 893,861 may be better represented as [][][]string
		// and then you do newlist = append(newlist, [][]string{starting, finishing})

		return &newlist, nil
	}

	return &newlist, nil

}

func LineEnumerator(sx float64, sy float64, fx float64, fy float64) {
	// Okay so what we need are 10 rows of numbers ending in 900
	// these go from 0 to 900
	// when there is nothing there we mark a ., otherwise increment number
	// I think the most efficient way would be to print according to each line after eliminating slopes
	// in fact hmm we should make an array of each
	// no, print them all out first
	// ugh that's so inefficent though
	// frick it though let's just do it for now
	// ugh in ruby you could just use cycle

	// okay so we need to sort the lists

	for i := 0; i < 1000; i++ {
		// could check here for every i but ugh that would take FOREVER unless we're usinga
		// GOT IT, we need to range over the ones already in it and just make everything else .s, duh
		if i%100 == 0 {
			// go that is not more legible why are you autosmushing it together
			print("\n")
		} else {
			print(".")
		}

	}

}
func LineSorter(noslopelist *[][]float64) {

	sorted := false

	for sorted != true {
		for i, v := range *noslopelist {
			if i != 0 {
				if v[0] < (*noslopelist)[i-1][0] {
					(*noslopelist)[i-1], (*noslopelist)[i] = (*noslopelist)[i], (*noslopelist)[i-1]
					// set to false here
					sorted = false
					continue
				}
				// this will only work if Go stops execution after the if statement
				// I feel like this wont work at all
				if i == len(*noslopelist)-1 {
					sorted = true
				}
			}

		}
		// okay so what we would need to know here is on what level does this modify it

	}
}

func LineCombiner(noslopelist *[][]float64) {
	combinedlist := []int{}
	for i, v := range *noslopelist {

	}

}

// GATED OFF FOR NOTES:

// Ugh, I've made a critical mistake. I should've perhaps combined the lists at the start but

// I don't think this is how we solve this either. At least not efficiently

// m = \frac{\text{rise}}{\text{run}} = \frac{y_2 - y_1}{x_2 - x_1}

// So what we need is something that checks for slopes and just, axes lines with slopes
// Then we need something that

// 0,9 -> 5,9
// 8,0 -> 0,8
// 9,4 -> 3,4
// 2,2 -> 2,1
// 7,0 -> 7,4
// 6,4 -> 2,0
// 0,9 -> 2,9
// 3,4 -> 1,4
// 0,0 -> 8,8
// 5,5 -> 8,2

// An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
// An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
// For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.
// AB: Oh, it literally tells you this

// So, the horizontal and vertical lines from the above list would produce the following diagram:

// .......1..
// ..1....1..
// ..1....1..
// .......1..
// .112111211
// ..........
// ..........
// ..........
// ..........
// 222111....

// In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9.
// AB: Okay so they did get the maximum and the minimum

// AB: What if I /did/ try to do the graph?
// This is 9 rows of 0 to 9
// Elisabet encourages me to view this as "what if we scaled up?"
// To answer that I need to ask "how did they get the minimum and maximum?"
// I would say the highest and lowest numbers in an x coordinate or y coordinate
// Elisabet points out that I could compute this
// We've agreed upon writin the parser

// To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.

// Consider only horizontal and vertical lines. At how many points do at least two lines overlap?

// So i could do
// linear algebra
// a very long calculation
// the problem here is that when you get into the hundreds of thousands we would have to keep slices that long
// could we boil it down by doing some sort of representation of 100 of thousands?
// they use 1s and 2s...
// The real issue here is we have a near unlimited set of numbers...
// how much is enough
// UNTIL THEYVE MET
// we don't care about negatives
// and we know there are maxs and mins here
//https://adventofcode.com/2021/day/5
// Each position is shown as the number of lines which cover that point
// or . if no line covers that point.
// so if we did this we could just tally them up at the end

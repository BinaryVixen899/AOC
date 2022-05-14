package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Notes:

func main() {
	var noslopelist [][]int

	// Default parsing
	pointslist, err := Parser()

	// Removing Slopes
	// WWID: Passing pointslist to to slopeChecker and therefore, rewriting slopechecker

	for _, v := range *pointslist {
		result := slopeChecker(v)

		if result == 0 {
			noslopelist = append(noslopelist, v)
		} else {
			continue

		}

	}

	// check results
	// Lines are sorted
	LineSorter(&noslopelist)
	// Lines are combined
	combinedlist := LineCombiner(&noslopelist)
	LineMaker(&combinedlist)

	LineEnumerator(&combinedlist)

	if err != nil {
		panic(err)
	}

}

func slopeChecker(pointslist [][]int) int {
	// Want to make sure that zero isn't what folks

	for i, starting := range pointslist {
		starting[0]
		starting[1]
		//WWID: have to iterate through the finishing values as well and this means no range. Maybe this is where I should start making classes

	}
	finishing := pointslist[0][1]

	if fy-sy == 0 || fx-sx == 0 {
		return 0

	}

	return -1
}

func Parser() (*[][][]int, error) {
	var newlist [][][]int
	bytes, err := ioutil.ReadFile("testpuzzleinput.txt")
	if err != nil {
		return &newlist, err
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	for _, v := range split {
		startingint := []int{}
		finishingint := []int{}
		entry := strings.Split(v, "->")
		starting := strings.Split(entry[0], ",")
		finishing := strings.Split(entry[1], ",")

		for _, v := range starting {
			v = strings.TrimSpace(v)
			v, _ := strconv.Atoi(v)
			startingint = append(startingint, v)
		}
		// Q: Why do I have to trim strings in order to not have strings with spaces become 0?

		for _, v := range finishing {
			v = strings.TrimSpace(v)
			v, _ := strconv.Atoi(v)
			finishingint = append(finishingint, v)
		}

		newlist = append(newlist, [][]int{startingint, finishingint})

		// CONSULT:
		// newlist = append(newlist, starting, finishing)
		// so a list of things like 60,28 -> 893,861 may be better represented as [][][]string
		// and then you do newlist = append(newlist, [][]string{starting, finishing})
	}

	return &newlist, nil

}

func LineMaker(combinedlist *[]int) (finallist []int) {
	finallist = make([]int, 1000)
	var outerindex int
	var count int
	var skipahead bool

	for i, v := range *combinedlist {

		// if skipahead == true && skipaheadvalue == i {
		// 	continue
		// 	// when skipaheadvalue no longer matches j will shortcircuit
		// }
		// split this off into its own function maybe?
		// Main function handling, we check if the index is equal to the value in the sorted list
		// // 7, 0

		if outerindex == v {
			count += 1
			// DO: HANDLE IF VALUE IS A REPEAT
			if (*combinedlist)[i+1] == v {
				if skipahead == true {
					continue
				} else {
					skipahead = true
				}

			} else {
				// DO: HANDLE PRINTING IF VALUE IS NOT A REPEAT
				skipahead = false
				finallist[outerindex] = count
				count = 0
				outerindex++
				continue

			}

			// okay so the issue is that we're checking if the index matches the CURRENT value not if the index matches every value in the list
			// so wee need to make sure it detects that 10 is a value in this first
			// okay so that works
			// repeatcount = 0
			// ALSO HAVE TO RESET SKIPAHEAD
			// reset the skipaheadvalue
			// for i := j; i+1 == v; i++ {
			// 	// does this get reset each time i increments or just set to j at the start?
			// 	// increment the count of repeats
			// 	repeatcount++

			// }
			// once it's over or if there were no repeats to begin with, print the count

			// We can set a variable here indicating that we need to skip ahead because we've cleared the hurdle
			// value to skip
			// skipaheadvalue = j

		} else {

			// okay so I think what we can do here is that if outerindex does not equal v, we iterate outerouterindexunouterindex != vt outerindex++s
			for outerindex < v {
				finallist[outerindex] = -1
				outerindex++
				// this might screw us up
			}
			if (*combinedlist)[i+1] == v {
				skipahead = true
				count += 1
			} else {
				skipahead = false
				finallist[outerindex] = count
				count = 0
				outerindex++
			}
			// TODO: When it does == v we need to hadnle that
			// okay so this should handle that
			// when i get back take care of the other cases
		}
	}

	return finallist
}

func LineEnumerator(combinedlist *[]int) {
	// var skipahead bool
	// var matched bool
	// skipaheadvalue := 1
	// // var repeatcount = 0
	// // Okay so what we need are 10 rows of numbers ending in 900
	// // these go from 0 to 900
	// // when there is nothing there we mark a ., otherwise increment number
	// // I think the most efficient way would be to print according to each line after eliminating slopes
	// // in fact hmm we should make an array of each
	// // no, print them all out first
	// // ugh that's so inefficent though
	// // frick it though let's just do it for now
	// // ugh in ruby you could just use cycle
	// // okay so we need to sort the lists

	// // Prolbem is that it's going to be over 1000 due to repeats
	// // Plan is to just do length,and then increment until the current i <= the last i
	// for i := 0; i < 1000; i++ {
	// 	// could check here for every i but ugh that would take FOREVER unless we're usinga
	// 	// GOT IT, we need to range over the ones already in it and just make everything else .s, duh

	// 	// We could go through everything within thi
	// 	// I IS ONLY ITERATED WHEN WE go through the WHOLE LIST

	// 	}
	// 	if i != 0 && i%100 == 0 {
	// 		print(".")
	// 		print("\n")
	// 		// move this below
	// 	} else {
	// 		if matched == true {
	// 			print()

	// 		} else {
	// 			print(".")
	// 		}
	// 		// okay so here is where we print it if we
	// 	}

	// }

}
func LineSorter(noslopelist *[][]int) {

	sorted := false

	// Sort the X's
	// TODO: DOes this just sort the Xs?

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

func FullLine(startingcoordinate int, endingcoordinate int) (line []int) {
	for i := int(startingcoordinate); i <= int(endingcoordinate); i++ {
		line = append(line, int(i))
	}
	if startingcoordinate == endingcoordinate {
		line = append(line, int(startingcoordinate))
	}
	return line
	// make this a pointer and such
}

func LineCombiner(noslopelist *[][]int) (combinedlist []int) {

	for _, v := range *noslopelist {
		xline := FullLine(v[0], v[2])
		yline := FullLine(v[1], v[3])
		print(xline)
		// okay so we're not getting the value if it's the same
		print(yline)
		combinedlist = append(combinedlist, xline...)
		combinedlist = append(combinedlist, yline...)
		// hopefully this will print out every single value
		// that is not what is happening
		// this only happens sometimes. ugh
		// So it doesn't include the values at the end
	}

	sort.Slice(combinedlist, func(i, j int) bool {
		return combinedlist[i] < combinedlist[j]
	})
	// Will this actually return it though?

	return combinedlist

	// combinedxlist := []int{}
	// combinedylist := []int{}

	// for _, v := range *noslopelist {
	// 	combinedxlist = append(combinedxlist, v[0], v[2])
	// }

	// for _, v := range *noslopelist {
	// 	combinedylist = append(combinedylist, v[1], v[3])
	// }

	// return &combinedxlist, &combinedylist
	// Okay so by the end of this everything should be modified.

}

// GATED OFF FOR NOTES:

// EXAMPLE PROBLEM:

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

// So, the horizontal and vertical lines from the above list would produce the following diagram:
// Each position is shown as the number of lines which cover that point or . if no line covers that point.

// 7, 0
// 7, 1

// 7, 3
// 7, 4

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

// To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.

// Consider only horizontal and vertical lines. At how many points do at least two lines overlap?

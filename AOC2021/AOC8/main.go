package main

// WWID: Fixing all the code so that we can run through inputs before moving onwards to an efficient way to do outputs

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type digit struct {
	letters            string
	uniqueletterscount int
	digittype          int
	digitcount         int
	// digit count, is the amount of times we've found this digit in the output
	// Is there subclass like functionality available here?
}

func newDigit(letters string, uniqueletterscount int, digittype int) *digit {
	d := digit{letters: letters, uniqueletterscount: uniqueletterscount, digittype: digittype}
	d.uniqueletterscount = uniqueletterscount
	d.digittype = digittype
	return &d
}

func newSpecialDigit(letters string, uniqueletterscount int, digitype int, digitcount int) *digit {
	d := digit{letters: letters, uniqueletterscount: uniqueletterscount, digittype: digitype, digitcount: digitcount}
	d.uniqueletterscount = uniqueletterscount
	d.digittype = digitype
	d.digitcount = digitcount
	return &d
}

type output struct {
	letters []string
}

// We might not need entry in the final version
type entry struct {
	digits [][]digit
	output []output
}

func main() {
	var uniquecount int
	var digittype int

	// Input parsing
	lines := parsetextfile()
	print(lines)

	// We will need these later when it comes to output
	// one := newSpecialDigit("one", 2, 0, 0)
	// four := newSpecialDigit("four", 4, 0, 0)
	// seven := newSpecialDigit("seven", 3, 0, 0)
	// eight := newSpecialDigit("eight", 8, 0, 0)

	// Okay so we're going to have to split on the spaces and such. Ugh.

	for _, v := range lines {
		// What if we split v[0]
		inputstring := strings.Split(v[0], " ")
		outputstring := strings.Split(v[1], " ")
		for _, word := range inputstring {
			uniquecount = getUniqueDigitsCount(word)
			digittype = identifyDigits(uniquecount)
			if digittype != -1 {
				newDigit(word, uniquecount, digittype)
				// Construct a temporary slice or dictionary
				print(word)
			}

			// The issue here is that word is a rune
			// So we have to split v[0] in order to do this

		}
		// Here is where we do the output string scanning

		// Okay so v[0] is the array with the problem
		// v[1] is the array with the output
		// Commenting these out for testing
		// uniquecount = getcountofuniquedigits(v)
		// digittype = v.identifydigits()
		// newDigit(v.string, uniquecount, digittype)
		print(v)
		// And then here we

	}

	// And then do here is where we print the counts when we're all done

	// for _, v := range lines {
	// 	// I feel like we shold use letter scanning here, some sort of regex instead of just arrays, because otherwise this thing is going ot be very inefficient

	// }

}

func identifyDigits(uniquecount int) int {
	// Is passed a uniquecount and returns what digit this is, returns -1 if unidentified
	// 1, 4, 7, 8
	switch uniquecount {
	case 2:
		return 1

	case 4:
		return 4

	case 3:
		return 7

	case 8:
		return 8

	default:
		return -1

	}

}

func getUniqueDigitsCount(letterchunk string) int {
	// Takes a chunk of letters and gives you back the count of unique letters
	var uniquecount int
	var uniqueslice []string

	letterchunkslice := strings.Split(letterchunk, "")
	sort.Strings(letterchunkslice)

	for i, v := range letterchunkslice {
		if i != len(letterchunkslice)-1 {
			if v != letterchunkslice[i+1] {
				uniqueslice = append(uniqueslice, v)
			}

		}
		// So at the end of this we should have a unique

	}

	uniquecount = len(uniqueslice)
	return uniquecount

}

// func identifyoutputdigits(letters string) {

// }

func parsetextfile() [][]string {
	var lines [][]string
	f, err := os.Open("testpuzzleinput.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		currentline := scanner.Text()
		splitline := strings.Split(currentline, "|")
		// okay so splitline 0 contains the problem and splitline 1 contains the output
		lines = append(lines, splitline)
	}

	return lines

}

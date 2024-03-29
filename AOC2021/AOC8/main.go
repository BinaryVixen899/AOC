package main

// WWID: And that the output is likewise sorted before going in

// For problem 2, we can almost build an 8 out of 1 4 and 7, once we do that we should be able to figure most things out
// for example, 8, but we can actually figure out a lot more things than just 8, for example, if something has six letters and the same letters as 1, it is a 9
// if it doesn't, it's a  6
// I'm tempted just to number these segments. But...
// I want to do the next problem, I've basically already solved this one

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type digit struct {
	letters            string
	sortedletters      []string
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

	// We will need these later when it comes to output
	one := newSpecialDigit("one", 2, 0, 0)
	four := newSpecialDigit("four", 4, 0, 0)
	seven := newSpecialDigit("seven", 3, 0, 0)
	eight := newSpecialDigit("eight", 7, 0, 0)

	// Okay so we're going to have to split on the spaces and such. Ugh.

	for _, v := range lines {
		// What if we split v[0]
		inputstring := strings.Split(v[0], " ")
		outputstring := strings.Split(v[1], " ")
		temporarymap := make(map[string]int)
		for _, word := range inputstring {
			uniquecount = 0
			uniquecount = getUniqueDigitsCount(word)
			digittype = identifyDigits(uniquecount)
			if digittype != -1 {
				somedigit := newDigit(word, uniquecount, digittype)
				somedigit.sortedletters = strings.Split(somedigit.letters, "")
				sort.Strings(somedigit.sortedletters)
				somedigit.letters = strings.Join(somedigit.sortedletters, "")
				temporarymap[somedigit.letters] = somedigit.digittype
				// Okay so what I THINK I've done here is that by saying *digit above I have given it a map of pointers which I will eventually need to dereference because ughhhhhhh
				// And I should ask later
				// Construct a temporary slice or dictionary
			}

			// The issue here is that word is a rune
			// So we have to split v[0] in order to do this

		}
		// Okay so now we're all done and we get to do the scanning

		for _, word := range outputstring {
			var wordslice []string
			// we can just do a case/switch statement here
			// The problem here is that we are mapping letters two things that are not organized
			// Also, we aren't identifying all numbers
			wordslice = strings.Split(word, "")
			sort.Strings(wordslice)
			word = strings.Join(wordslice, "")
			output := temporarymap[word]
			switch output {

			case 1:
				one.digitcount++
			case 4:
				four.digitcount++
			case 7:
				seven.digitcount++
			case 8:
				eight.digitcount++

			}

		}
		// Here is where we do the output string scanning

		// Okay so v[0] is the array with the problem
		// v[1] is the array with the output
		// Commenting these out for testing
		// uniquecount = getcountofuniquedigits(v)
		// digittype = v.identifydigits()
		// newDigit(v.string, uniquecount, digittype)
		// And then here we

	}

	// And then do here is where we print the counts when we're all done
	print("onecount:")
	print(one.digitcount)
	print("\n")
	print("fourcount")
	print(four.digitcount)
	print("\n")
	print("sevencount")
	print(seven.digitcount)
	print("\n")
	print("eightcount")
	print(eight.digitcount)

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

	case 7:
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
		} else {
			if v != letterchunkslice[i-1] {
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
	f, err := os.Open("puzzleinput.txt")

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

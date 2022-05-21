package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type digit struct {
	name               string
	uniqueletterscount int
}

type output struct {
	letters []string
}

type entry struct {
	digits [][]digit
	output []output
}

func main() {
	lines := parsetextfile()
	print(lines)
}

func identifyuniquecount(letters []string) {

}

func newDigit(name string, uniqueletterscount int) *digit {
	d := digit{name: name}
	d.uniqueletterscount = uniqueletterscount
	return &d
}

func identifydigits(uniquecount) {
	// Then do a switch/case statement here

}

func unique(letters []string) {
	mainletters := []string{"a", "b", "c", "d", "e", "f", "g"}
	var uniquecount int
	var uniqueslice []string

	for i, v := range letters {
		// if i != len - 1 basically
		if v == letters[i+1] {

		}

	}

}

func parsetextfile() [][]string {
	var lines [][]string
	f, err := os.Open("textpuzzleinput.txt")

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

package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type digit struct {
	letters            []string
	uniqueletterscount int
}

type output struct {
	letters []string
}

type problem struct {
	lines [][]string
}

func main() {
	lines := parsetextfile()
	print(lines)

}

func parsetextfile() []string {
	var lines []string
	f, err := os.Open("puzzleinput.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		currentline := scanner.Text()
		splitline := strings.Split(currentline, "|")
		// okay so splitline 0 contains the problem and splitline 1 contains the output
		lines = append(lines, splitline...)
	}

	return lines

}

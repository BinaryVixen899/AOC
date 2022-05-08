package main

import "fmt"

type lanternfish struct {
	timer int
	// Remember that the timer resets to 6
	firstfish    bool
	nlf          bool
	birthnextday bool
	// New lantern fish
}

// Notes:
// * LF.firstfish is there to track days
// Starting fish need to be manually intialized
// WWIUT: Putting in artifical days by creating a days struct or days global variable, seeing if after each tick fish can pass to the next fish (also, messaging Liz to ask if I should still do dictionaries when recursion makes so much more sense to me here )
// New lanternfish doesn't start counting down until the next day
// Also, realizing this is LITERALLY THE UNIQUE KEYS THING FROM LINKEDIN LEARNING

// Constructor, but bad in reality because there is ZERO way to force anyone to use it
// In reality you can just not export the type but export an itnerface
// WWID:

func NewLanternFish() *lanternfish {
	lf := lanternfish{}
	lf.timer = 8

	return &lf
}

func (lf *lanternfish) tick() {

	for lf.timer > 0 && lf.nlf != true {
		lf.timer -= 1
		return

	}

	if lf.timer == 0 && lf.birthnextday == false {
		// do a thing here to have fish birth the following day and restart 0
		lf.birthnextday = true
		return
	}

	if lf.nlf == true {
		lf.nlf = false
		return
	}

}

func main() {
	afish := &lanternfish{3, true, false, false}
	secondfish := &lanternfish{4, false, false, false}
	thirdfish := &lanternfish{3, false, false, false}
	fourthfish := &lanternfish{1, false, false, false}
	fifthfish := &lanternfish{2, false, false, false}
	var fishes = []*lanternfish{afish, secondfish, thirdfish, fourthfish, fifthfish}
	var day = 0
	// First we tick, then we check for -1 fish, if that's the case we call the birth fish method

	// and then we can do this up until we get the day we want
	for day <= 18 {
		if day == 0 {
			fmt.Println("The day is", day)
		}

		for _, fish := range fishes {
			print(fish.timer)
			print("\n")

		}

		for _, fish := range fishes {
			fish.tick()

		}

		for _, fish := range fishes {

			if fish.timer == 0 && fish.birthnextday == true {
				fishes = append(fishes, NewLanternFish())
				fish.timer = 6
				fish.birthnextday = false
				// this relies on the assumption that it will NOT count this new one. This may be a bad assumption
			}

		}

		day++
		fmt.Println("The day is", day)
		// Perhaps put whatever day it is here
	}
	print("test")

	// We can increase the date here

	// maybe do this until a certain number of days?
	// We would still need to find a way to have EVERYTHING execute ticks
}

// We need day 1 to be  2,3,2,0,1
// Currently day 1 is  2 3 2 6 1 8

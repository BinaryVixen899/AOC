package main

type lanternfish struct {
	timer int
	// Remember that the timer resets to 6
	firstfish bool
	nlf       bool
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

	for lf.timer >= 0 && lf.nlf != true {
		lf.timer -= 1
		print(lf.timer)
		print("\n")
		return

	}

	if lf.timer == -1 {
		print("test")
		lf.timer = 6
		return
	}

	if lf.nlf == true {
		lf.nlf = false
		return
	}

}

func main() {
	afish := &lanternfish{3, true, false}
	secondfish := &lanternfish{4, false, false}
	thirdfish := &lanternfish{3, false, false}
	fourthfish := &lanternfish{1, false, false}
	fifthfish := &lanternfish{2, false, false}
	var fishes = []*lanternfish{afish, secondfish, thirdfish, fourthfish, fifthfish}
	var day = 0
	// First we tick, then we check for -1 fish, if that's the case we call the birth fish method

	// and then we can do this up until we get the day we want
	for day < 18 {

		for _, fish := range fishes {
			fish.tick()
			if fish.timer == -1 {
				fishes = append(fishes, NewLanternFish())
				// this relies on the assumption that it will NOT count this new one. This may be a bad assumption
			}

		}
		print("A New Day Is Dawning \n")
		day++
		// Perhaps put whatever day it is here
	}
	print("test")

	// We can increase the date here

	// maybe do this until a certain number of days?
	// We would still need to find a way to have EVERYTHING execute ticks
}

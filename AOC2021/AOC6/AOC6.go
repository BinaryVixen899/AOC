package main

type lanternfish struct {
	timer int
	// Remember that the timer resets to 6
	startcountdown bool
	firstfish      bool
	nlf            bool
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

	for lf.timer >= 0 && lf.startcountdown == true && lf.nlf != true {
		if lf.firstfish == true {
			print("A day has passed \n")
		}
		lf.timer -= 1
		print(lf.timer)
		print("\n")
		return

	}

	if lf.timer == -1 {
		lf.timer = 6
		NewLanternFish().tick()
		if lf.firstfish == true {
			print("A day has passed ")
		}
	}

	if lf.nlf == true {
		lf.nlf = false
		return
	}

	if lf.startcountdown == false {
		lf.startcountdown = true
		lf.tick()
		// Recursive logic that should take into account the fake tick
	}

}

func main() {
	afish := lanternfish{3, true, true, false}
	secondfish := lanternfish{4, true, false, false}
	thirdfish := lanternfish{3, true, false, false}
	fourthfish := lanternfish{1, true, false, false}
	fifthfish := lanternfish{2, true, false, false}

	// maybe do this until a certain number of days?
	// We would still need to find a way to have EVERYTHING execute ticks
	afish.tick()
	secondfish.tick()
	thirdfish.tick()
	fourthfish.tick()
	fifthfish.tick()

}

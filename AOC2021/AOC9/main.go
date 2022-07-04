package main

import ic "github.com/WAY29/icecream-go/icecream"
import (
	"bytes"
	"os"
)

// Let's just put the notes here
// Okay so this is putting 576 out which is way too high, let's try off by one
// Okay that didn't work either, so my best guess is that this is something related to an edge case since the test data passes
// And I'm going to assume that's either at the edges or if a number above or below is the same
//5456789349886456890123985435578996543213456789656899996467789234989765442345789778999989652349879899
// I'm thinking the issue here is that our current algorithm will not take this and say "9 is a suspected lowpoint"
// If 5 was in the same position, it wouldn't do that either
// Theoretically we shouldn't need to. It should be caught by other stuff. But....Let's do it anyway

type basins struct {
	basinslice []*basin
}

type basin struct {
	heightmap          queue
	lowpoints          map[int]int
	suspectedlowpoints map[int]int
	nextbasin          *basin
	lastbasin          *basin
	islastbasin        bool
	isfirstbasin       bool
}

type queue struct {
	numbers []int
}

func (q *queue) DeQueue() {
	// Pointer reciever
	q.numbers = (q.numbers)[1:]
	// We can do this because of automatic dereferencing
}

func (b *basin) FindSuspectedLowPoints() {
	var lastnumber int
	b.suspectedlowpoints = make(map[int]int)
	// Initialize the map
	for i, v := range b.heightmap.numbers {
		// Checking for the first number
		if i == 0 {
			// if this is 0 we don't want to compare to nil
			// So instead let's compare to i+1
			if b.heightmap.numbers[i+1] > v {
				b.suspectedlowpoints[i] = b.heightmap.numbers[i]
			}
			continue
		}
		//Checking against the previous number
		lastnumber = b.heightmap.numbers[i-1]
		if i != len(b.heightmap.numbers)-1 && v < lastnumber && v < b.heightmap.numbers[i+1] {
			// Checking to make sure we are not at the end
			b.suspectedlowpoints[i] = b.heightmap.numbers[i]
		} else {
			// If we are at the end
			if i == len(b.heightmap.numbers)-1 && v < lastnumber {
				b.suspectedlowpoints[i] = b.heightmap.numbers[i]
				// A check solely for the last number
			}

		}
	}

	// Remove this in production
	for _, v := range b.suspectedlowpoints {
		ic.Ic("Suspected Lowpoints"+"\n", v)
	}
	ic.Ic("This is a new row")

}

func (b *basins) FindActualLowpoints() {

	for _, v := range b.basinslice {
		bsn := v
		for k, v := range bsn.suspectedlowpoints {
			// I have no idea if this is true
			if bsn.islastbasin == true {
				for k, v := range bsn.suspectedlowpoints {
					if bsn.lastbasin.heightmap.numbers[k] <= v {
						delete(bsn.suspectedlowpoints, k)
						// keep in suspected lowpoints, do nothing
					} else {
						continue
						//TODO: this also means that in the original findsuspectedlowpoints we will have to find a way to deal with lowpoints possibly already existing

					}

				}

			}

			if bsn.isfirstbasin == true {
				// We should be fine to skip this because we're handling it in the next one
				continue

			}

			// This is SUPPOSED to do the vertical checking
			// TODO: WE're getting to this with lastbasin and that's what's causing this
			if bsn.islastbasin == false {
				if bsn.nextbasin.heightmap.numbers[k] <= v || bsn.lastbasin.heightmap.numbers[k] <= v {
					delete(bsn.suspectedlowpoints, k)
					// keep in suspected lowpoints, do nothing
				} else {
					continue
					//TODO: this also means that in the original findsuspectedlowpoints we will have to find a way to deal with lowpoints possibly already existing

				}

			}

		}

		for _, v := range bsn.suspectedlowpoints {
			ic.Ic("actual Lowpoints"+"\n", v)
		}
		ic.Ic("This is a new row")

	}

}

func (b *basins) RiskCalculation() {
	var sum int
	for _, b := range b.basinslice {
		for _, v := range b.suspectedlowpoints {
			sum = sum + (v + 1)
		}

	}
	ic.Ic("Sum: %p", sum)

}

func (b *basins) CreateBasin(numbers []int) {
	somebasin := basin{}
	for _, v := range numbers {
		somebasin.heightmap.numbers = append(somebasin.heightmap.numbers, v)
		// then we need something that can intelligently comp through basin
		// I think we should be using a linkedlist here
	}
	b.basinslice = append(b.basinslice, &somebasin)

}

func (b *basins) LinkBasins() {
	// TD: Find a way to do this for the last basin
	//TD: So I think the issue here is that v is that basinslice is not a pointer
	// Yeah the more I think about this the more I think the issue is basinslice, becuase like we're able to modify v just fine
	// I still worry we've made copies though
	for i, v := range b.basinslice {
		if i+1 < len(b.basinslice) {
			test := b.basinslice[i+1]
			v.nextbasin = test
			// The reason this didn't work is it would set nextbasin to an address of a copy
			// if I had test = the address of &b.basinslice[i+1]
			b.basinslice[i+1].lastbasin = v
		} else {
			v.islastbasin = true
		}

	}
	b.basinslice[0].isfirstbasin = true
}

// Stolen from GobyExample
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (b *basins) CompBasins() {
	dat, err := os.ReadFile("./puzzleinput.txt")
	check(err)
	lines := bytes.Split(dat, []byte{'\n'})
	// clever trick for converting byte values to digits
	// bytes.split will need to get rid of the newlines or call trim, if it doesn't I will have to
	for _, line := range lines {
		intslice := make([]int, len(line))
		for i, b := range line {
			intslice[i] = int(b - '0')
		}
		b.CreateBasin(intslice)
		// create a new basin with intslice
	}

}

func main() {
	var basins basins
	ic.Enable()
	ic.Ic("test")
	basins.CompBasins()
	basins.LinkBasins()
	// Find Suspected Lowpoints

	for _, v := range basins.basinslice {
		v.FindSuspectedLowPoints()
	}
	basins.FindActualLowpoints()

	// do the risk calculation
	basins.RiskCalculation()

}

package main

import ic "github.com/WAY29/icecream-go/icecream"

type basins struct {
	basinslice []basin
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
		if i == 0 {
			// if this is 0 we don't want to compare to nil
			// So instead let's compare to i+1
			if b.heightmap.numbers[i+1] > v {
				b.suspectedlowpoints[i] = b.heightmap.numbers[i]
			}
			continue
		}
		lastnumber = b.heightmap.numbers[i-1]
		if i != len(b.heightmap.numbers)-1 && v < lastnumber && v < b.heightmap.numbers[i+1] {
			// Checking to make sure we are not at the end
			b.suspectedlowpoints[i] = b.heightmap.numbers[i]
		} else {
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
	ic.Ic("Sum:" + "\n")
	ic.Ic(sum)

}

func (b *basins) CreateBasin(numbers []int) {
	somebasin := basin{}
	for _, v := range numbers {
		somebasin.heightmap.numbers = append(somebasin.heightmap.numbers, v)
		// then we need something that can intelligently comp through basin
		// I think we should be using a linkedlist here
	}
	b.basinslice = append(b.basinslice, somebasin)

}

func (b *basins) LinkBasins() {
	// TD: Find a way to do this for the last basin
	//TD: So I think the issue here is that v is that basinslice is not a pointer
	// Yeah the more I think about this the more I think the issue is basinslice, becuase like we're able to modify v just fine
	// I still worry we've made copies though
	for i, _ := range b.basinslice {
		if i+1 < len(b.basinslice) {
			b.basinslice[i].nextbasin = &b.basinslice[i+1]
			// The reason this didn't work is it would set nextbasin to an address of a copy
			// if I had test = the address of &b.basinslice[i+1]
			b.basinslice[i+1].lastbasin = &b.basinslice[i]
		} else {
			b.basinslice[i].islastbasin = true
		}

	}
	b.basinslice[0].isfirstbasin = true
}

func (b *basins) CompBasins() {

}

func main() {
	var basins basins
	ic.Enable()
	ic.Ic("test")

	basins.CreateBasin([]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0})
	basins.CreateBasin([]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1})
	basins.CreateBasin([]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2})
	basins.CreateBasin([]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9})
	basins.CreateBasin([]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8})
	basins.LinkBasins()
	//TD: compute a list of heights depths
	// TD: Group these all into some sort of collection so we can just iterate through and call things
	// Find Suspected Lowpoints

	for _, v := range basins.basinslice {
		v.FindSuspectedLowPoints()
		//TD:Make "find actual lowpoints "
	}
	basins.FindActualLowpoints()

	// do the risk calculation
	basins.RiskCalculation()

}

package main

import ic "github.com/WAY29/icecream-go/icecream"

type basins struct {
	basinslice []basin
}

type basin struct {
	heightmap          queue
	lowpoints          map[int]int
	suspectedlowpoints map[int]int
	islastbasin        bool
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
	for _, v := range b.heightmap.numbers {

		ic.Ic("Suspected Lowpoints"+"\n", v)
	}

}
func (b *basin) FindActualLowpoints(nextbasin *basin, previousbasin *basin) {

	for k, v := range b.suspectedlowpoints {
		if nextbasin.heightmap.numbers[k] <= v || previousbasin.heightmap.numbers[k] <= v {
			delete(b.suspectedlowpoints, k)
			// keep in suspected lowpoints, do nothing
		} else {
			continue
			//TODO: this also means that in the original findsuspectedlowpoints we will have to find a way to deal with lowpoints possibly already existing

		}

	}
}

func (b *basin) FindActualLowpointsForLastBasin(previousbasin *basin) {

	for k, v := range b.suspectedlowpoints {
		if previousbasin.heightmap.numbers[k] <= v {
			delete(b.suspectedlowpoints, k)
			// keep in suspected lowpoints, do nothing
		} else {
			continue
			//TODO: this also means that in the original findsuspectedlowpoints we will have to find a way to deal with lowpoints possibly already existing

		}

	}
}

func (b *basin) printLowPoints() {
	for _, v := range b.lowpoints {
		print(v)

	}
	print("\n")
}

func (b *basin) printSuspectedLowPoints() {

	for _, v := range b.suspectedlowpoints {
		print(v)

	}
	print("\n")
}

func (b *basins) RiskCalculation() {
	var sum int
	for _, b := range b.basinslice {
		for _, v := range b.suspectedlowpoints {
			print(v + 1)
			sum = sum + (v + 1)
		}

	}
	print("\n")
	print("Sum:" + "\n")
	print(sum)

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

func main() {
	var basins basins
	ic.Enable()
	basins.CreateBasin([]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0})
	basins.CreateBasin([]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1})
	basins.CreateBasin([]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2})
	basins.CreateBasin([]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9})
	basins.CreateBasin([]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8})
	//TD: compute a list of heights depths
	// TD: Group these all into some sort of collection so we can just iterate through and call things
	// Find Suspected Lowpoints
	ic.Enable()

	for _, v := range basins.basinslice {
		v.FindSuspectedLowPoints()
		v.FindActualLowpoints()
		//TD:Make "find actual lowpoints "
	}

	//WWID: Print Suspected Lowpoints for Debugging Purposes (except inside of FindSuspectedLowPoints because why wouldn't I do that?)

	// do the risk calculation
	basins.RiskCalculation()

}

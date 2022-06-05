package main

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
	// takes a pointer
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
	for _, b := range b.basinslice {
		for _, v := range b.suspectedlowpoints {
			print(v + 1)
		}

	}

}

func main() {
	var a, b, c, d, e queue
	var ab, bb, cb, db, eb basin
	var basins basins

	a.numbers = []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}
	b.numbers = []int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1}
	c.numbers = []int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2}
	d.numbers = []int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9}
	e.numbers = []int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}
	// compute a list of heights depths
	ab.heightmap = a
	bb.heightmap = b
	cb.heightmap = c
	db.heightmap = d
	eb.heightmap = e
	eb.islastbasin = true
	// TD: Group these all into some sort of collection so we can just iterate through and call things
	// Find Suspected Lowpoints
	ab.FindSuspectedLowPoints()
	bb.FindSuspectedLowPoints()
	cb.FindSuspectedLowPoints()
	db.FindSuspectedLowPoints()
	eb.FindSuspectedLowPoints()
	// Print Suspected Lowpoints for Debugging Purposes
	print("Suspected Lowpoints" + "\n")
	ab.printSuspectedLowPoints()
	bb.printSuspectedLowPoints()
	cb.printSuspectedLowPoints()
	db.printSuspectedLowPoints()
	eb.printSuspectedLowPoints()

	// Find Actual LowPoints
	bb.FindActualLowpoints(&cb, &ab)
	cb.FindActualLowpoints(&db, &bb)
	db.FindActualLowpoints(&eb, &cb)
	eb.FindActualLowpointsForLastBasin(&db)

	// printLowPoints
	print("Actual Lowpoints" + "\n")
	ab.printSuspectedLowPoints()
	bb.printSuspectedLowPoints()
	cb.printSuspectedLowPoints()
	db.printSuspectedLowPoints()
	eb.printSuspectedLowPoints()
	basins.basinslice = append(basins.basinslice, ab, bb, cb, db, eb)
	// TD: Change this so we are no longer using a copy
	// do the risk calculation
	basins.RiskCalculation()

}

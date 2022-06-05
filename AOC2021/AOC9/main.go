package main

type basin struct {
	heightmap          queue
	lowpoints          map[int]int
	suspectedlowpoints map[int]int
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
			continue
		}
		lastnumber = b.heightmap.numbers[i-1]
		if v > lastnumber && i != len(b.heightmap.numbers) {
			// Checking to make sure we are not at the end
			b.suspectedlowpoints[i-1] = b.heightmap.numbers[i-1]
			// Add to suspected lowpoints if the current number is greater
		} else {
			if i == len(b.heightmap.numbers)-1 && v < lastnumber {
				b.suspectedlowpoints[i] = b.heightmap.numbers[i]
				// A check solely for the last number
			}

		}
	}

}
func (b *basin) FindActualLowpoints(nextbasin *basin) {
	// var nlp = nextbasin.suspectedlowpoints
	var nhmp = nextbasin.heightmap.numbers
	nextbasin.suspectedlowpoints = make(map[int]int)
	var nlp = nextbasin.suspectedlowpoints
	for k, v := range b.suspectedlowpoints {
		if nhmp[k] >= v {
			continue
			// keep in suspected lowpoints, do nothing
		} else {
			nlp[k] = v
			delete(nlp, k)
			//TODO: this also means that in the original findsuspectedlowpoints we will have to find a way to deal with lowpoints possibly already existing

		}

	}
}

func (b *basin) printLowPoints() {
	for _, v := range b.lowpoints {
		print(v)

	}
}

func (b *basin) printSuspectedLowPoints() {
	for _, v := range b.suspectedlowpoints {
		print(v)

	}
}

func (b *basin) RiskCalculation() {
}

func main() {
	var a, b, c, d, e queue
	var ab, bb, cb, db, eb basin
	a.numbers = []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}
	b.numbers = []int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1}
	c.numbers = []int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2}
	d.numbers = []int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9}
	e.numbers = []int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}
	ab.heightmap = a
	bb.heightmap = b
	cb.heightmap = c
	db.heightmap = d
	eb.heightmap = e
	// TD: Group these all into some sort of collection so we can just iterate through and call things
	ab.FindSuspectedLowPoints()

	// bb.FindSuspectedLowPoints()
	// cb.FindSuspectedLowPoints()
	// db.FindSuspectedLowPoints()
	// eb.FindSuspectedLowPoints()
	// ab.FindActualLowpoints(&bb)
	// bb.FindActualLowpoints(&cb)
	// cb.FindActualLowpoints(&db)
	// db.FindActualLowpoints(&eb)
	ab.printSuspectedLowPoints()
	// bb.printLowPoints()
	// cb.printLowPoints()
	// db.printLowPoints()
	// eb.printLowPoints()

	// compute a list of heights depths
	// Find suspected low points
	// find the actual low points
	// do the risk calculation

}

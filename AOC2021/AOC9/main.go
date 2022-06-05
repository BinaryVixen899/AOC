package main

type basin struct {
	heightmap          *queue
	lowpoints          *[]int
	suspectedlowpoints *map[int]int
	// don't need to have these be *
}

type queue struct {
	// I'm not sure if I'm doing this right but I want to get a pointer to this slice
	numbers *[]int
	// pointer to a slice
}

func (q *queue) DeQueue() {
	// takes a pointer
	p := (*q.numbers)[1:]
	q.numbers = &p
	// https://gosamples.dev/cannot-take-address/ I think that what I'm having to do is use a helper variable

	// q.numbers automatically dereferences, meaning I should be getting the values if I iterate
}

func (b *basin) FindSuspectedLowPoints() {
	var lastnumber int
	var heightmap = *b.heightmap.numbers
	var lp = *b.suspectedlowpoints
	// Set lastnumber to the initial number
	// premature optimization is the root of all evil

	for i, v := range heightmap {
		if i == 0 {
			lastnumber = v
			continue

		}
		lastnumber = heightmap[i-1]
		if v > lastnumber && i != len(heightmap) {
			lp[i-1] = heightmap[i-1]
		} else {
			if i == len(heightmap) && v < lastnumber {
				lp[i-1] = heightmap[i-1]
			}

		}
	}

}
func (b *basin) FindActualLowpoints(nextbasin *basin) {
	var nlp = *nextbasin.suspectedlowpoints
	var nhmp = *nextbasin.heightmap.numbers

	for k, v := range *b.suspectedlowpoints {
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

func (b *basin) RiskCalculation() {
}

func main() {
	var a, b, c, d, e queue
	a.numbers = &[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}
	b.numbers = &[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1}
	c.numbers = &[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2}
	d.numbers = &[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9}
	e.numbers = &[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}

	// Let's try to do this with a queue, which is a fifo structure

	// compute a list of heights depths
	// Find suspected low points
	// find the actual low points
	// do the risk calculation

}

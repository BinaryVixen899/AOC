package main

import (
	"sort"
	"strconv"
)

func main() {

	// I don't know how large my input will be, hence a slice
	inputslice := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	getmode(inputslice)
}

func getmode(myslice []int) {
	// Obviously this makes a copy, but if I had bugs in this code I wouldn't want it to be able to change the values of the slice. Plus this allows us to change the values of the int

	var greatestnumbervaluestring string
	var greatestnumbervalueint int
	var greatestnumberkey string

	if sort.IntsAreSorted(myslice) == false {
		sort.Ints(myslice)
	}

	numbermap := make(map[string]int)

	uniqueslice := unique(myslice)

	// screw it, we'll do multiple passes
	for _, v := range uniqueslice {
		h := strconv.Itoa(v)
		numbermap[h] = 1
	}

	for _, v := range myslice {
		h := strconv.Itoa((v))
		numbermap[h] += 1
	}

	for k, v := range numbermap {

		if greatestnumbervaluestring == "" {
			greatestnumbervalueint = v
			greatestnumberkey = k
			greatestnumbervaluestring = strconv.Itoa(greatestnumbervalueint)

		} else {

			if v > greatestnumbervalueint {
				greatestnumbervalueint = v
				greatestnumberkey = k
				greatestnumbervaluestring = strconv.Itoa(greatestnumbervalueint)

			}
		}

	}

	// print(greatestnumbervaluestring)
	print(greatestnumberkey)

	// Here is where we iterate over and increase the count

}

func unique(myslice []int) []int {
	// We have to pass a SORTED slice here or this will never work
	var uniqueslice []int

	for i, v := range myslice {
		if i != 0 {
			if myslice[i-1] != myslice[i] {
				uniqueslice = append(uniqueslice, v)

			}

		}

	}
	return uniqueslice

}

func fuelcalculation(myslice []int, mode int) {
	var fueltotal int
	for _, v := range myslice {
		fueltotal += v - mode

	}

	print(fueltotal)

}

package main

import "math"

// WWID: Making it actually run before I do anything else

// RULES
//First, we need to look around at all numbers around us and select the lowest one
// When we've selected the lowest one, we move there and mark both numbers as visited
// Anytime we have a decision point where two numbers are equal, we ignore it because we will eventually get to them on the next loop

// TODO:
// We need to make absolutely sure before we run this that we have a way to deal with the first and last rows, wher you literally cannot look up or down
// I'm not sure if this is a true depth first search. I'm not really popping things off :/
//ugh okay after looking at the definition this is probably not a true depth first search. Definitely room for improvement. But it will work.
// We have to do something with leastnumber after we find it
var visited []int
var answers []int

func main() {
	// var neighbors []int

	var currentnumber int
	// var currentnumberindex int

	combinednumbers := []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0, 3, 9, 8, 7, 8, 9, 4, 9, 2, 1, 9, 8, 5, 6, 7, 8, 9, 8, 9, 2, 8, 7, 6, 7, 8, 9, 6, 7, 8, 9, 9, 8, 9, 9, 9, 6, 5, 6, 7, 8}

	// a := []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}
	// b := []int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1}
	// c := []int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2}
	// d := []int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9}
	// e := []int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}
	for i := len(visited); i < len(combinednumbers); i++ {
		for i, _ := range combinednumbers {
			if contains(i, visited) {
				continue
			} else {
				if leftcornercheck(i) == true {
					leastnumber, solution := findleast(lookright(i, combinednumbers), lookdown(i, combinednumbers), lookup(i, combinednumbers))
					if solution != -1 {
						answers = append(answers, currentnumber)
						push(currentnumber, visited)
						continue
					}

				}
				if rightcornercheck(i) == true {
					leastnumber, solution := findleast(lookright(i, combinednumbers), lookdown(i, combinednumbers), lookup(i, combinednumbers))
					if solution != -1 {
						answers = append(answers, currentnumber)
						push(currentnumber, visited)
						continue
					}
				}
				if middlecheck(i) == true {
					leastnumber, solution := findleast(lookright(i, combinednumbers), lookdown(i, combinednumbers), lookup(i, combinednumbers))
					if solution != -1 {
						answers = append(answers, currentnumber)
						push(currentnumber, visited)
						continue
					}
				}

			}
		}
	}
}

// Navigation Steps
func lookright(index int, slice []int) int {

	if contains(index+1, slice) == false {
		push(index+1, visited)
		return slice[index+1]
	} else {
		return -1
	}
}

func lookleft(index int, slice []int) int {
	if contains(index-1, slice) == false {
		push(index-1, visited)
		return slice[index-1]
	} else {
		return -1
	}

}

func lookup(index int, slice []int) int {
	if index <= 9 {
		return -1
	}

	if contains(index-10, slice) == false {
		push(index-10, visited)
		return slice[index-10]
	} else {
		return -1
	}

}

func lookdown(index int, slice []int) int {
	if index >= 40 {
		return -1
	}

	if contains(index+10, slice) == false {
		push(index+10, visited)
		return slice[index+10]
	} else {
		return -1
	}

}

// CornerChecking
func leftcornercheck(index int) bool {
	if math.Mod(index, 10) == 0 {
		return true
	} else {
		return false
	}
}

func rightcornercheck(number int) bool {
	if math.Mod(index, 9) == 0 {
		return true
	} else {
		return false
	}
}

func middlecheck(number int) bool {
	if leftcornercheck && rightcornercheck == false {
		return true
	} else {
		return false
	}

}

// Misc Functions
func findleast(currentnumber int, numbers ...int) (leastnumber int, solution int) {
	// This will find the least but can only be fed non visited numbers
	// Also, this doesn't do anything with numbers that are the same... Maybe add them to a 'to visit' slice?
	//TODO: We need to remove the -1s as those are numbers that didn't count

	for _, number := range numbers {
		if leastnumber != 0 {
			if number < leastnumber {
				leastnumber = number
			}
		}
		if number < currentnumber {
			leastnumber = number
		}
	}

	if leastnumber != 0 {
		return leastnumber, -1
	} else {
		return -1, currentnumber
	}
}

func contains(index int, slice []int) bool {
	for i, v := range slice {
		if number == i {
			return true
		} else {
			return false
		}

	}

}

// Stack Methods
func push(number int, stack []int) {
	stack = append(stack, number)
}

func pop(number int, stack []int) int {
	var poppednumber int
	poppednumber = stack[len(stack)-1]
	remove(stack, poppednumber)
	return poppednumber

}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

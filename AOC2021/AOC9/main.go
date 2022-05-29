package main

// WWID: I didn't calculate visitednumbers well and that is what is messing this up

// RULES
//First, we need to look around at all numbers around us and select the lowest one
// When we've selected the lowest one, we move there and mark both numbers as visited
// Anytime we have a decision point where two numbers are equal, we ignore it because we will eventually get to them on the next loop

// TODO:
// We need to make absolutely sure before we run this that we have a way to deal with the first and last rows, wher you literally cannot look up or down
// I'm not sure if this is a true depth first search. I'm not really popping things off :/
//ugh okay after looking at the definition this is probably not a true depth first search. Definitely room for improvement. But it will work.
// We have to do something with leastnumber after we find it
// The contains function has to be modified so
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
					solutionnumber, solution := findleast(lookright(i, combinednumbers), lookdown(i, combinednumbers), lookup(i, combinednumbers))
					print(solutionnumber)
					if solution != false {
						answers = append(answers, currentnumber)
						visited = push(i, visited)
						continue
					} else {
						// this should be where we mark the current number as visited
						continue

					}

				}
				if rightcornercheck(i) == true {
					solutionnumber, solution := findleast(lookleft(i, combinednumbers), lookdown(i, combinednumbers), lookup(i, combinednumbers))
					print(solutionnumber)
					if solution != false {
						answers = append(answers, currentnumber)
						visited = push(i, visited)
						continue
					} else {
						push(i, visited)
						continue

					}
				}
				if middlecheck(i) == true {
					solutionnumber, solution := findleast(lookleft(i, combinednumbers), lookright(i, combinednumbers), lookdown(i, combinednumbers), lookup(i, combinednumbers))
					print(solutionnumber)
					if solution != false {
						answers = append(answers, currentnumber)
						visited = push(i, visited)
						continue
					} else {
						push(i, visited)
						continue

					}
				}

			}
		}
	}
	print(len(answers))
	for _, v := range answers {
		print(v)

	}
}

// Navigation Steps
func lookright(index int, slice []int) int {

	if contains(index+1, visited) == false {
		return slice[index+1]
		// THIS IS WHERE IT ERRORS OUT
	} else {
		return -1
	}
}

func lookleft(index int, slice []int) int {
	if contains(index-1, visited) == false {
		return slice[index-1]
	} else {
		return -1
	}

}

func lookup(index int, slice []int) int {
	if index <= 9 {
		return -1
	}

	if contains(index-10, visited) == false {
		return slice[index-10]
	} else {
		return -1
	}

}

func lookdown(index int, slice []int) int {
	if index >= 40 {
		return -1
	}

	if contains(index+10, visited) == false {
		return slice[index+10]
	} else {
		return -1
	}

}

// CornerChecking
func leftcornercheck(index int) bool {
	if index%10 == 0 {
		return true
	} else {
		return false
	}
}

func rightcornercheck(index int) bool {
	if index%10 == 9 {
		return true
	} else {
		return false
	}
}

func middlecheck(index int) bool {
	if (leftcornercheck(index) && rightcornercheck(index)) == false {
		return true
	} else {
		return false
	}

}

// Misc Functions
func findleast(currentnumber int, numbers ...int) (solution int, isasolution bool) {
	// This will find the least but can only be fed non visited numbers
	// Also, this doesn't do anything with numbers that are the same... Maybe add them to a 'to visit' slice?
	//TODO: We need to remove the -1s as those are numbers that didn't count
	leastnumber := 0

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

	if leastnumber != currentnumber {
		return solution, false
	} else {
		return solution, true
	}
}

func contains(number int, slice []int) bool {
	for _, v := range slice {
		if number == v {
			return true
		}

	}
	return false

}

// Stack Methods
func push(number int, stack []int) []int {
	stack = append(stack, number)
	return stack
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

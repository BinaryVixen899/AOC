package main

// WWID: Creating a stack

// We need some rules here. First, we need to look around at all numbers around us and select the lowest one
// When we've selected the lowest one, we move there and mark both numbers as visited
// anytime we have a decision point where two numbers are equal, we have to split off
// so the plan is to have each row be a, b, c, etc.
// Wait a second, we could totally do this horizontally. We just need to pass the next ones we're trying to the line below
// then we skip visited numbers when we restart
type branch struct {
	currentnumber int
}

func main() {
	var neighbors []int
	var visited []int
	var currentnumber int
	var currentnumberindex int

	a := []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}
	b := []int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1}
	c := []int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2}
	d := []int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9}
	e := []int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8}

// 	for i, v := range a {
// 		// only applicable in first row
// 		if somesortofcheckforfirstrow == true {
// 			// If we're at 0 or 9 in row a, set the number to that
// 			if i == 0 || 9 {
// 				currentnumber = a[i]
// 			}

// 			// We should check for visited first
// 			// Maybe a case statement to see if each is
// 			// Else, check up an down
// 			if a[i+1] < b[i] && a[i+1] {
// 				currentnumber = a[i+1]
// 			} else {
// 				currentnumber = b[i]
// 			}

// 		}

// 		push(currentnumber, visited)

// 	}

// }

// Ugh, I don't remember which way in Golang it's big o efficient to do this in when it comes to slices

//

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

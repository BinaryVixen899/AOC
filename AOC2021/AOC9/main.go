package main

// WWID: Creating a stack

// We need some rules here. First, we need to look around at all numbers around us and select the lowest one
// When we've selected the lowest one, we move there and mark both numbers as visited
// anytime we have a decision point where two numbers are equal, we have to split off

func main() {
	var neighbors []int
	var visited []int

}

// Ugh, I don't remember which way in Golang it's big o efficient to do this in when it comes to slices

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

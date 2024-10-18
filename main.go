package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var dLoop time.Duration
	var totalLoopDuration time.Duration
	added_nums := 0
	number_limit := 4000000
	iterations := 100

	start := time.Now()

	for j := 0; j < iterations; j++ {
		// Start timer and init array
		startLoop := time.Now()
		var fib []int

		// Initialize fibonacci sequence
		a, b := 0, 1

		// Compute fibonacci numbers
		for b < number_limit {
			fib = append(fib, b)
			a, b = b, a+b
		}

		added_nums = 0

		// Add even fibonacci numbers to variable added_nums
		for i := 1; i < len(fib); i++ {
			if fib[i]&1 != 1 {
				added_nums += fib[i]
			}
		}

		// Compute average loop duration
		totalLoopDuration += time.Since(startLoop)
		dLoop = totalLoopDuration / time.Duration(iterations)
	}

	// Print result
	addNums := strconv.Itoa(added_nums)
	numLim := strconv.Itoa(number_limit)
	fmt.Println("Even fibinacci numbers under " + numLim + ": " + addNums)
	fmt.Println("Average loop runtime:")
	fmt.Println(dLoop)
	fmt.Println("Total runtime for " + strconv.Itoa(iterations) + " iterations:")
	fmt.Println(time.Since(start))
}

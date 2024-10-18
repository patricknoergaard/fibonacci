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
		startLoop := time.Now()
		var fib []int

		a, b := 0, 1

		for b < number_limit {
			fib = append(fib, b)
			a, b = b, a+b
		}

		added_nums = 0

		for i := 1; i < len(fib); i++ {
			if fib[i]&1 != 1 {
				added_nums += fib[i]
			}
		}

		totalLoopDuration += time.Since(startLoop)
		dLoop = totalLoopDuration / time.Duration(iterations)
	}

	addNums := strconv.Itoa(added_nums)
	numLim := strconv.Itoa(number_limit)
	fmt.Println("Even fibinacci numbers under " + numLim + ": " + addNums)
	fmt.Println("Average loop runtime:")
	fmt.Println(dLoop)
	fmt.Println("Total runtime for " + strconv.Itoa(iterations) + " iterations:")
	fmt.Println(time.Since(start))
}

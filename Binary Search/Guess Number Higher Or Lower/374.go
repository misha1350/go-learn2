// Guess Number Higher or Lower
// https://leetcode.com/problems/guess-number-higher-or-lower/
package main

import "fmt"

func guessNumber(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := left + (right-left)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(guessNumber((6 * 7 * 2))) // should return 42
}

func guess(mid int) int {
	return 0 // placeholder
}

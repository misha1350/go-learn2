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
	fmt.Println(guessNumber(84)) // should return 42
}

func guess(mid int) int {
	return 0 // placeholder
}

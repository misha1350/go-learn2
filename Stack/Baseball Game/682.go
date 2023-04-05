package main

import (
	"fmt"
	"strconv"
)

func sum(nums []int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

func main() {
	operations := []string{"5", "2", "C", "D", "+"}
	record := []int{}
	for _, s := range operations {
		if len(record) > 0 && s == "C" {
			record = record[:len(record)-1]
		} else if len(record) > 0 && s == "D" {
			record = append(record, record[len(record)-1]*2)
		} else if len(record) > 0 && s == "+" {
			record = append(record, record[len(record)-1]+record[len(record)-2])
		} else {
			iop, _ := strconv.Atoi(s)
			record = append(record, iop)
		}
	}
	// return sum(record)
	fmt.Println(sum(record))
}

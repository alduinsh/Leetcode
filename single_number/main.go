package main

import "fmt"

func main() {
	result := singleNumber([]int{2, 2, 1, 4, 5, 4, 5})
	fmt.Println("Single number:", result)
}

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

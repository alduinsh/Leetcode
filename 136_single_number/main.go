/*
136. Single Number
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

package main

import "fmt"

/*var 1
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
*/

//var 2

func main() {
	//инициализируем срез целых чисел
	nums := []int{2, 3, 3, 4, 4, 10, 5, 5}
	result := singleNumber(nums)
	fmt.Println(result)
}

func singleNumber(nums []int) int {
	uniq := make(map[int]int)
	for _, n := range nums {
		uniq[n]++
	}
	for num, count := range uniq {
		if count == 1 {
			return num
		}
	}
	return 0
}

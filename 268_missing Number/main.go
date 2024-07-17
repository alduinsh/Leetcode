package main

import "fmt"

func main() {
	nums1 := []int{0, 1}
	fmt.Println("Недостающее число:", missingNumber(nums1))
}

func missingNumber(nums []int) int {
	numMap := make(map[int]bool)
	n := len(nums)

	//заполнение карты
	for _, num := range nums {
		numMap[num] = true
	}
	// ищем число
	for i := 0; i <= n; i++ {
		if !numMap[i] {
			return i
		}
	}
	return -1
}

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 11, 56, 1}
	fmt.Println(containsDuplicate(nums))
}

// проверяет, есть ли в массиве дубликаты
func containsDuplicate(nums []int) bool {
	//создаем мапу для отслеживания, пустую
	boolTrue := make(map[int]bool)

	//цикл для прохождения по массиву
	for _, n := range nums {
		//если ранее встречался элемент, то возвращаем true
		if boolTrue[n] {
			return true
		}
		// если ранее его не было то просто запишем его в нашу мапу
		boolTrue[n] = true
	}

	//если не встретили в цикле дубликаты, то просто false
	return false
}

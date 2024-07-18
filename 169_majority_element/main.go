/*
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2
*/

package main

import "fmt"

func main() {
	s := []uint8{1, 3, 2, 1, 3, 4, 4, 0, 10, 0}
	fmt.Println(majorityElement(s))

}

func majorityElement(nums []uint8) uint8 {
	countMap := make(map[uint8]uint8, 10) //создадим мапу для подсчета вхождений

	//цикл для записи кол-ва вхождений в срез nums каждого числа
	for _, number := range nums {
		countMap[number]++ //при встрече увеличиваем значение на 1 для найденного ключа
	}

	//ищем наибольшее количество вхождений
	//создаем без инициализации две пер
	var maxCount uint8
	var frequentNum uint8

	//пройдемся циклом по полученной мапе
	for number, count := range countMap {
		if count > maxCount {
			maxCount = count     //перезапишем в макскоунт число если оно больше
			frequentNum = number //перезаписываем number на каждой итерации
		}

	}
	return frequentNum
}

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.


Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
*/

package main

import "fmt"

func main() {
	fmt.Println(isAnagram("listen", "silent"))
	fmt.Println(isAnagram("anagram", "silent"))
}

func isAnagram(s string, t string) bool {
	//сравним дилну строк, если не равны, то сразу false
	if len(s) != len(s) {
		return false
	}

	//создадим хэш-табл из ключ=руна, значение=числов вхождений
	count := make(map[rune]int)

	//посчитаем вхождения первой строки
	for _, i := range s {
		count[i]++
	}
	//пройдем по строке t с уменьшением счетчика и вернем false если будет значение меньше нуля
	for _, i := range t {
		count[i]--
		if count[i] < 0 {
			return false
		}
	}

	return true
}

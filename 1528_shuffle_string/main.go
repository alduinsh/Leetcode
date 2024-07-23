/*
You are given a string s and an integer array indices of the same length. The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

Return the shuffled string.

*/

package main

import "fmt"

func main() {
	s := "cba"
	indices := []int{0, 1, 2}

	result := restoreString(s, indices)
	fmt.Println(result)
}

func restoreString(s string, indices []int) string {
	//создаем срез рун длиной равной s
	shuff := make([]rune, len(s))

	//переберем в цикле
	for i, char := range s {
		shuff[indices[i]] = char
	}
	//преобразуем в строку
	return string(shuff)
}

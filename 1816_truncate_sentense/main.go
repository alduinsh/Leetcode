/*
A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each of the words consists of only uppercase and lowercase English letters (no punctuation).

For example, "Hello World", "HELLO", and "hello world hello world" are all sentences.
You are given a sentence s​​​​​​ and an integer k​​​​​​. You want to truncate s​​​​​​ such that it contains only the first k​​​​​​ words. Return s​​​​​​ after truncating it.



Example 1:

Input: s = "Hello how are you Contestant", k = 4
Output: "Hello how are you"
Explanation:
The words in s are ["Hello", "how" "are", "you", "Contestant"].
The first 4 words are ["Hello", "how", "are", "you"].
Hence, you should return "Hello how are you".

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello how are you Contestant"
	k := 4
	fmt.Println(truncateSentence(s, k))
}

func truncateSentence(s string, k int) string {
	//разбиваем строку
	words := strings.Split(s, " ")

	if k > len(words) {
		k = len(words)
	}

	//соберем строку до числа к
	return strings.Join(words[:k], " ")
}

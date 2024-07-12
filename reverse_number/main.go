//344. Reverse String
//Write a function that reverses a string.
//The input string is given as an array of characters s.

//var 1
/*
package main

import "fmt"

var name = []byte("Людмила")

func main() {
	reverseString(name)
	fmt.Println(string(name))
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
*/

// var2 for unicode and >1byte
package main

import "fmt"

var name = []rune("Людмила")

func main() {
	reverseString(name)
	fmt.Println(string(name))
}

func reverseString(s []rune) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

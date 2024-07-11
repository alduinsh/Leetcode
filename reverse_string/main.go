package main

import "fmt"

var name = []byte("Людмила")

func main() {
	reverseString(name)
	fmt.Println(string(name))
}

func reverseString(s []byte) {
	left, right := 0, len(name)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

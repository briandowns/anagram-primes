package main

import "fmt"

const maxSize = 26

var primes = []int{
	2, 3, 5, 7, 11, 13, 17, 19,
	23, 29, 31, 37, 41, 43, 47,
	53, 59, 61, 67, 71, 73, 79,
	83, 89, 97, 101,
}

var letters = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u',
	'v', 'w', 'x', 'y', 'z',
}

// uniqueNum generates a unique value assinged to the
// word passed in based on the letters it contains.
func uniqueNum(w1 string) int {
	prod := 1
	for _, l := range w1 {
		for j := 0; j < maxSize-1; j++ {
			if l == letters[j] {
				prod *= primes[j]
			}
		}
	}
	return prod
}

// isAnagram determines whether the given words are
// anagrams of one another.
func isAnagram(w1, w2 string) bool {
	return uniqueNum(w1) == uniqueNum(w2)
}

func main() {
	if isAnagram("ear", "are") {
		fmt.Println("true")
	}
}

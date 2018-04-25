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

// uniqueNum2
func uniqueNum2(w1 string) int {
	prod := 1
	for _, l := range w1 {
		switch l {
		case 'a':
			prod *= 2
		case 'b':
			prod *= 3
		case 'c':
			prod *= 5
		case 'd':
			prod *= 7
		case 'e':
			prod *= 11
		case 'f':
			prod *= 13
		case 'g':
			prod *= 17
		case 'h':
			prod *= 19
		case 'i':
			prod *= 23
		case 'j':
			prod *= 29
		case 'k':
			prod *= 31
		case 'l':
			prod *= 37
		case 'm':
			prod *= 41
		case 'n':
			prod *= 43
		case 'o':
			prod *= 47
		case 'p':
			prod *= 53
		case 'q':
			prod *= 59
		case 'r':
			prod *= 61
		case 's':
			prod *= 67
		case 't':
			prod *= 71
		case 'u':
			prod *= 73
		case 'v':
			prod *= 79
		case 'w':
			prod *= 83
		case 'x':
			prod *= 89
		case 'y':
			prod *= 97
		case 'z':
			prod *= 101
		}
	}
	return prod
}

// uniqueNum2
func uniqueNum3(w1 string) int {
	prod := 1
	for _, l := range w1 {
		switch l {
		case 'e':
			prod *= 11
		case 't':
			prod *= 71
		case 'a':
			prod *= 2
		case 'o':
			prod *= 47
		case 'i':
			prod *= 23
		case 'n':
			prod *= 43
		case 's':
			prod *= 67
		case 'h':
			prod *= 19
		case 'r':
			prod *= 61
		case 'd':
			prod *= 7
		case 'l':
			prod *= 37
		case 'c':
			prod *= 5
		case 'u':
			prod *= 73
		case 'm':
			prod *= 41
		case 'w':
			prod *= 83
		case 'f':
			prod *= 13
		case 'g':
			prod *= 17
		case 'y':
			prod *= 97
		case 'p':
			prod *= 53
		case 'b':
			prod *= 3
		case 'v':
			prod *= 79
		case 'k':
			prod *= 31
		case 'j':
			prod *= 29
		case 'x':
			prod *= 89
		case 'q':
			prod *= 59
		case 'z':
			prod *= 101
		}
	}
	return prod
}

// isAnagram determines whether the given words are
// anagrams of one another.
func isAnagram(w1, w2 string) bool {
	return uniqueNum(w1) == uniqueNum(w2)
}

func isAnagram2(w1, w2 string) bool {
	return uniqueNum2(w1) == uniqueNum2(w2)
}

func isAnagram3(w1, w2 string) bool {
	return uniqueNum3(w1) == uniqueNum3(w2)
}

func main() {
	if isAnagram("ear", "are") {
		fmt.Println("true")
	}
}

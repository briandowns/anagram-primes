package main

import "testing"

type testCase struct {
	word1    string
	word2    string
	expected bool
}

func TestIsAnagram(t *testing.T) {
	tc := []testCase{
		{
			word1:    "dog",
			word2:    "god",
			expected: true,
		},
		{
			word1:    "og",
			word2:    "god",
			expected: false,
		},
		{
			word1:    "thing",
			word2:    "night",
			expected: true,
		},
		{
			word1:    "silent",
			word2:    "listen",
			expected: true,
		},
		{
			word1:    "silet",
			word2:    "listen",
			expected: false,
		},
	}
	for _, i := range tc {
		if res := isAnagram(i.word1, i.word2); res != i.expected {
			t.Errorf("expected %t but got %t\n", i.expected, res)
		}
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram("silent", "listen")
	}
}

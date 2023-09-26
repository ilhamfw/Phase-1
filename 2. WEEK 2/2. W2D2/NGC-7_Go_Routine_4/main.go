package main

import (
	"fmt"
	"sync"
)

var letterValues = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

func main() {
	input := []string{
		"exampleword",
		"scrabble",
		"go",
	}

	var wg sync.WaitGroup

	for _, word := range input {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			score := calculateScrabbleScore(w)
			fmt.Printf("%s | scrabble score %d\n", w, score)
		}(word)
	}

	wg.Wait()
}

func calculateScrabbleScore(word string) int {
	score := 0
	for _, letter := range word {
		letter = toUpper(letter)
		value, found := letterValues[letter]
		if found {
			score += value
		}
	}
	return score
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 'a' + 'A'
	}
	return r
}

package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	input := []string{
		"As Soon As Possible",
		"Liquid-crystal display",
		"Thank George It's Friday!",
	}

	var wg sync.WaitGroup

	for _, word := range input {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			acronym := generateAcronym(w)
			fmt.Printf("%s - %s\n", w, acronym)
		}(word)
	}

	wg.Wait()
}

func generateAcronym(input string) string {
	words := strings.Fields(input)
	acronym := ""

	for _, word := range words {
		// Remove punctuation and convert to uppercase
		normalizedWord := strings.ToUpper(strings.TrimFunc(word, func(r rune) bool {
			return !isLetter(r)
		}))

		if normalizedWord != "" {
			acronym += string(normalizedWord[0])
		}
	}

	return acronym
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

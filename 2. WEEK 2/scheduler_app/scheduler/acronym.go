package scheduler

import (
	"fmt"
	"strings"
	"sync"
)

// AcronymTask adalah fungsi untuk menghasilkan akronim yang akan dijadwalkan.
func AcronymTask(wg *sync.WaitGroup, input string) {
	defer wg.Done()
	acronym := generateAcronym(input)
	fmt.Printf("%s - %s\n", input, acronym)
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

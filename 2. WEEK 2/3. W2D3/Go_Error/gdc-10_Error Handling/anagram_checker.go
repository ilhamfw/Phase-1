package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	if len(os.Args) != 3 {
		panic("Harap masukkan dua kata sebagai argumen.")
	}

	word1 := os.Args[1]
	word2 := os.Args[2]

	if !isValidInput(word1) || !isValidInput(word2) {
		panic("Input tidak valid")
	}

	if areAnagrams(word1, word2) {
		fmt.Printf("[%s] & [%s] merupakan anagram\n", word1, word2)
	} else {
		fmt.Printf("[%s] & [%s] bukan merupakan anagram\n", word1, word2)
	}
}

func isValidInput(word string) bool {
	if len(word) > 10 {
		return false
	}
	for _, char := range word {
		if !isValidChar(char) {
			return false
		}
	}
	return true
}

func isValidChar(char rune) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
}

func areAnagrams(word1, word2 string) bool {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	if len(word1) != len(word2) {
		return false
	}

	counts := make(map[rune]int)

	for _, char := range word1 {
		counts[char]++
	}

	for _, char := range word2 {
		if counts[char] == 0 {
			return false
		}
		counts[char]--
	}

	return true
}

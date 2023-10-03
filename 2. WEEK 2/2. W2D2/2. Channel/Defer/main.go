package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	// Open File
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error Opening file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	wordCount := 0

	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		words := strings.Fields(line) // Perbaiki strings.Fields(line)
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of lines: %d\n", lineCount)
	fmt.Printf("Number of words: %d\n", wordCount)
}

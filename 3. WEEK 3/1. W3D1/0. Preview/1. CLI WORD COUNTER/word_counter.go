package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run word_counter.go <filename>")
		return
	}

	filename := os.Args[1]

	// Read the input text file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			wg.Add(1)
			go func(w string) {
				defer wg.Done()
				mutex.Lock()
				wordCounts[w]++
				mutex.Unlock()
			}(word)
		}
	}

	wg.Wait()

	fmt.Println("Word Count:")
	for word, count := range wordCounts {
		fmt.Printf("%s : %d\n", word, count)
	}
}

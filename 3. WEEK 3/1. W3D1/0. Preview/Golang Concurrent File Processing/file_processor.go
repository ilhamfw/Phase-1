package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run file_processor.go input.csv output.csv")
		return
	}

	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	// Open the input CSV file
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inputFile.Close()

	// Create the output CSV file
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Initialize CSV readers and writers
	inputReader := csv.NewReader(inputFile)
	outputWriter := csv.NewWriter(outputFile)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	// Read and process CSV records concurrently
	for {
		record, err := inputReader.Read()
		if err != nil {
			break
		}

		wg.Add(1)
		go func(record []string) {
			defer wg.Done()

			// Process the record (in this case, capitalize name and add "Mr." prefix)
			record[0] = strings.ToUpper(record[0])
			record[2] = "Mr." + record[2]

			// Write the processed record to the output CSV file
			mutex.Lock()
			if err := outputWriter.Write(record); err != nil {
				fmt.Printf("Error writing to output file: %v\n", err)
			}
			mutex.Unlock()
		}(record)
	}

	wg.Wait()
	outputWriter.Flush()

	if err := outputWriter.Error(); err != nil {
		fmt.Printf("Error flushing output writer: %v\n", err)
	}
}

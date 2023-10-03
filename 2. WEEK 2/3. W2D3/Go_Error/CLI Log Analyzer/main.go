package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// LogAnalyzer represents the log analysis tool
type LogAnalyzer struct {
	LogFile string
}

// NewLogAnalyzer creates a new instance of LogAnalyzer
func NewLogAnalyzer(logFile string) *LogAnalyzer {
	return &LogAnalyzer{LogFile: logFile}
}

// AnalyzeLogs reads and analyzes log files
func (la *LogAnalyzer) AnalyzeLogs() {
	// Open the log file
	file, err := os.Open(la.LogFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Initialize a map to store log level counts
	logLevels := make(map[string]int)

	// Read and analyze each log line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		logLevel := extractLogLevel(line)
		logLevels[logLevel]++
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Display log analysis results
	fmt.Println("Log Analysis Results:")
	for level, count := range logLevels {
		fmt.Printf("[%s] Level: %d occurrences\n", level, count)
	}
}

// extractLogLevel extracts the log level from the log line
func extractLogLevel(line string) string {
	// Assume the log level is enclosed in square brackets
	start := strings.Index(line, "[")
	end := strings.Index(line, "]")
	if start != -1 && end != -1 && end > start {
		return line[start+1 : end]
	}
	return "UNKNOWN"
}

func main() {
	// Parse command-line flags
	logFile := flag.String("logfile", "", "Path to the log file")
	flag.Parse()

	if *logFile == "" {
		fmt.Println("Usage: loganalyzer -logfile <path_to_log_file>")
		return
	}

	// Create a new log analyzer
	la := NewLogAnalyzer(*logFile)

	// Handle panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred:", r)
		}
	}()

	// Analyze logs and handle errors
	la.AnalyzeLogs()

	// Successful execution
}

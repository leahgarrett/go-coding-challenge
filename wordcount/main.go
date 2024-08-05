package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// This version was implemented by Copilot

	// sample main version runs:
	// "moby.txt": 181239 words, duration: 585907ms
	// "moby.txt": 181239 words, duration: 581428ms
	// "moby.txt": 181239 words, duration: 585015ms

	// sample run on this branch
	// "moby.txt": 215838 words, duration: 17211ms
	// "moby.txt": 215838 words, duration: 12804ms
	// "moby.txt": 215838 words, duration: 11604ms

	// Read the file
	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	start := time.Now()

	// Convert the content to string
	text := string(content)

	// Count the words
	wordCount := countWords(text)

	fmt.Printf("%q: %d words, duration: %dms\n", os.Args[1], wordCount, time.Since(start)/1000)
}

func countWords(text string) int {
	// Split the text into words
	words := strings.Fields(text)

	// Return the count of words
	return len(words)
}

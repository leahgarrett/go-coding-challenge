package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

func main() {
	// This version was implemented by Copilot

	// Read the file
	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fc, err := os.Create("cpuprofile.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(fc)
	defer pprof.StopCPUProfile()

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

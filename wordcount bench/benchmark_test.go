package main

import (
	"log"
	"os"
	"testing"
)

func BenchmarkCountWords(b *testing.B) {
	// Prepare the input data
	bytes, _ := ReadFile("test_data.txt")
	countWords(bytes)
}

func BenchmarkByteCountWordsV2(b *testing.B) {
	f, err := os.Open("test_data.txt")
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}
	byteCountWordsV2(f)
}

func BenchmarkByteCountWordsV3(b *testing.B) {
	// Prepare the input data
	bytes, _ := ReadFile("test_data.txt")
	byteCountWordsV3(bytes)
}

func BenchmarkFastWordCount(b *testing.B) {
	// Prepare the input data
	f, err := os.Open("test_data.txt")
	if err != nil {
		log.Fatalf("could not open file %q: %v", "test_data.txt", err)
	}
	defer f.Close()

	for i := 0; i < b.N; i++ {
		// Reset the file pointer to the beginning for each iteration
		f.Seek(0, 0)
		fastWordCount(f)
	}
}

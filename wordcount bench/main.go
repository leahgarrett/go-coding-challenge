package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	filePath := os.Args[1]
	bytes, _ := ReadFile(filePath)
	fmt.Printf("ReadFile: %q %d bytes\n", os.Args[1], len(bytes))
	var start = time.Now()
	var wordCount = countWords(bytes)
	fmt.Printf("countWords: %d words, duration: %dns\n", wordCount, time.Since(start)/1000)

	start = time.Now()
	wordCount, err := fastestCountWords(filePath)
	if err != nil {
		fmt.Printf("%d", err)
	}
	fmt.Printf("fastestCounWords: %d words, duration: %dns\n", wordCount, time.Since(start)/1000)

	start = time.Now()
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}
	wordCount = byteCountWords(f)
	fmt.Printf("byteCountWords: %d words, duration: %dns\n", wordCount, time.Since(start)/1000)

	wordCount = byteCountWordsV2(f)
	fmt.Printf("byteCountWordsV2: %d words, duration: %dns\n", wordCount, time.Since(start)/1000)

	start = time.Now()
	wordCount = byteCountWordsV3(bytes)
	fmt.Printf("byteCountWordsV3: %d words, duration: %dns\n", wordCount, time.Since(start)/1000)

}

func ReadFile(filePath string) ([]byte, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Could not open file %q: %v", os.Args[1], err)
		return nil, err
	}
	return bytes, nil
}

// Producing the most generalizable and fast word count
func fastestCountWords(filePath string) (int, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("could not open file %q: %v", filePath, err)
	}
	defer fd.Close()

	words := 0
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file %q: %v", filePath, err)
	}

	return words, nil
}

func countWords(bytes []byte) int {
	text := string(bytes)
	words := strings.Fields(text)

	return len(words)
}

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func byteCountWordsV3(bytes []byte) int {
	words := 0

	inword := false
	for i := 0; i < len(bytes); i++ {
		r := rune(bytes[i])

		if unicode.IsSpace(r) {
			if inword {
				words++
				inword = false
			}
		} else {
			inword = true
		}
	}

	return words
}

func byteCountWordsV2(f *os.File) int {
	words := 0
	inword := false
	b := bufio.NewReader(f)
	for {
		r, err := readbyte(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}

	return words
}

func byteCountWords(f *os.File) int {
	words := 0
	inword := false
	for {
		r, err := readbyte(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}

	return words
}

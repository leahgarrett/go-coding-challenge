package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"time"
	"unicode"
)

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

// before adding buffering
// "shakespeare.txt": 741200 words, duration: 2673868ms

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}

	fc, err := os.Create("cpuprofile.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(fc)
	defer pprof.StopCPUProfile()

	start := time.Now()
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
	fmt.Printf("%q: %d words, duration: %dms\n", os.Args[1], words, time.Since(start)/1000)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"time"
	"unicode"
)

var buf [1]byte

func readbyte(r io.Reader) (rune, error) {

	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

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
	b := bufio.NewReader(f)
	inword := false
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
	fmt.Printf("%q: %d words, duration: %dms\n", os.Args[1], words, time.Since(start)/1000)
}

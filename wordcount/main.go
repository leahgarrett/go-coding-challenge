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

func main() {
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	log.Fatalf("could not open file %q: %v", os.Args[1], err)
	// }

	fc, err := os.Create("cpuprofile.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(fc)
	defer pprof.StopCPUProfile()

	s, _ := os.ReadFile(os.Args[1])

	start := time.Now()
	words := 0
	// b := bufio.NewReader(f)

	inword := false
	for i := 0; i < len(s); i++ {
		r := rune(s[i])

		if unicode.IsSpace(r) {
			if inword {
				words++
				inword = false
			}
		} else {
			inword = true
		}
	}
	fmt.Printf("%q: %d words, duration: %dms\n", os.Args[1], words, time.Since(start)/1000)
}

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	log.SetFlags(0)

	byFlag := flag.String("by", "line", "line, byte, or rune")
	flag.Parse()

	switch *byFlag {
	case "line":
		m, err := LineFreq(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		for _, k := range keysSortedByDescValue(m) {
			fmt.Printf("%d\t%s\n", m[k], k)
		}
	case "byte":
		m, err := ByteFreq(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		for _, k := range keysSortedByDescValue(m) {
			fmt.Printf("%d\t%02x\n", m[k], k)
		}
	case "rune":
		m, err := RuneFreq(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		for _, k := range keysSortedByDescValue(m) {
			fmt.Printf("%d\t%q\n", m[k], k)
		}
	default:
		log.Fatalln("error: invalid -by flag")
	}
}

func LineFreq(r io.Reader) (map[string]int, error) {
	m := make(map[string]int)
	s := bufio.NewScanner(r)
	for s.Scan() {
		m[s.Text()]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return m, nil
}

func ByteFreq(r io.Reader) (map[byte]int, error) {
	m := make(map[byte]int)
	br := bufio.NewReader(r)
	for {
		b, err := br.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		m[b]++
	}
	return m, nil
}

func RuneFreq(r io.Reader) (map[rune]int, error) {
	m := make(map[rune]int)
	br := bufio.NewReader(r)
	for {
		rn, sz, err := br.ReadRune()
		if rn == unicode.ReplacementChar && sz == 1 {
			return nil, errors.New("invalid utf8")
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		m[rn]++
	}
	return m, nil
}

// keysSortedByDescValue returns the keys in map m sorted by descending value.
func keysSortedByDescValue[K comparable](m map[K]int) []K {
	keys := maps.Keys(m)
	slices.SortFunc(keys, func(a, b K) int { return m[a] - m[b] })
	slices.Reverse(keys)
	return keys
}

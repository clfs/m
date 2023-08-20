package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/slices"
)

func main() {
	log.SetFlags(0)

	byFlag := flag.String("by", "line", "line, byte, rune, or word")
	flag.Parse()

	var splitFunc bufio.SplitFunc

	switch *byFlag {
	case "line":
		splitFunc = bufio.ScanLines
	case "byte":
		splitFunc = bufio.ScanBytes
	case "rune":
		splitFunc = bufio.ScanRunes
	case "word":
		splitFunc = bufio.ScanWords
	default:
		log.Fatalln("error: invalid -by value")
	}

	s := bufio.NewScanner(os.Stdin)
	s.Split(splitFunc)

	m := make(map[string]int)
	for s.Scan() {
		m[s.Text()]++
	}
	if err := s.Err(); err != nil {
		log.Fatalf("error: %v", err)
	}

	var format string

	switch *byFlag {
	case "line":
		format = "%d\t%s\n"
	case "byte":
		format = "%d\t0x%02x\n"
	case "rune":
		format = "%d\t%q\n"
	case "word":
		format = "%d\t%s\n"
	default:
		panic("unreachable")
	}

	pairs := toPairs(m)
	slices.SortFunc(pairs, cmp)

	for _, e := range pairs {
		fmt.Printf(format, e.count, e.value)
	}
}

type pair struct {
	count int
	value string
}

// {3,"a"} > {2,"a"} > {2,"b"}
func cmp(a, b pair) int {
	if a.count == b.count {
		switch {
		case a.value > b.value:
			return 1
		case a.value < b.value:
			return -1
		default:
			return 0
		}
	}
	return b.count - a.count
}

func toPairs(m map[string]int) []pair {
	pairs := make([]pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, pair{v, k})
	}
	return pairs
}

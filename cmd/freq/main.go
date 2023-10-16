package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

type mode struct {
	splitFunc bufio.SplitFunc
	format    string
}

var supportedModes = map[string]mode{
	"line": {bufio.ScanLines, "%d\t%s\n"},
	"byte": {bufio.ScanBytes, "%d\t0x%02x\n"},
	"rune": {bufio.ScanRunes, "%d\t%q\n"},
	"word": {bufio.ScanWords, "%d\t%s\n"},
}

func main() {
	log.SetFlags(0)

	byFlag := flag.String("by", "line", "line, byte, rune, or word")
	flag.Parse()

	byMode, ok := supportedModes[*byFlag]
	if !ok {
		log.Fatalln("error: invalid -by value")
	}

	s := bufio.NewScanner(os.Stdin)
	s.Split(byMode.splitFunc)

	m := make(map[string]int)
	for s.Scan() {
		m[s.Text()]++
	}
	if err := s.Err(); err != nil {
		log.Fatalf("error: %v", err)
	}

	pairs := toPairs(m)
	slices.SortFunc(pairs, cmp)

	for _, e := range pairs {
		fmt.Printf(byMode.format, e.count, e.value)
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

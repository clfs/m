package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	log.SetFlags(0)

	seed := flag.Int64("seed", 0, "seed for random number generator")
	flag.Parse()

	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error: %v", err)
	}

	rng := rand.New(rand.NewSource(*seed))
	perm := rng.Perm(len(lines))

	for _, i := range perm {
		fmt.Println(lines[i])
	}
}

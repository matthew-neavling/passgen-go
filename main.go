package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"passgen-go/wordlist"
	"strings"
)

func randomInt64Slice(n int, max int64) (slice []int64) {
	// Get a slice of `n` random integers using the `crypto/rand` library
	slice = []int64{}
	for range n {
		v, err := rand.Int(rand.Reader, big.NewInt(max))
		if err != nil {
			panic(err)
		}
		slice = append(slice, v.Int64())
	}
	return
}

func asciiProperCase(s string) string {
	// Quick and dirty proper case.
	// `s` is assumed to be a lower-case ASCII string
	return string(append([]byte{s[0] ^ 0x20}, s[1:]...))
}

func main() {
	effWordlist := wordlist.FromFile("eff.txt")
	ints := randomInt64Slice(3, int64(effWordlist.TotalLines()))

	words := []string{}
	for i := range len(ints) {
		words = append(words, asciiProperCase(effWordlist.GetWord(ints[i])))
	}
	fmt.Printf("%s\n", strings.Join(words,""))
}

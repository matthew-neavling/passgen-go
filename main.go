package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"passgen-go/wordlist"
	"regexp"
	"strings"
)

var (
	passwordLength     = flag.Int("length", 3, "Number of words to use in password")
	passwordIterations = flag.Int("iterations", 1, "Generate this number of passwords")
)

// randomInt64Slice returns a slice of `n` random integers using the `crypto/rand` library.
func randomInt64Slice(n int, max int64) (slice []int64) {
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

// asciiProperCase performs quick and dirty proper case.
// `s` is assumed to be a lower-case ASCII string.
func asciiProperCase(s string) string {
	return string(append([]byte{s[0] ^ 0x20}, s[1:]...))
}

type urlType int

const (
	file = iota
	remote
)

func inferURLType(url string) urlType {
	match, err := regexp.MatchString(`^https?://.*$`, url)
	if err != nil {
		return -1
	}
	if match {
		return remote
	}
	return file
}

func exit(msg string) {
	fmt.Println(msg)
	flag.Usage()
	os.Exit(1)
}

func main() {
	var wl *wordlist.Wordlist
	flag.Parse()

	wordlistURL := flag.Arg(0)
	if wordlistURL == "" {
		exit("no wordlist url or file provided")
	}
	switch inferURLType(wordlistURL) {
	case file:
		wordlist, err := wordlist.FromFile(wordlistURL)
		if err != nil {
			exit(err.Error())
		}
		wl = wordlist
	case remote:
		wordlist, err := wordlist.FromURL(wordlistURL)
		if err != nil {
			exit(err.Error())
		}
		wl = wordlist
	case -1:
		exit("provided wordlist url or file is invalid")
	}

	for range *passwordIterations {
		ints := randomInt64Slice(*passwordLength, int64(wl.TotalLines()))

		words := []string{}
		for i := range len(ints) {
			words = append(words, asciiProperCase(wl.GetWord(ints[i])))
		}
		fmt.Printf("%s\n", strings.Join(words, ""))
	}
}

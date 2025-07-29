package wordlist

import (
	"bufio"
	"os"
)

// wordlist contains the hashmap and convenience
// methods for interacting with a wordlist
type wordlist struct {
	words map[int64]string
}

// TotalLines returns the number of words present in the wordlist
func (wl *wordlist) TotalLines() int {
	return len(wl.words)
}

// GetWord returns the word with key `i`, which
// corresponds to its line number in the text file
func (wl *wordlist) GetWord(i int64) string {
	return wl.words[i]
}

// FromFile creates a wordlist from a text file.
// The file is assumed to be a list of words with
// a single word on each line.
func FromFile(filename string) wordlist {
	reader, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(reader)
	var lineNumber int64 = 0
	wl := wordlist{words: make(map[int64]string)}
	for {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		wl.words[lineNumber] = scanner.Text()

		hasMore := scanner.Scan()
		if !hasMore {
			break
		}

		lineNumber += 1
	}

	return wl
}

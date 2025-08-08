// Package wordlist contains the `Wordlist` struct for reading
// and returning words from wordlist files.
//
// Do not instantiate the `Wordlist` struct directly.
// Use `wordlist.FromFile`
package wordlist

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

type style int

const (
	regular = iota
	diceware
	notStyle
)

func inferStyle(line string) style {
	isRegular, err := regexp.MatchString(`^\w+$`, line)
	if err != nil {
		panic(err)
	}
	if isRegular {
		return regular
	}
	isDiceware, err := regexp.MatchString(`^\d+\s\w+$`, line)
	if err != nil {
		panic(err)
	}
	if isDiceware {
		return diceware
	}

	return notStyle
}

func dicewareToRegular(line string) string {
	pattern, err := regexp.Compile(`^\d+\s(\w+)$`)
	if err != nil {
		panic(err)
	}
	match := pattern.FindStringSubmatch(line)
	// fmt.Println(match)
	if len(match) == 0 {
		return ""
	}
	return match[1]
}

// Wordlist contains the hashmap and convenience
// methods for interacting with a Wordlist
type Wordlist struct {
	words map[int64]string
}

// TotalLines returns the number of words present in the wordlist
func (wl *Wordlist) TotalLines() int {
	return len(wl.words)
}

// GetWord returns the word with key `i`, which
// corresponds to its line number in the text file
func (wl *Wordlist) GetWord(i int64) string {
	return wl.words[i]
}

func readWordlist(reader io.Reader) (*Wordlist, error) {
	scanner := bufio.NewScanner(reader)
	var lineNumber int64 = 0
	var inferredStyle style = regular
	text := ""
	wl := Wordlist{words: make(map[int64]string)}
	for {
		if scanner.Err() != nil {
			return nil, fmt.Errorf("failed to read buffer %s", reader)
		}

		// Infer if wordlist is regular or diceware from first line
		if lineNumber == 1 {
			inferredStyle = inferStyle(scanner.Text())
			if inferredStyle > diceware {
				return nil, errors.New("file not a valid wordlist")
			}
		}

		text = scanner.Text()

		if inferredStyle == diceware {
			text = dicewareToRegular(text)
		}

		wl.words[lineNumber] = text

		hasMore := scanner.Scan()
		if !hasMore {
			break
		}

		lineNumber += 1
	}
	return &wl, nil

}

// FromFile creates a wordlist from a text file.
// The file is assumed to be a list of words with
// a single word on each line.
func FromFile(filename string) (*Wordlist, error) {
	reader, error := os.Open(filename)
	if error != nil {
		return nil, error
	}
	wl, err := readWordlist(reader)
	if err != nil {
		return nil, err
	}
	return wl, nil
}

// FromURL creates a wordlist from a file at a remote url
// The file is assumed to be a list of words with
// a single word on each line
func FromURL(url string) (*Wordlist, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	wl, err := readWordlist(resp.Body)
	if err != nil {
		return nil, err
	}
	return wl, nil
}

package wordlist

import (
	"bufio"
	"os"
)

type wordlist struct {
	words map[int64]string
}

func (wl *wordlist) TotalLines() int {
	return len(wl.words)
}

func (wl *wordlist) GetWord(i int64) string {
	return wl.words[i]
}

func FromFile(filename string) wordlist {
	reader, error := os.Open(filename)
	if error != nil { panic(error) }
	scanner := bufio.NewScanner(reader)
	var lineNumber int64 = 0
	wl := wordlist { words: make(map[int64]string) }
	for {
		if scanner.Err() != nil { panic(scanner.Err()); }

		wl.words[lineNumber] = scanner.Text()	

		hasMore := scanner.Scan()
		if !hasMore { break }

		lineNumber += 1
	}

	return wl
}


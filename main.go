package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type pattern struct {
	delim, from, to string
	word            string
}

func (p pattern) String() string {
	return p.word
}

func main() {
	var path string
	flag.StringVar(&path, "path", "words.txt", "path to file with words")
	flag.Parse()

	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	words := strings.Split(string(b), "\n")
	wordMap := map[string]struct{}{}

	re := regexp.MustCompile("^s([a-z])")
	patterns := []pattern{}
	for _, word := range words {
		wordMap[word] = struct{}{}

		submatches := re.FindStringSubmatch(word)
		if n := len(submatches); n == 0 { // delim not found
			continue
		} else if n == 1 {
			panic("this shouldn't have happened")
		}
		delim := submatches[1]

		p := fmt.Sprintf("^s%s([^%s]*)%s([^%s]*)%s$", delim, delim, delim, delim, delim)
		if x := regexp.MustCompile(p).FindStringSubmatch(word); len(x) == 3 {
			patterns = append(patterns, pattern{delim, x[1], x[2], word})
		}
	}

	sort.Slice(patterns, func(i, j int) bool {
		return patterns[i].word < patterns[j].word
	})

	wg := sync.WaitGroup{}
	for _, w := range words {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			for _, pattern := range patterns {
				after := strings.Replace(word, pattern.from, pattern.to, 1)
				if _, ok := wordMap[after]; ok && after != word {
					fmt.Println(word, pattern, after)
				}
			}
		}(w)
	}
	wg.Wait()
}

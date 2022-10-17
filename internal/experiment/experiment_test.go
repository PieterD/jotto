package experiment

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

const WordsFile = `../../words_alpha.txt`

func TestExperiment(t *testing.T) {
	wr := NewWordReader(WordsFile)
	ws := NewWordStore()
	err := wr.ReadWords(NewWordFilter(
		ws.AddWord,
		wordHasLengthFive,
		wordHasOnlyAlpha,
		wordHasUniqueLetters,
		//specificWordFilter("fldxt", "vejoz", "fconv", "gconv", "hdqrs", "expdt"),
	))
	require.NoError(t, err)
	t.Run("stats", func(t *testing.T) {
		for i, words := range ws.firstLetterIndex {
			fmt.Printf("letter %c: %d\n", i+'a', len(words))
		}
	})
	t.Run("enumerate", func(t *testing.T) {
		cur := NewWordletCursor(ws)
		totalWordlets := 0
		totalWords := 0
		for cur.Next() {
			totalWordlets++
			wl := cur.Value()
			words := ws.WordsFromWordlet(wl)
			//fmt.Printf("wordlet: %08X\n", wl)
			for _, word := range words {
				word = word
				//fmt.Printf("\tword: %s\n", word)
				totalWords++
			}
		}
		fmt.Printf("total wordlets: %d\n", totalWordlets)
		fmt.Printf("total words: %d\n", totalWords)
	})
	t.Run("find five", func(t *testing.T) {
		found := findFirstTetra(ws)
		fmt.Printf("found: %v\n", found)
		for _, wl := range found {
			fmt.Printf("\t%v\n", ws.WordsFromWordlet(wl))
		}
	})
	t.Run("find all", func(t *testing.T) {
		printAllTetra(ws)
	})
}

func findFirstTetra(store *WordStore) (found []Wordlet) {
	cur := NewWordletCursor(store)
	stack := make([]Wordlet, 0, 5)
	recursiveTetra(cur, stack, func(wordlets []Wordlet) (more bool) {
		found = make([]Wordlet, len(wordlets))
		copy(found, wordlets)
		return false
	})
	return found
}

func printAllTetra(store *WordStore) {
	cur := NewWordletCursor(store)
	stack := make([]Wordlet, 0, 5)
	totalSets := 0
	recursiveTetra(cur, stack, func(wordlets []Wordlet) (more bool) {
		totalSets++
		fmt.Printf("found: %v\n", wordlets)
		for _, wl := range wordlets {
			fmt.Printf("\t%v\n", store.WordsFromWordlet(wl))
		}
		return true
	})
	fmt.Printf("total sets: %d\n", totalSets)
}

func recursiveTetra(cur WordletCursor, stack []Wordlet, found func(wordlets []Wordlet) (more bool)) bool {
	if len(stack) == 5 {
		return found(stack)
	}
	for cur.Next() {
		wl := cur.Value()
		curNext := cur.Copy()
		curNext.Disallow(wl)
		biggerStack := append(stack, wl)
		more := recursiveTetra(curNext, biggerStack, found)
		if !more {
			return false
		}
	}
	return true
}

func wordHasLengthFive(word []byte) bool {
	return len(word) == 5
}

func wordHasOnlyAlpha(word []byte) bool {
	for _, c := range word {
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}

func wordHasUniqueLetters(word []byte) bool {
	if len(word) <= 1 {
		return true
	}
	var set [26]byte
	for _, c := range word {
		ci := c - 'a'
		if set[ci] == 1 {
			return false
		}
		set[ci] = 1
	}
	return true
}

func specificWordFilter(words ...string) WordFilter {
	wordMap := make(map[string]struct{})
	for _, word := range words {
		wordMap[word] = struct{}{}
	}
	return func(word []byte) (accept bool) {
		_, ok := wordMap[string(word)]
		return !ok
	}
}

func NewWordFilter(consumer WordConsumer, filters ...func(word []byte) (accept bool)) WordConsumer {
	return func(word []byte) error {
		for _, filter := range filters {
			if accept := filter(word); accept == false {
				return nil
			}
		}
		return consumer(word)
	}
}

func NewWordPrinter() WordConsumer {
	count := 0
	return func(word []byte) error {
		count++
		fmt.Printf("word %5d: %s\n", count, word)
		return nil
	}
}

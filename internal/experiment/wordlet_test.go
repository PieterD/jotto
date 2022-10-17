package experiment

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWordlet_FirstLetter(t *testing.T) {
	tests := []struct {
		word        string
		firstLetter byte
	}{
		{word: "hello", firstLetter: 'e'},
		{word: "world", firstLetter: 'd'},
		{word: "albert", firstLetter: 'a'},
		{word: "bra", firstLetter: 'a'},
		{word: "orb", firstLetter: 'b'},
		{word: "bore", firstLetter: 'b'},
		{word: "snore", firstLetter: 'e'},
		{word: "zzzz", firstLetter: 'z'},
	}
	for _, test := range tests {
		t.Run(test.word, func(t *testing.T) {
			wl := NewWordletFromWord([]byte(test.word))
			require.Equal(t, test.firstLetter, wl.FirstLetter())
		})
	}
}

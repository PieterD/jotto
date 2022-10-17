package experiment

type Word [5]byte

func (w Word) String() string {
	return string(w[:])
}

type WordStore struct {
	wordletToWords   map[Wordlet][]Word
	firstLetterIndex [26][]Wordlet
}

func NewWordStore() *WordStore {
	return &WordStore{
		wordletToWords: make(map[Wordlet][]Word),
	}
}

func (ws *WordStore) AddWord(word []byte) error {
	wl := NewWordletFromWord(word)
	var Word Word
	copy(Word[:], word)
	_, wordletAlreadyExists := ws.wordletToWords[wl]
	if !wordletAlreadyExists {
		idx := wl.FirstLetter() - 'a'
		ws.firstLetterIndex[idx] = append(ws.firstLetterIndex[idx], wl)
	}
	ws.wordletToWords[wl] = append(ws.wordletToWords[wl], Word)
	return nil
}

func (ws *WordStore) WordsFromWordlet(wl Wordlet) []Word {
	return ws.wordletToWords[wl]
}

type WordletCursor struct {
	store *WordStore

	disallowedLetters   Wordlet
	currentFirstLetter  byte
	currentWordletIndex int
	current             Wordlet
}

func NewWordletCursor(store *WordStore) WordletCursor {
	return WordletCursor{
		store: store,
	}
}

func (c WordletCursor) Copy() WordletCursor {
	return c
}

func (c *WordletCursor) Disallow(wl Wordlet) {
	c.disallowedLetters.AddWordlet(wl)
}

func (c *WordletCursor) Next() bool {
	if c.currentFirstLetter == 0 {
		c.currentFirstLetter = 'a'
	}
	for {
		if c.currentFirstLetter > 'z' {
			return false
		}
		if c.disallowedLetters.HasLetter(c.currentFirstLetter) {
			c.currentWordletIndex = 0
			c.currentFirstLetter++
			continue
		}
		fli := c.currentFirstLetter - 'a'
		storedWordlets := c.store.firstLetterIndex[fli]
		if c.currentWordletIndex >= len(storedWordlets) {
			c.currentWordletIndex = 0
			c.currentFirstLetter++
			continue
		}
		for c.currentWordletIndex < len(storedWordlets) {
			c.current = storedWordlets[c.currentWordletIndex]
			c.currentWordletIndex++
			if !c.disallowedLetters.Overlap(c.current) {
				return true
			}
		}
	}
}

func (c WordletCursor) Value() Wordlet {
	return c.current
}

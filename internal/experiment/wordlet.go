package experiment

import "math/bits"

type Wordlet uint32

func NewEmptyWordlet() Wordlet {
	return 0
}

func NewWordletFromWord(word []byte) Wordlet {
	s := NewEmptyWordlet()
	s.AddWord(word)
	return s
}

func (s *Wordlet) AddLetter(l byte) {
	l -= 'a'
	*s |= 1 << l
}

func (s *Wordlet) AddWord(word []byte) {
	for _, c := range word {
		s.AddLetter(c)
	}
}

func (s *Wordlet) AddWordlet(wl Wordlet) {
	*s |= wl
}

func (s Wordlet) HasLetter(l byte) bool {
	l -= 'a'
	return s&(1<<l) != 0
}

func (s Wordlet) Overlap(s2 Wordlet) bool {
	return s&s2 != 0
}

func (s Wordlet) Uint32() uint32 {
	return uint32(s)
}

func (s Wordlet) FirstLetter() byte {
	idx := bits.TrailingZeros32(s.Uint32())
	if idx >= 26 {
		panic("WTF")
	}
	return 'a' + byte(idx)
}

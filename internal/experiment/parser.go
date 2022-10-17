package experiment

import (
	"bufio"
	"fmt"
	"os"
)

type WordReader struct {
	fileName string
}

type (
	WordConsumer func(word []byte) error
	WordFilter   func(word []byte) bool
)

func NewWordReader(fileName string) *WordReader {
	return &WordReader{
		fileName: fileName,
	}
}

func (r *WordReader) ReadWords(consumer WordConsumer, filters ...WordFilter) error {
	h, err := os.Open(r.fileName)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer func() { _ = h.Close() }()
	s := bufio.NewScanner(h)
ScanLoop:
	for s.Scan() {
		word := s.Bytes()
		for _, filter := range filters {
			if accept := filter(word); accept == false {
				continue ScanLoop
			}
		}
		if err := consumer(word); err != nil {
			return err
		}
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("reading from file: %w", err)
	}
	return nil
}

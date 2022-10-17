# The Jotto Problem

From all known english words (dataset provided below):
- find a set of 5 words
- each 5 letters long
- with no repeated letters among them.

Example:

	dwarf
	vibex
	jocks
	glyph
	muntz

5 words with 5 letters each, 25 different letters.

For extra points, find all such sets.

Video that inspired this: https://www.youtube.com/watch?v=_-AfhLQfb6w

## Goals

- Make it faster than Matt's
- Use recursion because I think it fits
- No memory allocations

## My results

- experiment_test.go/find_five: <3s
- experiment_test.go/find_all:  <1m

## Source of "all known English words"

https://github.com/dwyl/english-words

It's very dirty but it's nice and big.

## License

Everything is free and may be used, copied, or reproduced by anyone for any reason.

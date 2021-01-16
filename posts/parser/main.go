package main

import (
	"io"
	"strings"
)

type parser struct {
	words int
	pos   int
}

type item int

type state int

const (
	Space item = iota
	Word
	Emph
	END
)

var handle = [4]func(rd io.Reader) int{
	start, space, word, emph, end,
}

func start(rd io.Reader) int {
	r, _, _ := rd.Rune()
	return 1
}

func space(rd io.Reader) int { return 1 }
func word(rd io.Reader) int  { return 1 }
func emph(rd io.Reader) int  { return 1 }
func end(rd io.Reader) int   { return 1 }

var rd = strings.NewReader(" ajdfkads fald *fa* df ads flakdsfalksdf jadf alsdkfj adslkfja ")


func main() {

	for {
		r, _, _ := rd.ReadRune()
		for {
			switch r {
			case ' ':
				n := new(

			}
		}
	}
}

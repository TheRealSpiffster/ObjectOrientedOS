package main

import "fmt"

type CLASSFILE struct {
	codeSegment   []byte
	dataSegment   []byte
	globalSegment []byte
	codeNames     map[string]uint32
	dataNames     map[string]uint32
	globalNames   map[string]uint32
}

type RULE struct {
	symbols []Symbol
	output  Symbol
	Reduce  func(int, []WORD) WORD
}

func parse(words []WORD) []WORD {
	var r RULE
	rules := []RULE{
		{[]Symbol{COMMENT, COMMENT}, COMMENT, r.DoCombine},
	}

	index := 0
	for len(words) > 1 {
		for _, r := range rules {
			if r.match(index, words) {
				fmt.Println("Matched rule ", r, "at", index)
				newWord := r.Reduce(index, words)
				newWord.symbol = r.output
				stopLoc := index + len(r.symbols)
				fwords := words[:index]
				lwords := words[stopLoc:]
				words = append(fwords, newWord)
				words = append(words, lwords...)
				index = 0
				break
			}
		}

		index++
		if index >= len(words) {
			fmt.Println("Could not finish parsing")
			break
		}
	}

	return words
}

func (r *RULE) match(index int, words []WORD) bool {
	for i, s := range r.symbols {
		if s != words[index+i].symbol {
			return false
		}
	}

	return true
}

func (r *RULE) DoCombine(index int, words []WORD) WORD {
	first := words[0]
	rest := words[1:]
	for _, w := range rest {
		first.code = append(first.code, w.code...)
		first.data += "\n" + w.data
	}

	return first
}

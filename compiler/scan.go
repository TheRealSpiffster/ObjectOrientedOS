package main

const (
	LOOKING = iota
	INCOMMENT
	INSPACES
	INQUOTE
	INNUMBER
	INWORD
)

type WORD struct {
	symbol Symbol
	data   string
	line   int
	code   []byte
}

type Scanner struct {
	words                      []WORD
	comment                    WORD
	quote                      WORD
	space                      WORD
	number                     WORD
	word                       WORD
	state                      int
	DoState                    []func(byte) bool
	startComment, startLiteral bool
	linecount                  int
}

func (s *Scanner) Scan(dat []byte) []WORD {
	s.words = make([]WORD, 0)
	s.state = LOOKING
	s.comment = WORD{COMMENT, "", 0, nil}
	s.quote = WORD{QUOTE, "", 0, nil}
	s.space = WORD{SPACE, "", 0, nil}
	s.number = WORD{NUMBER, "", 0, nil}
	s.word = WORD{ANYWORD, "", 0, nil}
	s.startComment = false
	s.startLiteral = false
	s.linecount = 1
	s.DoState = make([]func(byte) bool, 10)
	s.DoState[LOOKING] = s.inLooking
	s.DoState[INCOMMENT] = s.inComment
	s.DoState[INQUOTE] = s.inQuote
	s.DoState[INSPACES] = s.inSpaces
	s.DoState[INWORD] = s.inWord
	s.DoState[INNUMBER] = s.inNumber

	for _, c := range dat {
		// keep doing the state until we get false
		for s.DoState[s.state](c) {
		}
	}

	return s.words
}

func (s *Scanner) inLooking(c byte) bool {
	punctWord := map[byte]WORD{
		'*': {STAR, "*", 0, nil},
		'+': {PLUS, "+", 0, nil},
		'-': {MINUS, "-", 0, nil},
		'=': {EQUALS, "=", 0, nil},
		'&': {AMP, "&", 0, nil},
		'!': {BANG, "!", 0, nil},
		'~': {TILDE, "~", 0, nil},
		'.': {DOT, ".", 0, nil},
		',': {COMMA, ",", 0, nil},
		'<': {LESSTHAN, "<", 0, nil},
		'>': {GREATERTHAN, ">", 0, nil},
		':': {COLON, ":", 0, nil},
		';': {SEMICOLON, ";", 0, nil},
		'(': {LPAREN, "(", 0, nil},
		')': {RPAREN, ")", 0, nil},
		'[': {LBOX, "[", 0, nil},
		']': {RBOX, "]", 0, nil},
		'{': {LSQUIG, "{", 0, nil},
		'}': {RSQUIG, "}", 0, nil},
		'%': {PERCENT, "%", 0, nil},
		'^': {CARAT, "^", 0, nil},
	}

	spaceWord := map[byte]bool{
		' ':  true,
		'\n': true,
		'\r': true,
		'\t': true,
	}

	if c == '/' {
		if s.startComment {
			s.comment.data = ""
			s.comment.line = s.linecount
			s.state = INCOMMENT
			s.startComment = false
		} else {
			s.startComment = true
		}
	} else if c == '"' {
		s.state = INQUOTE
		s.quote.data = ""
		s.quote.line = s.linecount
	} else {
		if s.startComment {
			// process the singleton /
			s.words = append(s.words, WORD{DIVIDE, "/", s.linecount, nil})
			s.startComment = false
		}

		// check punctuation
		val, ok := punctWord[c]
		if ok {
			val.line = s.linecount
			s.words = append(s.words, val)
		} else if spaceWord[c] {
			s.state = INSPACES
			s.space.data = string(c)
		} else if c >= '0' && c <= '9' {
			s.state = INNUMBER
			s.number.data = string(c)
		} else {
			s.state = INWORD
			s.word.data = string(c)
		}
	}

	return false
}

func (s *Scanner) inComment(c byte) bool {
	if c != 13 && c != 10 {
		s.comment.data += string(c)
	}
	if c == 13 {
		s.comment.line = s.linecount
		s.words = append(s.words, s.comment)
		s.state = LOOKING
		s.linecount++
	}

	return false
}

func (s *Scanner) inQuote(c byte) bool {
	literals := map[byte]string{
		'\\': "\\",
		'n':  "\n",
		'r':  "\r",
		't':  "\t",
		'"':  "\"",
	}

	if c == 13 {
		s.linecount++
	}

	if s.startLiteral {
		s.quote.data += literals[c]
		s.startLiteral = false
	} else if c == '\\' {
		s.startLiteral = true
	} else if c == '"' {
		s.words = append(s.words, s.quote)
		s.state = LOOKING
	} else {
		s.quote.data += string(c)
	}

	return false
}

func (s *Scanner) inSpaces(c byte) bool {
	spaceWord := map[byte]bool{
		' ':  true,
		'\n': true,
		'\r': true,
		'\t': true,
	}

	if c == 13 {
		s.linecount++
	}

	if spaceWord[c] {
		s.space.data += string(c)
		return false
	} else {
		s.state = LOOKING
		return true // the last c needs to be consumed elsewhere
	}

}

func (s *Scanner) inNumber(c byte) bool {
	if c >= '0' && c <= '9' {
		s.number.data += string(c)
		return false
	} else {
		s.state = LOOKING
		s.words = append(s.words, s.number)
		return true
	}
}

func (s *Scanner) inWord(c byte) bool {
	if (c >= '0' && c <= '9') || c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		s.word.data += string(c)
		return false
	} else {
		s.state = LOOKING
		s.words = append(s.words, s.word)
		return true
	}
}

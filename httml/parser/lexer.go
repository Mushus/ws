package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

var debug = false

func setDebug(f bool) {
	debug = f
}

type lexer struct {
	input     *bufio.Reader
	buffer    bytes.Buffer
	nextToken []Token
	result    []Element
	pos       int
	start     int
}

func lex(r io.Reader) *lexer {
	return &lexer{
		input: bufio.NewReader(r),
	}
}

const eof = -1

func (l *lexer) peek() rune {
	lead, err := l.input.Peek(1)
	if err == io.EOF {
		return eof
	} else if err != nil {
		l.Error(err.Error())
	}

	p, err := l.input.Peek(runeLen(lead[0]))
	if err == io.EOF {
		return eof
	} else if err != nil {
		l.Error(err.Error())
	}
	r, _ := utf8.DecodeRune(p)
	return r
}

func runeLen(lead byte) int {
	if lead < 0xC0 {
		return 1
	} else if lead < 0xE0 {
		return 2
	} else if lead < 0xF0 {
		return 3
	} else {
		return 4
	}
}

const (
	leftDoctype        = "<!doctype"
	leftStartElement   = "<"
	rightElement       = ">"
	leftEndElement     = "</"
	rightSingleElement = "/>"
	leftComment        = "<!--"
	rightComment       = "-->"
	leftCDATA          = "<![CDATA["
	rightCDATA         = "]]>"
	whiteSpace         = " "
	tab                = "\t"
)

func (l *lexer) Lex(lval *yySymType) int {
	token := lval.token
	switch lval.token.Token {
	case LT: // <
		str := l.nextToWord(whiteSpace, tab, rightElement, rightSingleElement)
		if str == "" {
			l.Error("unexpected token")
		}
		token = Token{
			Token:   TAG_NAME,
			Literal: str,
		}
	case TAG_NAME:
		if l.hasPrefix(rightElement) {
			token = Token{
				Token:   GT,
				Literal: rightElement,
			}
			l.nextStringCount(rightElement)
		} else {
			l.Error("unexpected token")
		}
	case LEFT_COMMENT:
		if l.hasPrefix(rightComment) {
			token = Token{
				Token:   RIGHT_COMMENT,
				Literal: rightComment,
			}
			l.nextStringCount(rightComment)
		} else {
			str := l.nextToWord(rightComment)
			if str == "" {
				return eof
			}
			token = Token{
				Token:   COMMENT_TEXT,
				Literal: str,
			}
		}
	case COMMENT_TEXT:
		if l.hasPrefix(rightComment) {
			token = Token{
				Token:   RIGHT_COMMENT,
				Literal: rightComment,
			}
			l.nextStringCount(rightComment)
		} else {
			l.Error("found EOF")
		}
	case GT, RIGHT_COMMENT, TEXT_CONTENT:
		if l.hasPrefix(leftComment) {
			token = Token{
				Token:   LEFT_COMMENT,
				Literal: leftComment,
			}
			l.nextStringCount(leftComment)
		} else if l.hasPrefix(leftStartElement) {
			token = Token{
				Token:   LT,
				Literal: leftStartElement,
			}
			l.nextStringCount(leftStartElement)
		} else {
			str := l.nextToWord(leftStartElement)
			fmt.Println(str)
			if str == "" {
				return eof
			}
			token = Token{
				Token:   TEXT_CONTENT,
				Literal: str,
			}
		}
	default:
		if l.hasPrefix(leftComment) {
			token = Token{
				Token:   LEFT_COMMENT,
				Literal: leftComment,
			}
			l.nextStringCount(leftComment)
		} else if l.hasPrefix(leftStartElement) {
			token = Token{
				Token:   LT,
				Literal: leftStartElement,
			}
			l.nextStringCount(leftStartElement)
		} else {
			str := l.nextToWord(leftStartElement)
			token = Token{
				Token:   TEXT_CONTENT,
				Literal: str,
			}
			l.nextStringCount(str)
		}
	}
	l.clearBuffer()
	lval.token = token
	if debug {
		fmt.Printf("token: %#v\n", token)
	}
	return token.Token
}

func (l *lexer) hasPrefix(prefix string) bool {
	p, err := l.input.Peek(len(prefix))
	if err == io.EOF {
		return false
	} else if err != nil {
		l.Error(err.Error())
		return false
	}
	return string(p) == prefix
}

func (l *lexer) next() rune {
	r, w, err := l.input.ReadRune()
	if err == io.EOF {
		return eof
	}
	l.pos += w
	l.buffer.WriteRune(r)
	return r
}

func (l *lexer) nextToWord(word ...string) string {
Loop:
	for {
		for _, w := range word {
			if l.hasPrefix(w) {
				break Loop
			}
		}
		if r := l.next(); r == eof {
			return l.buffer.String()
		}
	}
	return l.buffer.String()
}

func (l *lexer) nextRuneCount(count int) {
	for i := 0; i < count; i++ {
		l.next()
	}
}

func (l *lexer) nextStringCount(text string) {
	l.nextRuneCount(utf8.RuneCountInString(text))
}

func (l *lexer) clearBuffer() {
	l.start = l.pos
	l.buffer.Truncate(0)
}

func (l *lexer) Error(e string) {
	panic(e)
}

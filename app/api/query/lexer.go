package query

import (
	"errors"
)

// lex クエリをトークンに分解します
func lex(query string) ([]Token, error) {
	l := &lexer{
		source: []rune(query),
	}
	if err := l.start(); err != nil {
		return nil, err
	}
	return l.tokens, nil
}

const (
	eof   rune = -1
	empty rune = -2
)

type lexer struct {
	source   []rune
	buf      runeStack
	pos      int
	startPos int
	tokens   []Token
}

// start 解析を開始します
func (l *lexer) start() error {
	for {
		r := l.peek()
		switch {
		case l.parseWhiteSpace():
		case r == eof:
			return nil
		case l.parseString():
		case l.parseVariable():
		case l.parseFloat():
		case l.parseInt():
		case l.parseOperatior():
		default:
			return errors.New("invalid token")
		}
	}
}

// rollback emit 直前までの状態に戻す
func (l *lexer) rollback() {
	l.pos = l.startPos
	l.buf.clear()
}

// next カーソルを1文字進めつつ読み込みます
func (l *lexer) next() rune {
	var r rune
	if len(l.source) <= l.pos {
		r = eof
	} else {
		r = l.source[l.pos]
	}
	l.buf.push(r)
	l.pos++
	return r
}

// peek カーソルを進めず1文字読み込みます
func (l *lexer) peek() rune {
	if len(l.source) <= l.pos {
		return eof
	}
	return l.source[l.pos]
}

// rewind カーソルを1文字巻き戻します
func (l *lexer) rewind() {
	if l.pos == 0 {
		return
	}
	l.pos--
	l.buf.pop()
}

// ignore 1文字を無視して取り出します
func (l *lexer) ignore() rune {
	if len(l.source) <= l.pos {
		return eof
	}
	r := l.source[l.pos]
	l.buf.pushEmpty()
	l.pos++
	return r
}

// emit 読み込んだトークンを確定します
func (l *lexer) emit(t tokenType) {
	token := Token{
		Typ:   t,
		Value: string(l.buf.takeout()),
	}
	l.tokens = append(l.tokens, token)
	l.clear()
}

func (l *lexer) clear() {
	l.startPos = l.pos
	l.buf.clear()
}

// parse statement

func (l *lexer) parseWhiteSpace() bool {
	for {
		switch l.next() {
		case ' ', '\n', '\r':
		default:
			l.rewind()
			l.clear()
			return false
		}
	}
}

// オペレーターをパースします
func (l *lexer) parseOperatior() bool {
	defer l.rollback()
	switch l.next() {
	case '!':
		switch l.peek() {
		case '=':
			l.next()
			l.emit(TokenTypCompOp)
		case '(':
			l.next()
			l.emit(TokenTypOpenNot)
		}
	case '(':
		l.emit(TokenTypOpenRB)
	case ')':
		l.emit(TokenTypCloseRB)
	case '@':
		if l.peek() == '>' {
			l.next()
		}
		l.emit(TokenTypCompOp)
	case '=':
		if l.next() != '=' {
			return false
		}
		l.emit(TokenTypCompOp)
	case '<':
		c := l.peek()
		if c == '=' || c == '@' {
			l.next()
		}
		l.emit(TokenTypCompOp)
	case '>':
		if l.peek() == '=' {
			l.next()
		}
		l.emit(TokenTypCompOp)
	case '&':
		if l.next() != '&' {
			return false
		}
		l.emit(TokenTypCondOp)
	case '|':
		if l.next() != '|' {
			return false
		}
		l.emit(TokenTypCondOp)
	case '.':
		l.emit(TokenTypOp)
	default:
		return false
	}
	return true
}

// parseString 文字列をパースします
func (l *lexer) parseString() bool {
	defer l.rollback()
	if l.ignore() != '"' {
		return false
	}
	for {
		r := l.peek()
		switch {
		case r == '\\':
			l.ignore()
			l.next()
		case r == '"':
			l.ignore()
			l.emit(TokenTypStr)
			return true
		case r == eof:
			return false
		default:
			l.next()
		}
	}
}

// parseVariable 変数をパースします
func (l *lexer) parseVariable() bool {
	defer l.rollback()
	c := l.next()
	switch {
	case 'A' <= c && c <= 'Z':
	case 'a' <= c && c <= 'z':
	case c == '_':
	default:
		return false
	}
	for {
		c := l.next()
		switch {
		case 'A' <= c && c <= 'Z':
		case 'a' <= c && c <= 'z':
		case '0' <= c && c <= '9':
		case c == '_':
		default:
			l.rewind()
			l.emit(TokenTypVar)
			return true
		}
	}
}

// parseFloat 少数をパースします
func (l *lexer) parseFloat() bool {
	defer l.rollback()
	hasDot := false
	hasSmall := false
	for {
		c := l.next()
		switch {
		case '0' <= c && c <= '9':
			if hasDot {
				hasSmall = true
			}
		case c == '.':
			if !hasDot {
				hasDot = true
			} else {
				// "." が複数あるのは少数ではない
				return false
			}
		default:
			if hasDot && hasSmall {
				l.rewind()
				l.emit(TokenTypFloat)
				return true
			}
			// 小数点と "." のあとに数字が存在しないのは少数ではない
			return false
		}
	}
}

// parseInt 整数をパースします
func (l *lexer) parseInt() bool {
	defer l.rollback()
	c := l.next()
	switch {
	case '1' <= c && c <= '9':
	default:
		return false
	}
	for {
		c := l.next()
		switch {
		case '0' <= c && c <= '9':
		default:
			l.rewind()
			l.emit(TokenTypInt)
			return true
		}
	}
}

// Token トークン
type Token struct {
	Typ   tokenType
	Value string
}

type tokenType int

const (
	// TokenTypUnknown 不明
	TokenTypUnknown tokenType = iota
	// TokenTypEOF 終端
	TokenTypEOF
	// TokenTypStr 文字列
	TokenTypStr
	// TokenTypVar 変数名
	TokenTypVar
	// TokenTypInt 整数値
	TokenTypInt
	// TokenTypFloat 少数
	TokenTypFloat
	// TokenTypOp オペレーター
	TokenTypOp
	// TokenTypCompOp 比較演算子
	TokenTypCompOp
	// TokenTypCondOp 条件演算子
	TokenTypCondOp
	// TokenTypCloseRB ")" 丸かっこ
	TokenTypCloseRB
	// TokenTypOpenRB ")"
	TokenTypOpenRB
	// TokenTypOpenNot "!("
	TokenTypOpenNot
)

// Eq トークンが一致しているかどうか
func (t Token) Eq(typ tokenType, value string) bool {
	return t.Typ == typ && t.Value == value
}

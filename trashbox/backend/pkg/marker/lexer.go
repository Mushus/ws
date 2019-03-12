package marker

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
		switch {
		case l.parseWhiteSpace():
		case l.peek() == eof:
			return nil
		case l.parseHeading():
		case l.parseHashTag():
		case l.parseText():
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
	for i := 0; ; i++ {
		switch l.next() {
		case ' ', '\n', '\r':
		default:
			if i == 0 {
				l.rollback()
				return false
			}
			l.rewind()
			l.emit(TokenTypWhiteSpace)
			return true
		}
	}
}

func (l *lexer) parseHeading() bool {
	for i := 0; ; i++ {
		switch l.next() {
		case '#':
		case ' ':
			if i == 0 {
				l.rollback()
				return false
			}
			l.rewind()
			l.emit(TokenTypHeading)
			return true
		default:
			l.rollback()
			return false
		}
	}
}

func (l *lexer) parseHashTag() bool {
	if l.next() != '#' {
		return false
	}
	switch l.peek() {
	case ' ', '\n', '\r', eof:
		return false
	}
	l.emit(TokenTypHashTag)
	for i := 0; ; i++ {
		switch l.next() {
		case ' ', '\n', '\r', eof:
			l.rewind()
			l.emit(TokenTypHashTagText)
			return true
		}
	}
}

func (l *lexer) parseText() bool {
	for i := 0; ; i++ {
		switch l.next() {
		case ' ', '\n', '\r', eof:
			l.rewind()
			if i == 0 {
				return false
			}
			l.emit(TokenTypText)
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
	// TokenTypLineBreak 改行
	TokenTypLineBreak
	// TokenTypWhiteSpace 空白文字
	TokenTypWhiteSpace
	// TokenTypHeading 見出し
	TokenTypHeading
	// TokenTypHash ハッシュタグ
	TokenTypHashTag
	// TokenTypHashTagText ハッシュタグテキスト
	TokenTypHashTagText
	// TokenTypText テキスト
	TokenTypText
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

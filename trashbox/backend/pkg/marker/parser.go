package marker

import (
	"fmt"
)

// Parse パースを行います
func Parse(query string) (AST, error) {
	tokens, err := lex(query)
	if err != nil {
		return nil, err
	}

	p := parser{
		tokens: tokens,
	}
	ast, err := p.start()
	if err != nil {
		return nil, err
	}
	return ast, nil
}

type parser struct {
	tokens []Token
	pos    int
}

// TokenEOF End Of File
var TokenEOF = Token{
	Typ:   TokenTypEOF,
	Value: "EOF",
}

func (p *parser) peek() Token {
	if len(p.tokens) <= p.pos {
		return TokenEOF
	}
	return p.tokens[p.pos]
}
func (p *parser) next() Token {
	if len(p.tokens) <= p.pos {
		return TokenEOF
	}
	t := p.tokens[p.pos]
	p.pos++
	return t
}

func (p *parser) start() (AST, error) {
	ast, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	t := p.peek()
	if p.peek() != TokenEOF {
		return nil, fmt.Errorf("unexpected token: %v", t.Value)
	}
	return ast, nil
}

func (p *parser) parseExpr() (AST, error) {
	token := p.peek()
	switch {
	case token.Typ == TokenTypOpenNot:
		p.next()
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		token := p.next()
		if !token.Eq(TokenTypCloseRB, ")") {
			return nil, fmt.Errorf(`unexpected token: %v`, token.Value)
		}
		return NotExpr{Expr: expr}, nil
	case token.Eq(TokenTypOp, "("):
		p.next()
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		token := p.next()
		if !token.Eq(TokenTypCloseRB, ")") {
			return nil, fmt.Errorf(`unexpected token: %v`, token.Value)
		}
		return expr, nil
	case token.Typ == TokenTypVar:
		left, err := p.parseCompExpr()
		if err != nil {
			return nil, err
		}

		op := p.peek()
		if !op.Eq(TokenTypCondOp, "&&") && !op.Eq(TokenTypCondOp, "||") {
			return left, nil
		}
		p.next()

		right, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		return BinaryExpr{
			Left:  left,
			Op:    &op,
			Right: right,
		}, nil
	default:
		return nil, fmt.Errorf(`unexpected token: %v`, token.Value)
	}
}

func (p *parser) parseCompExpr() (AST, error) {
	variable := p.next()
	if variable.Typ != TokenTypVar {
		return nil, fmt.Errorf(`unexpected token: %v`, variable.Value)
	}
	op := p.next()
	if op.Typ != TokenTypCompOp {
		return nil, fmt.Errorf(`unexpected token: %v`, op.Value)
	}
	val := p.next()
	switch val.Typ {
	case TokenTypStr:
	case TokenTypInt:
	case TokenTypFloat:
	default:
		return nil, fmt.Errorf(`unexpected token: %v`, val.Value)
	}
	return CompExpr{
		Var: &variable,
		Op:  &op,
		Val: &val,
	}, nil
}

// AST 抽象構文木
type AST interface{}

// BinaryExpr 条件式
type BinaryExpr struct {
	Op    *Token
	Left  AST
	Right AST
}

// NotExpr 否定式
type NotExpr struct {
	Expr AST
}

// CompExpr 比較式
type CompExpr struct {
	Var *Token
	Op  *Token
	Val *Token
}

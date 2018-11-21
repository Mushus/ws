package parser

import (
	"strings"
	"testing"
)

var okHTMLs = []string{
	``,
	`<!--hoge-->`,
	`<hoge>fuga`,
}

func TestOK(t *testing.T) {
	setDebug(true)
	for _, html := range okHTMLs {
		t.Logf("%#v\n", html)
		l := lex(strings.NewReader(html))
		yyParse(l)
		t.Logf("%#v\n", l.result)
	}
}

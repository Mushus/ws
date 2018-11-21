package query

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	cases := []struct {
		query string
		want  []Token
	}{
		{
			query: `"hello"`,
			want: []Token{
				{
					Typ:   TokenTypStr,
					Value: "hello",
				},
			},
		},
		{
			query: `"hel\"lo"`,
			want: []Token{
				{
					Typ:   TokenTypStr,
					Value: `hel"lo`,
				},
			},
		},
		{
			query: `hoge`,
			want: []Token{
				{
					Typ:   TokenTypVar,
					Value: "hoge",
				},
			},
		},
		{
			query: `hoge"hoge"`,
			want: []Token{
				{
					Typ:   TokenTypVar,
					Value: "hoge",
				},
				{
					Typ:   TokenTypStr,
					Value: "hoge",
				},
			},
		},
		{
			query: `100012345`,
			want: []Token{
				{
					Typ:   TokenTypInt,
					Value: "100012345",
				},
			},
		},
		{
			query: `1000.12345`,
			want: []Token{
				{
					Typ:   TokenTypFloat,
					Value: "1000.12345",
				},
			},
		},
		{
			query: `.12345`,
			want: []Token{
				{
					Typ:   TokenTypFloat,
					Value: ".12345",
				},
			},
		},
		{
			query: `(hoge)`,
			want: []Token{
				{
					Typ:   TokenTypOpenRB,
					Value: "(",
				},
				{
					Typ:   TokenTypVar,
					Value: "hoge",
				},
				{
					Typ:   TokenTypCloseRB,
					Value: ")",
				},
			},
		},
		{
			query: `( hoge )`,
			want: []Token{
				{
					Typ:   TokenTypOpenRB,
					Value: "(",
				},
				{
					Typ:   TokenTypVar,
					Value: "hoge",
				},
				{
					Typ:   TokenTypCloseRB,
					Value: ")",
				},
			},
		},
		{
			query: `hoge != "hoge"`,
			want: []Token{
				{
					Typ:   TokenTypVar,
					Value: "hoge",
				},
				{
					Typ:   TokenTypCompOp,
					Value: "!=",
				},
				{
					Typ:   TokenTypStr,
					Value: "hoge",
				},
			},
		},
	}
	for _, c := range cases {
		tokens, err := lex(c.query)
		if err != nil {
			t.Fatalf("case %q, not be parse error: %v", c.query, err)
		}
		if !reflect.DeepEqual(tokens, c.want) {
			t.Fatalf("case %q, expect %v, got %v", c.query, c.want, tokens)
		}
	}
}

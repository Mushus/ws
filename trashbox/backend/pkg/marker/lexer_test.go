package marker

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
			query: `# heading`,
			want: []Token{
				{
					Typ:   TokenTypHeading,
					Value: "#",
				},
				{
					Typ:   TokenTypWhiteSpace,
					Value: " ",
				},
				{
					Typ:   TokenTypText,
					Value: "heading",
				},
			},
		},
		{
			query: `#### heading`,
			want: []Token{
				{
					Typ:   TokenTypHeading,
					Value: "####",
				},
				{
					Typ:   TokenTypWhiteSpace,
					Value: " ",
				},
				{
					Typ:   TokenTypText,
					Value: "heading",
				},
			},
		},
		{
			query: `#hashtag ok`,
			want: []Token{
				{
					Typ:   TokenTypHashTag,
					Value: "#",
				},
				{
					Typ:   TokenTypHashTagText,
					Value: "hashtag",
				},
				{
					Typ:   TokenTypWhiteSpace,
					Value: " ",
				},
				{
					Typ:   TokenTypText,
					Value: "ok",
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

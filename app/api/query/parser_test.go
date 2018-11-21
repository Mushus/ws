package query_test

import (
	"reflect"
	"testing"

	"github.com/Mushus/app/api/query"
)

func TestParser(t *testing.T) {
	cases := []struct {
		query string
		want  query.AST
	}{
		{
			query: `!(hoge == "fuga")`,
			want: query.NotExpr{
				Expr: query.CompExpr{
					Var: &query.Token{
						Typ:   query.TokenTypVar,
						Value: "hoge",
					},
					Op: &query.Token{
						Typ:   query.TokenTypCompOp,
						Value: "==",
					},
					Val: &query.Token{
						Typ:   query.TokenTypStr,
						Value: "fuga",
					},
				},
			},
		},
		{
			query: `hoge != 0.1 && fuga @> "piyo"`,
			want: query.BinaryExpr{
				Left: query.CompExpr{
					Var: &query.Token{
						Typ:   query.TokenTypVar,
						Value: "hoge",
					},
					Op: &query.Token{
						Typ:   query.TokenTypCompOp,
						Value: "!=",
					},
					Val: &query.Token{
						Typ:   query.TokenTypFloat,
						Value: "0.1",
					},
				},
				Op: &query.Token{
					Typ:   query.TokenTypCondOp,
					Value: "&&",
				},
				Right: query.CompExpr{
					Var: &query.Token{
						Typ:   query.TokenTypVar,
						Value: "fuga",
					},
					Op: &query.Token{
						Typ:   query.TokenTypCompOp,
						Value: "@>",
					},
					Val: &query.Token{
						Typ:   query.TokenTypStr,
						Value: "piyo",
					},
				},
			},
		},
	}
	for _, c := range cases {
		token, err := query.Parse(c.query)
		if err != nil {
			t.Fatalf("case %q, not be parse error: %v", c.query, err)
		}
		if !reflect.DeepEqual(token, c.want) {
			t.Fatalf("expect %#v, got %#v", c.want, token)
		}
	}
}

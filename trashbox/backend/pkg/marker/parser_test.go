package marker_test

import (
	"reflect"
	"testing"

	"github.com/Mushus/trashbox/backend/pkg/marker"
)

func TestParser(t *testing.T) {
	cases := []struct {
		query string
		want  marker.AST
	}{
		{
			query: `!(hoge == "fuga")`,
			want: marker.NotExpr{
				Expr: marker.CompExpr{
					Var: &marker.Token{
						Typ:   marker.TokenTypVar,
						Value: "hoge",
					},
					Op: &marker.Token{
						Typ:   marker.TokenTypCompOp,
						Value: "==",
					},
					Val: &marker.Token{
						Typ:   marker.TokenTypStr,
						Value: "fuga",
					},
				},
			},
		},
		{
			query: `hoge != 0.1 && fuga @> "piyo"`,
			want: marker.BinaryExpr{
				Left: marker.CompExpr{
					Var: &marker.Token{
						Typ:   marker.TokenTypVar,
						Value: "hoge",
					},
					Op: &marker.Token{
						Typ:   marker.TokenTypCompOp,
						Value: "!=",
					},
					Val: &marker.Token{
						Typ:   marker.TokenTypFloat,
						Value: "0.1",
					},
				},
				Op: &marker.Token{
					Typ:   marker.TokenTypCondOp,
					Value: "&&",
				},
				Right: marker.CompExpr{
					Var: &marker.Token{
						Typ:   marker.TokenTypVar,
						Value: "fuga",
					},
					Op: &marker.Token{
						Typ:   marker.TokenTypCompOp,
						Value: "@>",
					},
					Val: &marker.Token{
						Typ:   marker.TokenTypStr,
						Value: "piyo",
					},
				},
			},
		},
	}
	for _, c := range cases {
		token, err := marker.Parse(c.query)
		if err != nil {
			t.Fatalf("case %q, not be parse error: %v", c.query, err)
		}
		if !reflect.DeepEqual(token, c.want) {
			t.Fatalf("expect %#v, got %#v", c.want, token)
		}
	}
}

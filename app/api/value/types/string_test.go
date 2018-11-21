package types_test

import (
	"testing"

	"github.com/Mushus/app/api/value/types"
	"github.com/Mushus/app/api/value/validation"
)

func TestStringValidation(t *testing.T) {
	cases := []struct {
		typ  types.String
		in   interface{}
		want validation.Result
	}{
		// MaxLength が 0 のときは指定なし
		{
			typ: types.String{
				Validation: validation.Option{
					MaxLength: 0,
				},
			},
			in:   "hoge",
			want: validation.Valid,
		},
		// min, max の指定値はエラーにならない(エッジケース)
		{
			typ: types.String{
				Validation: validation.Option{
					MaxLength: 4,
					MinLength: 4,
				},
			},
			in:   "hoge",
			want: validation.Valid,
		},
		// 👨‍👩‍👦‍👦 は7文字扱い
		{
			typ: types.String{
				Validation: validation.Option{
					MaxLength: 7,
					MinLength: 7,
				},
			},
			in:   "👨‍👩‍👦‍👦",
			want: validation.Valid,
		},
		// 改行文字の数の扱い
		{
			typ: types.String{
				Validation: validation.Option{
					MaxLength: 2,
					MinLength: 2,
				},
			},
			in:   "\r\n",
			want: validation.Valid,
		},
	}

	for _, c := range cases {
		got := c.typ.Validate(c.in)
		if got != c.want {
			t.Fatalf("Validation(v, %q) got %q, want %q", c.in, got, c.want)
		}
	}
}

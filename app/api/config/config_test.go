package config_test

import (
	"testing"

	"github.com/Mushus/app/api/config"
)

func TestCheckConfig(t *testing.T) {
	cases := []struct {
		in   config.Config
		want error
	}{
		// エンティティ未定義でエラーになるかどうか
		{
			in:   config.Config{},
			want: config.ErrEntitiesIsRequired,
		},
	}

	for _, c := range cases {
		got := config.Check(c.in)
		if got != c.want {
			t.Fatalf("Check(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}

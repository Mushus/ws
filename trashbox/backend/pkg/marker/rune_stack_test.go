package marker

import (
	"testing"
)

func TestRuneStack(t *testing.T) {
	rs := &runeStack{}
	rs.push('h')
	rs.push('e')
	rs.push('l')
	rs.push('l')
	rs.push('o')
	rs.push('世')
	rs.push('界')
	rs.push('!')
	want := "hello世界!"
	got := string(rs.takeout())
	if got != want {
		t.Fatalf("expect %v, got %v", want, got)
	}
}

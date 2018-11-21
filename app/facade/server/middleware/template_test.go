package middleware

import (
	"log"
	"path"
	"path/filepath"
	"testing"
)

func TestFilepath(t *testing.T) {
	name := "../hoge"
	fp := filepath.Join(ViewDir, "template", path.Clean("/"+name))
	log.Println(fp)
}

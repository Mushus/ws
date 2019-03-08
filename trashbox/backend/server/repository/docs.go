package server

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Mushus/trashbox/backend/pkg/encode/base64url"
	"golang.org/x/xerrors"
)

// DocumentNotFound is a error thrown when the document does not exist
var DocumentNotFound = xerrors.New("document not found")

// RawDocument is a document that is not parsed
type RawDocument struct {
	Body string
}

// DocRepo is a document repository
type DocRepo struct {
	docsDir string
}

// NewDocRepo is creating new document repository
func NewDocRepo() *DocRepo {
	return &DocRepo{
		docsDir: "docs",
	}
}

// Get is getting the document named name from fs
func (d DocRepo) Get(name string) (RawDocument, error) {
	// calc dirname
	b := sha256.Sum256([]byte(name))
	hash := hex.EncodeToString(b[:])
	dirName := hash[:2]

	// calc file name
	fileName := base64url.Encode(name)

	// document path
	path := filepath.Join(d.docsDir, dirName, fileName)

	if _, err := os.Stat(path); err != nil {
		return RawDocument{}, DocumentNotFound
	}

	body, err := ioutil.ReadFile(path)
	if err != nil {
		return RawDocument{}, xerrors.Errorf("cannot read document: %w", err)
	}

	return RawDocument{
		Body: string(body),
	}, nil
}

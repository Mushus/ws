package server

import (
	"bytes"
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
	Title   string
	Content string
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

// Get is getting the document named title from fs
func (d DocRepo) Get(title string) (RawDocument, error) {
	path := d.getFilePath(title)

	if _, err := os.Stat(path); err != nil {
		return RawDocument{}, DocumentNotFound
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return RawDocument{}, xerrors.Errorf("cannot read document: %w", err)
	}

	return RawDocument{
		Content: string(content),
	}, nil
}

// Put is putting the document doc
func (d DocRepo) Put(doc RawDocument) error {
	path := d.getFilePath(doc.Title)

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0666); err != nil {
		return xerrors.Errorf("cannot creat document directory: %w", err)
	}

	var body bytes.Buffer
	body.Write([]byte(doc.Title))
	body.Write([]byte("\n"))
	body.Write([]byte(doc.Content))

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return xerrors.Errorf("cannot open document: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(body.Bytes()); err != nil {
		return xerrors.Errorf("cannot write document: %w", err)
	}

	return nil
}

func (d DocRepo) getFilePath(title string) string {
	// calc dirname
	b := sha256.Sum256([]byte(title))
	hash := hex.EncodeToString(b[:])
	dirName := hash[:2]

	// calc file name
	fileName := base64url.Encode(title)

	// document path
	return filepath.Join(d.docsDir, dirName, fileName)
}

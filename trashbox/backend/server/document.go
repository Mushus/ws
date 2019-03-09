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

// Document is a document that is not parsed
type Document struct {
	Title   string
	Content string
}

// FileDocumentRepository is a document repository
type FileDocumentRepository struct {
	dir string
}

// NewFileDocumentRepository is creating new document repository
func NewFileDocumentRepository() *FileDocumentRepository {
	return &FileDocumentRepository{
		dir: "docs",
	}
}

// Get is getting the document named title from fs
func (f FileDocumentRepository) Get(title string) (Document, error) {
	path := f.getFilePath(title)

	if _, err := os.Stat(path); err != nil {
		return Document{}, DocumentNotFound
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Document{}, xerrors.Errorf("cannot read document: %w", err)
	}

	return Document{
		Content: string(content),
	}, nil
}

// Put is putting the document doc
func (f FileDocumentRepository) Put(doc Document) error {
	path := f.getFilePath(doc.Title)

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return xerrors.Errorf("cannot creat document directory: %w", err)
	}

	var body bytes.Buffer
	body.Write([]byte(doc.Title))
	body.Write([]byte("\n"))
	body.Write([]byte(doc.Content))

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return xerrors.Errorf("cannot open document: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(body.Bytes()); err != nil {
		return xerrors.Errorf("cannot write document: %w", err)
	}

	return nil
}

func (f FileDocumentRepository) getFilePath(title string) string {
	// calc dirname
	b := sha256.Sum256([]byte(title))
	hash := hex.EncodeToString(b[:])
	dirName := hash[:2]

	// calc file name
	fileName := base64url.Encode(title)

	// document path
	return filepath.Join(f.dir, dirName, fileName)
}

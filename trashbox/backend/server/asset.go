package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/rs/xid"
	"golang.org/x/xerrors"
)

var (
	// AssetNotFound is an error that occurs when an asset does not exist.
	AssetNotFound = xerrors.New("document not found")
	// IDFormat reprecents the format of ID.
	IDFormat = regexp.MustCompile("^[0-9a-v]{20}$")
	// FormatFormat is
	AssetFormatFormat = regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")
)

const (
	// InfoFileName is a file name of the asset information
	InfoFileName = "info.json"
	// OfirinalFileName is a file name of the raw asset
	OfirinalFileName = "@orig"
)

type (
	// Asset is a file uploaded by user
	Asset struct {
		ID          string        `json:"id"`
		Stream      io.ReadCloser `json:"-"`
		ContentType string        `json:"contentType"`
		FileName    string        `json:"fileName"`
	}

	// AssetInfo is a external information of assets
	AssetInfo struct {
		FileName    string `json:"fileName"`
		ContentType string `json:"contentType"`
	}
	// FileAssetRepository is a repository saving asset to file system
	FileAssetRepository struct {
		dir string
	}
)

func (a Asset) Read(p []byte) (n int, err error) {
	return a.Stream.Read(p)
}

// Close is
func (a Asset) Close() error {
	return a.Stream.Close()
}

// NewFileAssetRepository create new FileAssetRepository
func NewFileAssetRepository() *FileAssetRepository {
	return &FileAssetRepository{
		dir: "assets",
	}
}

// GetFormatedInStream gets asset by id
func (f FileAssetRepository) Get(id string) (Asset, error) {
	return f.getFormated(id, OfirinalFileName)
}

// GetFormatedInStream gets formated asset by id
func (f FileAssetRepository) getFormated(id, format string) (Asset, error) {
	if !f.ValidateID(id) || !f.ValidateFormat(format) {
		return Asset{}, AssetNotFound
	}

	dir := f.getDirPath(id)
	// read files
	infoFilePath := filepath.Join(dir, InfoFileName)
	formatedFilePath := filepath.Join(dir, format)

	// check file exists
	if _, err := os.Stat(formatedFilePath); err != nil {
		return Asset{}, AssetNotFound
	}

	// read info
	infoFile, err := os.Open(infoFilePath)
	if err != nil {
		return Asset{}, xerrors.Errorf("cannot open asset info: %w", err)
	}

	var info AssetInfo
	decoder := json.NewDecoder(infoFile)
	if err := decoder.Decode(&info); err != nil {
		return Asset{}, xerrors.Errorf("cannt read asset info: %w", err)
	}

	// read original
	file, err := os.Open(formatedFilePath)
	if err != nil {
		return Asset{}, xerrors.Errorf("cannot open asset: %w", err)
	}

	return Asset{
		ID:          id,
		Stream:      file,
		FileName:    info.FileName,
		ContentType: info.ContentType,
	}, nil
}

func (f FileAssetRepository) Add(asset Asset) (string, error) {
	id := xid.New().String()
	dir := f.getDirPath(id)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", xerrors.Errorf("cannot creat asset directory: %w", err)
	}

	if err := f.createInfo(dir, asset); err != nil {
		return "", err
	}
	if err := f.createOriginal(dir, asset); err != nil {
		return "", err
	}

	return id, nil
}

// ValidateID is used to Validate ID format.
// This method returns true if id Format is valid.
func (f FileAssetRepository) ValidateID(id string) bool {
	return IDFormat.MatchString(id)
}

// ValidateFormat is used to Validate format format.
func (f FileAssetRepository) ValidateFormat(format string) bool {
	return AssetFormatFormat.MatchString(format)
}

func (f FileAssetRepository) getDirPath(id string) string {
	// calc dirname
	b := sha256.Sum256([]byte(id))
	hash := hex.EncodeToString(b[:])
	dirName := hash[:2]

	// document path
	return filepath.Join(f.dir, dirName, id)
}

func (f FileAssetRepository) createInfo(dir string, asset Asset) error {
	infoFilePath := filepath.Join(dir, InfoFileName)
	// open asset info
	file, err := os.OpenFile(infoFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return xerrors.Errorf("cannot open asset info: %w", err)
	}
	defer file.Close()

	info := AssetInfo{
		FileName:    asset.FileName,
		ContentType: asset.ContentType,
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(info); err != nil {
		return xerrors.Errorf("cannot write asset info content: %w", err)
	}
	return nil
}

func (f FileAssetRepository) createOriginal(dir string, asset Asset) error {
	origFilePath := filepath.Join(dir, OfirinalFileName)
	// open original asset file
	file, err := os.OpenFile(origFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return xerrors.Errorf("cannot open original asset: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, asset); err != nil {
		return xerrors.Errorf("cannot write original asset: %w", err)
	}

	return nil
}

func (f FileAssetRepository) Remove(id string) error {
	return nil
}

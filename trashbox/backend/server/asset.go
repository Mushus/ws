package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/rs/xid"
	"golang.org/x/xerrors"
)

// AssetNotFound is a error thrown when the asset does not exist
var AssetNotFound = xerrors.New("document not found")

const (
	// InfoFileName is a file name of the asset information
	InfoFileName = "info.json"
	// OfirinalFileName is a file name of the raw asset
	OfirinalFileName = "orig"
)

type (
	// Asset is a file uploaded by user
	Asset struct {
		ID          string `json:"id"`
		ContentType string `json:"contentType"`
	}

	// StreamAsset is ...
	StreamAsset struct {
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
)

func (s StreamAsset) Read(p []byte) (n int, err error) {
	return s.Stream.Read(p)
}

func (s StreamAsset) Close() error {
	return s.Stream.Close()
}

// FileAssetRepository is a repository saving asset to file system
type FileAssetRepository struct {
	dir string
}

// NewFileAssetRepository create new FileAssetRepository
func NewFileAssetRepository() *FileAssetRepository {
	return &FileAssetRepository{
		dir: "assets",
	}
}

func (f FileAssetRepository) Get(id string) (Asset, error) {
	return Asset{
		ID: id,
	}, nil
}

func (f FileAssetRepository) GetInStream(id string) (StreamAsset, error) {
	dir := f.getDirPath(id)
	// read files
	infoFilePath := filepath.Join(dir, InfoFileName)
	originalFilePath := filepath.Join(dir, OfirinalFileName)

	// check file exists
	if _, err := os.Stat(infoFilePath); err != nil {
		return StreamAsset{}, AssetNotFound
	}

	// read info
	infoFile, err := os.Open(infoFilePath)
	if err != nil {
		return StreamAsset{}, xerrors.Errorf("cannot open asset info: %w", err)
	}

	var info AssetInfo
	decoder := json.NewDecoder(infoFile)
	if err := decoder.Decode(&info); err != nil {
		return StreamAsset{}, xerrors.Errorf("cannt read asset info: %w", err)
	}

	// read original
	file, err := os.Open(originalFilePath)
	if err != nil {
		return StreamAsset{}, xerrors.Errorf("cannot open asset: %w", err)
	}

	return StreamAsset{
		ID:          id,
		Stream:      file,
		FileName:    info.FileName,
		ContentType: info.ContentType,
	}, nil
}

func (f FileAssetRepository) Add(asset StreamAsset) (string, error) {
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

func (f FileAssetRepository) getDirPath(id string) string {
	// calc dirname
	b := sha256.Sum256([]byte(id))
	hash := hex.EncodeToString(b[:])
	dirName := hash[:2]

	// document path
	return filepath.Join(f.dir, dirName, id)
}

func (f FileAssetRepository) createInfo(dir string, asset StreamAsset) error {
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

func (f FileAssetRepository) createOriginal(dir string, asset StreamAsset) error {
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

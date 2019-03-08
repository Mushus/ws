package server

type DocumentRepository interface {
	Get(id string) Document
	Put(doc Document) error
}

type AssetRepository interface {
	Get(id string) Asset
	Put(doc Asset) error
}

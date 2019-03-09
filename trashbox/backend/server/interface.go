package server

type DocumentRepository interface {
	Get(id string) (Document, error)
	Put(doc Document) error
}

type AssetRepository interface {
	Get(id string) (Asset, error)
	Add(streamAsset Asset) (string, error)
	Remove(id string) error
}
type AssetCacheRepository interface {
	GetCache(id, format string) (Asset, error)
	PutCache(asset Asset, format string) error
	PurgeAll(id string) error
}

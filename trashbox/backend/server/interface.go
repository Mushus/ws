package server

type DocumentRepository interface {
	Get(id string) (Document, error)
	Put(doc Document) error
}

type AssetRepository interface {
	Get(id string) (Asset, error)
	GetInStream(id string) (StreamAsset, error)
	Add(streamAsset StreamAsset) (string, error)
}

package repository

import (
	"github.com/Mushus/app/api/value/model"
)

type DataStore interface {
	Migrate() error
	Restore(modelKey string, ids ...string) ([]model.Value, error)
	Save(model ...*model.Value) error
	Update(model ...*model.Value) error
}

type Model struct {
	Data map[string]interface{}
}

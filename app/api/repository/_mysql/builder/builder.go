package builder

import (
	"github.com/Mushus/app/api/repository/mysql/builder/schema"
	"github.com/Mushus/app/api/value/model"
)

// Builder SQLを生成するします
type Builder struct {
	models  model.Models
	schemas schema.Tables
}

func NewBulder(models model.Models) Builder {
	schemas := schema.NewTables(schemas)
	return Builder{
		models: models,
	}
}

func (b Builder) SaveSQL(entities entities) ([]query, error) {
	return entities.organize().toInsertSQLs(b.schema)
}

func (b Builder) RestoreSQL(key string, ids ...string) {

}

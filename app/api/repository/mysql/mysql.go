package mysql

import (
	"fmt"
	"strings"

	"github.com/Mushus/app/api/repository/mysql/schema"
	"github.com/Mushus/app/api/value/model"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	models         model.Models
	tables         schema.Tables
	host           string
	port           int
	user           string
	password       string
	database       string
	forceMigration bool
	db             *sqlx.DB
}

// New は MySQL のデータストアを作成します
func New(models model.Models, opts ...Option) (*MySQL, error) {
	tables := schema.NewTablesFromModels(models)
	ds := &MySQL{
		models:         models,
		tables:         tables,
		host:           "localhost",
		port:           3306,
		user:           "root",
		password:       "",
		database:       "app",
		forceMigration: false,
	}
	for _, opt := range opts {
		opt(ds)
	}
	db, err := sqlx.Connect(ds.ConnectTo())
	if err != nil {
		return nil, err
	}
	ds.db = db
	return ds, nil
}

// ConnectTo 接続先を取得します
func (d MySQL) ConnectTo() (string, string) {
	return "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.user, d.password, d.host, d.port, d.database)
}

func (d MySQL) ShowCreateTables() string {
	return strings.Join(d.tables.ToCreateSQL(), "\n")
}

func (m MySQL) Migrate() error {
	return nil
}

func (m MySQL) Restore(modelKey string, ids ...string) ([]model.Value, error) {
	return nil, nil
}

func (m MySQL) Save(model ...*model.Value) error {
	return nil
}

func (m MySQL) Update(model ...*model.Value) error {
	return nil
}

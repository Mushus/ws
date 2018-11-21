package mysql

import (
	"fmt"
	"log"

	"github.com/Mushus/app/api/repository"
	"github.com/Mushus/app/api/value"
	"github.com/Mushus/app/api/value/model"

	"github.com/jmoiron/sqlx"
)

// DataStore MySQL のデータストアです
type DataStore struct {
	builder 		builder.Builder
	models         value.Models
	tableSchemas   tableSchemas
	host           string
	port           int
	user           string
	password       string
	database       string
	forceMigration bool
	db             *sqlx.DB
}

// New は MySQL のデータストアを作成します
func New(models value.Models, opts ...Option) (*DataStore, error) {
	tableSchema := getModelSchema(models.List())
	ds := &DataStore{
		models:         models,
		tableSchemas:   tableSchema,
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
	db, err := sqlx.Connect(ds.connectTo())
	if err != nil {
		return nil, err
	}
	ds.db = db
	return ds, nil
}

// DB はデータベースクライアントを取得します
func (d DataStore) DB() *sqlx.DB {
	return d.db
}

// 接続先を取得します
func (d DataStore) connectTo() (string, string) {
	return "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.user, d.password, d.host, d.port, d.database)
}

// Migrate マイグレーションを実行します
func (d DataStore) Migrate() error {
	// _, err := d.gatherAllSchemas()
	// if err != nil {
	// 	return fmt.Errorf("cannot gather table schemas: %v", err)
	// }

	// st := getModelSchema(d.models.List())
	// log.Printf("%#v", st)

	// TODO: create 文しか生成してないのでAlter文生成するようにする
	// 上を比較すれば作れる
	sqls := d.Schemas()
	log.Printf("%#v", sqls)
	for _, sql := range sqls {
		_, err := d.db.Exec(sql)
		if err != nil {
			return fmt.Errorf("failed to exec migrate sql: %v", err)
		}
	}
	return nil
}

// Schemas は MySQLのスキーマを生成します
func (d DataStore) Schemas() []string {
	return d.tableSchemas.toCreateSQL()
}

// Restore MySQL からエンティティを取得します
func (d DataStore) Restore(modelKey string, ids ...string) (repository.Model, error) {
	return repository.Model{}, nil
}

// Save 値を保存します
func (d DataStore) Save(values ...*model.Value) error {
	entities := entities{}
	for _, v := range values {
		// NOTE: 引数 value に ID を振って返したいので
		v.ID()
		entities = append(entities, toEntities(*v)...)
	}

	queries, err := d.tableSchemas.createInsertSQL(entities)
	if err != nil {
		fmt.Errorf("failed to create Insert SQL: %v", err)
	}

	if err := d.execute(queries...); err != nil {
		return fmt.Errorf("failed to execute sql: %v", err)
	}

	return nil
}

// Update 記事をアップデートする
func (d DataStore) Update(values ...*model.Value) error {
	entities := []entity{}
	for _, v := range values {
		// HACK: ForUpdateを使って気をつけてSQLをかく
		// TODO: 更新するコードを書く
		entities = append(entities, toEntities(*v)...)
	}
	return nil
}

func (d DataStore) execute(queries ...query) error {
	for _, query := range queries {
		_, err := d.db.Exec(query.sql, query.args...)
		if err != nil {
			return err
		}
	}
	return nil
}

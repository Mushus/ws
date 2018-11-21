package datastore_test

import (
	"testing"

	"github.com/Mushus/app/api/repository/mysql"
	"github.com/Mushus/app/api/value/model"
	"github.com/Mushus/app/api/value/types"
	"github.com/Mushus/app/api/value/validation"
	_ "github.com/go-sql-driver/mysql"
)

var nop = validation.Option{}
var models = model.Models{
	model.New("top", types.Props{
		types.NewProp("title", types.String{}, false, nop),
		types.NewProp("sentence", types.String{}, false, nop),
		types.NewProp("main", types.NewList(types.String{}, nop), false, nop),
		types.NewProp("main2",
			types.NewList(
				types.NewList(
					types.NewGroup(types.Props{
						types.NewProp("hoge", types.String{}, false, nop),
						types.NewProp("fuga", types.String{}, false, nop),
					}, nop), nop,
				), nop,
			), false, nop,
		),
		types.NewProp("main3",
			types.NewList(
				types.NewGroup(types.Props{
					types.NewProp("hoge",
						types.NewList(
							types.NewGroup(types.Props{
								types.NewProp("hoge", types.String{}, false, nop),
								types.NewProp("fuga", types.String{}, false, nop),
							}, nop), nop,
						), false, nop,
					),
				}, nop), nop,
			),
			false, nop,
		),
	}),
}


func connect(t *testing.T) *mysql.DataStore {
	db, err := mysql.New(
		models,
		mysql.Host("localhost"),
		mysql.Port(3306),
		mysql.User("test"),
		mysql.Password("test"),
		mysql.Database("test"),
		mysql.ForceMigration(true),
	)
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.DB().MustExec("DROP TABLE IF EXISTS top")
	db.DB().MustExec("DROP TABLE IF EXISTS top_main_prop")
	db.DB().MustExec("DROP TABLE IF EXISTS top_photo_prop")
	return db
}

func TestMySQL(t *testing.T) {
	db := connect(t)

	migrate(t, db)
	save(t, db)
}

func TestMySQLRestore(t *testing.T) {
	/*db := connect(t)

	_, err := db.Restore("top")
	if err != nil {
		t.Fatalf("db.Restore() must not be error: %v", err)
	}*/
}

func migrate(t *testing.T, db *mysql.DataStore) {
	if err := db.Migrate(); err != nil {
		t.Fatalf("db.Migrate() must not be error: %v", err)
	}
}

func save(t *testing.T, db *mysql.DataStore) {
	key := "top"
	model, ok := models.Get(key)
	if !ok {
		t.Fatalf("invalid test code: model not found: %q", key)
	}

	Value, err := model.Cast(map[string]interface{}{
		"id":       "hogehoge",
		"title":    "hoge",
		"sentence": "hogehoge",
		"main":     []interface{}{"hoge", "fuga", "piyo"},
	})
	if err != nil {
		t.Fatalf("invalid test code: failed to cast: %v", err)
	}

	if err := db.Save(&Value); err != nil {
		t.Fatalf("db.Save() must not be error: %v", err)
	}
}

func restore(t *testing.T, db *mysql.DataStore) {

}

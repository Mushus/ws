package service

import (
	"github.com/jmoiron/sqlx"
)

// NewDB is
func NewDB() *sqlx.DB {
	cfg := NewCfg()
	if dbInstance == nil {
		db, err := sqlx.Open("sqlite", cfg.Get().DB.DSN)
		if err != nil {
			panic(err)
		}
		dbInstance = db
	}
	return dbInstance
}

var dbInstance *sqlx.DB

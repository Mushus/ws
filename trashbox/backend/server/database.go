package server

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

// TxFunc トランザクション処理
type TxFunc func(tx *sqlx.Tx) error

type dbx interface {
	Get(dest interface{}, query string, args ...interface{}) error
}

// DB データベース
type DB struct {
	db *sqlx.DB
}

// NewDB 指定した filepath をリポジトリにして、データベースを生成する
func NewDB(filepath string) (*DB, error) {
	sq, err := sqlx.Open("sqlite3", filepath)
	if err != nil {
		return nil, xerrors.Errorf("failed to open database: %w", err)
	}

	db := &DB{
		db: sq,
	}
	if err := db.ProvisionDB(); err != nil {
		return nil, xerrors.Errorf("failed to provision database: %w", err)
	}
	return db, nil
}

// ProvisionDB データベースを初期構築する
func (d *DB) ProvisionDB() error {
	db := d.db
	{
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
login TEXT NOT NULL,
password TEXT NOT NULL
)`)
		if err != nil {
			return xerrors.Errorf("failed to create user table: %w", err)
		}
	}
	{
		// デフォルトユーザーとして admin / admin を作成する
		u, err := NewUser("admin", "admin")
		if err != nil {
			return err
		}

		if err := d.AddUserIfNotExist(u); err != nil {
			return err
		}
	}
	return nil
}

// Tx トランザクション処理 txFunc を実行する
func (d *DB) Tx(txFunc TxFunc) error {
	tx, err := d.db.Beginx()
	if err != nil {
		return err
	}

	if err := txFunc(tx); err != nil {
		// ロールバック処理は失敗してもリカバリできないので放置
		// ロールバック処理失敗しても時間で解決される
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func findUserByLogin(db dbx, login string) (User, error) {
	var user User
	err := db.Get(&user, `SELECT id, login, password FROM users WHERE login = ? LIMIT 1`, login)
	if err != nil {
		return user, err
	}
	return user, nil
}

// AddUserIfNotExist 同一のログイン名の user がいない場合にユーザーを作成する
func (d *DB) AddUserIfNotExist(user User) error {
	return d.Tx(func(tx *sqlx.Tx) error {
		_, err := findUserByLogin(tx, user.Login)
		// ユーザーが存在する
		if err == nil {
			return nil
		}
		// エラー
		if err != sql.ErrNoRows {
			return xerrors.Errorf("cannot find user: %w", err)
		}
		_, err = tx.Exec(`INSERT INTO users (login, password) VALUES (?, ?)`, user.Login, user.Password)
		if err != nil {
			return xerrors.Errorf("cannot add user: %w", err)
		}
		return nil
	})
}

// VerifyUser ユーザー user を検証する
func (d *DB) VerifyUser(login, password string) (User, bool, error) {
	target, err := findUserByLogin(d.db, login)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, false, nil
		}
		return User{}, false, xerrors.Errorf("cannot verify user: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(target.Password), []byte(password)); err == nil {
		return target, true, nil
	}

	return User{}, false, nil
}

// User ユーザー情報
type User struct {
	ID       int    `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

// NewUser ユーザーを作成する
func NewUser(login, password string) (User, error) {
	u := User{
		Login: login,
	}
	err := u.SetPassword(password)
	if err != nil {
		return u, xerrors.Errorf("failed to create user: %w", err)
	}
	return u, nil
}

// SetPassword パスワードを password に変更する
func (u *User) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return xerrors.Errorf("cannot generate user password: %w", err)
	}
	u.Password = string(hashed)
	return nil
}

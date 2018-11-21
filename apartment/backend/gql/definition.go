package gql

// Schema is SQLite3 DB schema
const Schema = `
PRAGMA foreign_keys = ON;

CREATE TABLE articles (
	id   INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);

CREATE TABLE rooms (
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	article_id INTEGER NOT NULL,
	name       TEXT    NOT NULL,
	rent       INTEGER NOT NULL,
	'index'    INTEGER NOT NULL,
	FOREIGN KEY(article_id) REFERENCES articles(id)
);
CREATE INDEX idx__rooms__article_id ON rooms(article_id);

CREATE TABLE tenants (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,
	room_id     INTEGER   NOT NULL,
	name        TEXT      NOT NULL,
	rent        INTEGER   NOT NULL,
	since       DATE      NOT NULL,
	until       DATE      NOT NULL,
	FOREIGN KEY(room_id) REFERENCES rooms(id)
);
CREATE INDEX idx__tenants__room_id ON tenants(room_id);

CREATE TABLE billing_terms (
	id        INTEGER PRIMARY KEY AUTOINCREMENT,
	name      TEXT      NOT NULL,
	until     DATE      NOT NULL,
	since     DATE      NOT NULL
);

CREATE TABLE bills (
	id              INTEGER PRIMARY KEY AUTOINCREMENT,
	billing_term_id INTEGER   NOT NULL,
	tenant_id       INTEGER   NOT NULL,
	until           DATE      NOT NULL,
	since           DATE      NOT NULL,
	rent            INTEGER   NOT NULL,
	FOREIGN KEY(billing_term_id) REFERENCES billing_terms(id)
);
CREATE INDEX idx__bills__billing_term_id ON bills(billing_term_id);
`

// DateFormat 日付フォーマット
const DateFormat = "2006-01-02"

const (
	// StatusLiving 住んでる
	StatusLiving = iota
	// StatusEmpty 空き
	StatusEmpty
	// StatusReserved 予約
	StatusReserved
)

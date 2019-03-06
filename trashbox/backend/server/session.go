package server

import (
	"github.com/averagesecurityguy/random"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/michaeljs1990/sqlitestore"
	"golang.org/x/xerrors"
)

// SessionMaxAge is Max age of Sessions
const SessionMaxAge = 60 * 60 * 24 * 100

// SessionKeyUserId is used to obtain user ID from the session
const SessionKeyUserId = "userId"

// NewSession is coreate in-memory session
func NewSession() (echo.MiddlewareFunc, error) {
	secretKey, err := random.Token()
	if err != nil {
		return nil, xerrors.Errorf("cannot create session secret key: %w", err)
	}
	// TODO: this store has not feature of autoremove session
	store, err := sqlitestore.NewSqliteStore(":memory:", "sessions", "/", SessionMaxAge, []byte(secretKey))
	if err != nil {
		return nil, err
	}
	return session.Middleware(store), nil
}

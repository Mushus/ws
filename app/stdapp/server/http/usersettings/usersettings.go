package usersettings

import (
	"github.com/Mushus/app/stdapp/server/middleware"
)

func NewHandler() func(*middleware.SessCtx) error {
	return Handler
}

func Handler(c *middleware.SessCtx) error {
	return nil
}

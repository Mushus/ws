package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
)

type (
	SessCtx struct {
		echo.Context
		Session Sess
	}
	Sess struct {
		Hoge string `json:"hoge"`
	}
)

func Session(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess := &SessCtx{
			c,
		}
		if err := loadSession(sess); err != nil {
			return err
		}
		if err := next(sess); err != nil {
			return err
		}
		if err := saveSession(sess); err != nil {
			return err
		}
	}
}

// セッションを読み込む
func loadSession(c *SessCtx) error {
	header := c.Response().Header()
	sessJson := header.Get("X-Session", string(sessJSON))
	// NOTE: ヘッダーが存在しない場合空文字が帰るのでセッションを空にする
	if sessJson == "" {
		return nil
	}

	sess := Session{}

	// NOTE; セッションが壊れていたらLogで問題を追跡できるようにする
	if err := json.Unmarshal([]byte(sessJson), &sess); err != nil {
		c.Logger().Warrnf("broken session header, use zero values: %v", err)
	}

	c.Session = sess
	return nil
}

// セッションを保存します
func saveSession(c *SessCtx) error {
	sessJSON, err := json.Marshal(c.Session)
	if err != nil {
		return fmt.Errorf("failed to Marshal session: %v", err)
	}

	header := c.Response().Header()
	header.Add("X-Session", string(sessJSON))
	return nil
}

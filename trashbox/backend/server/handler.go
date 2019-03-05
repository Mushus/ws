package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

func loginPage(c echo.Context) error {
	return c.HTML(http.StatusOK, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Login</title>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
</head>
<body>
<form method="POST" action="login">
<input type="text" name="user" placeholder="user name">
<input type="text" name="password" placeholder="passowrd">
<button type="submit">Login</button>
</form>
</body>
</html>`)
}

// LoginParam ログイン
type LoginParam struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
}

func login(c echo.Context) error {
	var prm LoginParam
	if err := c.Bind(&prm); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := c.Validate(prm); err != nil {
		ReportValidation(err)
		//return err
	}

	return c.HTML(http.StatusOK, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Login</title>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
</head>
<body>
<form method="POST" action="login">
<input type="text" name="user" placeholder="user name">
<input type="text" name="password" placeholder="passowrd">
<button type="submit">Login</button>
</form>
</body>
</html>`)
}

func logout(c echo.Context) error {
	return c.HTML(http.StatusOK, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>Login</title>
<meta name="viewport" content="width=device-width,initial-scale=1.0">
</head>
ログアウトしました。
</body>
</html>`)
}

// ReportValidation バリデーション結果を分解します
func ReportValidation(errs error) map[string][]string {
	result := map[string][]string{}
	for _, err := range errs.(validator.ValidationErrors) {
		key := err.Field()
		if v, ok := result[key]; ok {
			result[key] = append(v, err.Tag())
		} else {
			result[key] = []string{err.Tag()}
		}
	}
	return result
}

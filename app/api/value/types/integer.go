package types

import (
	"errors"

	"github.com/Mushus/app/api/value/validation"
)

// Integer 整数値です
type Integer struct {
}

// Key は識別子を取得します
func (i Integer) Key() string {
	return "integer"
}

// Cast は値をキャストします
func (i Integer) Cast(value interface{}) (interface{}, error) {
	v, ok := value.(int)
	if !ok {
		return nil, errors.New("type is not Integer")
	}
	return v, nil
}

// Validate はバリデーションを行います
func (i Integer) Validate(value interface{}, vs ...validation.Option) validation.Result {
	return validation.Valid
}

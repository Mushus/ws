package types

import (
	"fmt"

	"github.com/Mushus/app/api/value/validation"
)

// ID 文字列を表します
type ID struct {
}

// Key は識別子を取得します
func (i ID) Key() string {
	return "id"
}

// Cast は値をキャストします
func (i ID) Cast(value interface{}) (interface{}, error) {
	v, ok := value.(string)
	if !ok {
		return TypedValue{}, fmt.Errorf("invalid type: %T", value)
	}
	return v, nil
}

// Validate はバリデーションを行います
func (i ID) Validate(value interface{}, vs ...validation.Option) validation.Result {
	_, ok := value.(string)
	if !ok {
		return validation.InvalidType
	}
	//TODO: バリデーション
	return validation.Valid
}

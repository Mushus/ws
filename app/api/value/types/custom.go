package types

import "github.com/Mushus/app/api/value/validation"

// Custom はユーザー定義型です
type Custom struct {
	key        string
	base       Type
	validation validation.Option
}

// Base 基底型を取得します
func (c Custom) Base() Type {
	base, ok := c.base.(Custom)
	if !ok {
		return c.base
	}
	return base.Base()
}

// Key は識別子を取得します
func (c Custom) Key() string {
	return c.key
}

// Cast は値をキャストします
func (c Custom) Cast(value interface{}) (interface{}, error) {
	return c.base.Cast(value)
}

// Validate バリデーションを実行します
func (c Custom) Validate(value interface{}, vs ...validation.Option) validation.Result {
	validations := append([]validation.Option{c.validation}, vs...)
	return c.base.Validate(value, validations...)
}

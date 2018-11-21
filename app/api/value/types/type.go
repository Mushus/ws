package types

import "github.com/Mushus/app/api/value/validation"

// TypedValue は型が設定された値です
type TypedValue struct {
	typ Type
	val interface{}
}

// Value 値を取得します
func (t TypedValue) Value() interface{} {
	return t.val
}

// Type はエンティティのプロパティの型を表します。
type Type interface {
	// 識別子
	Key() string
	Cast(value interface{}) (interface{}, error)
	Validate(value interface{}, vs ...validation.Option) validation.Result
}

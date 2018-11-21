package types

import (
	"errors"

	"github.com/Mushus/app/api/value/validation"
)

// NewVariable グループを生成します。
// グループは props を持っています。
// グループは nullable 不可です。
func NewVariable(props Props, v validation.Option) Variable {
	return Variable{
		props:      props,
		validation: v,
	}
}

// Group 型です
type Variable struct {
	props      Props
	validation validation.Option
}

// Key 識別子を取得します
func (g Variable) Key() string {
	return "group"
}

// Props プロパティを取得します
func (g Variable) Props() Props {
	return g.props
}

// Cast リストへの変換を行います
func (g Variable) Cast(value interface{}) (interface{}, error) {
	// TODO: 未実装
	return value, nil
}

// Solve 自身の型の値に変換します
func (g Variable) Solve(value interface{}) (map[string]interface{}, error) {
	v, myValue := value.(map[string]interface{})
	if !myValue {
		return nil, errors.New("unsupported value")
	}
	return v, nil
}

// Validate リストのバリデーションを行います
func (g Variable) Validate(value interface{}, vs ...validation.Option) validation.Result {
	// TODO: 未実装
	return validation.Valid
}

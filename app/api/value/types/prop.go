package types

import (
	"fmt"

	"github.com/Mushus/app/api/value/validation"
)

type PropKey string

// Props はプロパティ一覧です
type Props []Prop

func (p Props) Find(key PropKey) (prop *Prop, isfound bool) {
	for _, prop := range p {
		if prop.Key() == key {
			return &prop, true
		}
	}
	return nil, false
}

// NewProp プロパティを作成します
func NewProp(key PropKey, typ Type, nullable bool, validation validation.Option) Prop {
	return Prop{
		key:        key,
		typ:        typ,
		nullable:   nullable,
		validation: validation,
	}
}

// Prop エンティティのプロパティを表します
type Prop struct {
	key        PropKey
	typ        Type
	nullable   bool
	validation validation.Option
}

// Key プロパティ名を取得します
func (p Prop) Key() PropKey {
	return p.key
}

// Type 型を取得します
func (p Prop) Type() Type {
	return p.typ
}

// Nullable null 可能かどうかを取得します
func (p Prop) Nullable() bool {
	return p.nullable
}

// Validation バリデーションのオプションを取得します
func (p Prop) Validation() validation.Option {
	return p.validation
}

// MaxSize プロパティの最大容量を取得します
func (p Prop) MaxSize() int {
	v := p.Validation()
	return v.MaxLength
}

// Cast キャストします
func (p Prop) Cast(value interface{}) (interface{}, error) {
	if !p.nullable && value == nil {
		return nil, fmt.Errorf("prop %q must not be null", p.Key())
	}
	typ := p.Type()
	typedValue, err := typ.Cast(value)
	if err != nil {
		return nil, err
	}
	return typedValue, nil
}

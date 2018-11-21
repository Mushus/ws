package types

import (
	"errors"

	"github.com/Mushus/app/api/value/validation"
	"github.com/rs/xid"
)

// NewList はリスト型を生成します。
// 生成されるリストは要素の型 typ です。
// list は nullable 不可です
func NewList(typ Type, v validation.Option) List {
	return List{
		itemsType:  typ,
		validation: v,
	}
}

// List 型です
type List struct {
	itemsType  Type
	validation validation.Option
}

func (l List) GenerateID() string {
	return xid.New().String()
}

// Key 識別子を取得します
func (l List) Key() string {
	return "list"
}

// ItemsType アイテムの型を取得します
func (l List) ItemsType() Type {
	return l.itemsType
}

// Cast リストへの変換を行います
func (l List) Cast(value interface{}) (interface{}, error) {
	list, ok := value.([]interface{})
	if !ok {
		return nil, errors.New("type is not List")
	}
	for _, v := range list {
		_, err := l.ItemsType().Cast(v)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

// Solve 自身の型の値に変換します
func (l List) Solve(value interface{}) ([]interface{}, error) {
	v, myValue := value.([]interface{})
	if !myValue {
		return nil, errors.New("unsupported value")
	}
	return v, nil
}

// Validate リストのバリデーションを行います
func (l List) Validate(value interface{}, vs ...validation.Option) validation.Result {
	list, ok := value.([]interface{})
	if !ok {
		return validation.InvalidType
	}

	validations := append([]validation.Option{l.validation}, vs...)

	for _, v := range validations {
		length := len(list)
		if v.MaxLength < length {
			return validation.InvalidMaxListLength
		}
		for _, listVal := range list {
			res := l.ItemsType().Validate(listVal)
			if res != validation.Valid {
				return res
			}
		}
	}

	return validation.Valid
}

package model

import (
	"github.com/Mushus/app/api/value/types"
	"github.com/rs/xid"
)

// Value は型検査済みの値を表します
type Value struct {
	model *Model
	val   map[types.PropKey]interface{}
}

// Model モデルを取得します
func (v Value) Model() *Model {
	return v.model
}

// Value 値を取得します
func (v Value) Value() map[types.PropKey]interface{} {
	return v.val
}

// ID を取得します
func (v *Value) ID() string {
	for key, val := range v.val {
		if key == "id" {
			id, ok := val.(string)
			if ok {
				return id
			}
		}
	}
	id := xid.New().String()
	v.val["id"] = id
	return id
}

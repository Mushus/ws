package model

import (
	"fmt"

	"github.com/Mushus/app/api/value/types"
	"github.com/Mushus/app/api/value/validation"
)

// モデルのキー
type Key string

// Model エンティティです
type Model struct {
	// 識別子
	key Key
	// ラベル用プロパティ
	labelProp string
	// プロパティ一覧
	props types.Props
}

var idProp = types.NewProp("id", types.ID{}, false, validation.Option{})

// New はエンティティを作成します
func New(key Key, props types.Props) *Model {
	props = append(types.Props{idProp}, props...)
	return &Model{
		key:   key,
		props: props,
	}
}

// Cast は値 value を型に当てはめます。
// 当てはめた結果が返ります。
func (e Model) Cast(value interface{}) (Value, error) {
	values, ok := value.(map[types.PropKey]interface{})
	if !ok {
		return Value{}, fmt.Errorf("invalid value: %T", value)
	}

	propsValue := map[types.PropKey]interface{}{}
	for _, prop := range e.props {
		key := prop.Key()
		val, ok := values[key]
		if !ok {
			return Value{}, fmt.Errorf("prop %q not exist", key)
		}
		castVal, err := prop.Cast(val)
		if err != nil {
			return Value{}, err
		}
		propsValue[key] = castVal
	}

	return Value{
		model: &e,
		val:   propsValue,
	}, nil
}

// Key はエンティティの識別子を返します
func (e Model) Key() Key {
	return e.key
}

// Props はエンティティのプロパティを返します
func (e Model) Props() types.Props {
	return e.props
}

// Models モデルのリストです
type Models []*Model

// Add エンティティをリストに追加します
func (m *Models) Add(model *Model) {
	*m = append(*m, model)
}

// Get エンティティを取得します
func (m Models) Get(key Key) *Model {
	for _, v := range m {
		if v.Key() == key {
			return v
		}
	}
	return nil
}

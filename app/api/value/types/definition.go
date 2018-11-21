package types

import (
	"fmt"

	"github.com/Mushus/app/api/config"
	"github.com/Mushus/app/api/value/validation"
)

// Types タイプ一覧
type Types map[string]Type

// Get は定義された型を取得します
func (t Types) Get(key string) (Type, error) {
	typ, ok := t[key]
	if !ok {
		return nil, fmt.Errorf(`type "%s" is not defied`, key)
	}
	return typ, nil
}

// CreateList リストを作成します
func CreateList(typCfg config.Types) (Types, error) {
	typs := CreateDefaults()
	for _, typ := range typCfg {
		base, err := typs.Get(typ.Base)
		if err != nil {
			return nil, err
		}

		typs[typ.Key] = Custom{
			key:        typ.Key,
			base:       base,
			validation: validation.CreateOption(typ.Validation),
		}
	}
	return typs, nil
}

// CreateDefaults デフォルトのタイプを定義する
func CreateDefaults() Types {
	// なるべく汎用的な基本型を定義して、エイリアスとして名前をつけて派生型を使用する方針
	return Types{
		// NOTE: わかりやすいという理由で string ばかり使われると雑な設計をされてしまうので使用しない
		// "string": types.String{},
		"text": String{
			Validation: validation.Option{
				NGWord: []string{"\r", "\n"},
			},
		},
		"sentence": String{},
		"integer":  Integer{},
	}
}

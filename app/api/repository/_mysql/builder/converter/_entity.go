package converter

import (
	"fmt"

	"github.com/Mushus/app/api/value/model"
	"github.com/Mushus/app/api/value/types"
)

// toEntities モデルからエンティティを作成します
func CreateTableSchema(v model.Value) Tables {
	entities := []Record{}
	values := v.Value()
	model := v.Model()
	props := model.Props()

	entityVal := entityValue{}
	for key, value := range values {
		prop, found := props.Find(key)
		if !found {
			// NOTE: 事前条件ミス
			// value の cast あたりが間違ってるかもしれない
			panic("invalid precondition: toEntities")
		}
		entities = append(entities, createEntityValues(prop.Type(), key, value, model.Key(), v.ID(), &entityVal)...)
	}

	return append(entities, entity{
		tableName: tableName(model.Key()),
		value:     entityVal,
	})
}

// createEntityValues エンティティの値を作成します
func createEntityValues(propType types.Type, propKey string, propVal interface{}, parentKey string, parentID string, ev *entityValue) []entity {
	// NOTE: listに直接value等、名前のつけようもないカラムは value
	if propKey == "" {
		propKey = "value"
	}
	entities := []entity{}

	switch typ := propType.(type) {
	case types.Group:
		addGroupValue(typ, propVal, propKey, ev)
	case types.List:
		entities = append(entities, toListEntities(typ, propVal, parentKey, parentID, propKey)...)
	case types.Custom:
		entities = append(entities, createEntityValues(typ.Base(), propKey, propVal, parentKey, parentID, ev)...)
	default:
		(*ev)[propKey] = propVal
	}

	return entities
}

// リストテーブル
func toListEntities(typ types.Type, value interface{}, parentKey string, parentID string, propKey string) []entity {
	switch listType := typ.(type) {
	case types.List:
		list, err := listType.Solve(value)
		if err != nil {
			panic(fmt.Errorf("invalid precondition: %v", err))
		}

		entities := make([]entity, len(list))
		for i, v := range list {
			myID := listType.GenerateID()
			foreignKey := parentKey + "_id"
			ev := entityValue{
				"id":       myID,
				foreignKey: parentID,
				"order":    i,
			}

			entities = append(entities, createEntityValues(listType.ItemsType(), "", v, propKey, myID, &ev)...)
			entities[i] = entity{
				tableName: tableName(parentKey + "_" + propKey + "_prop"),
				value:     ev,
			}
		}
		return entities
	case types.Custom:
		return toListEntities(listType.Base(), value, parentKey, parentID, propKey)
	default:
		panic("invalid precondition: toListEntities")
	}
}

// addGroupValue groupをvalueに変換します
func addGroupValue(group types.Group, value interface{}, parentKey string, ev *entityValue) {
	prefix := parentKey + "_"

	groupValues, err := group.Solve(value)
	if err != nil {
		// NOTE: ここに来たら事前条件ミス
		panic("invalid precondition")
	}

	for key, gv := range groupValues {
		entityKey := prefix + key
		(*ev)[entityKey] = gv
	}
}

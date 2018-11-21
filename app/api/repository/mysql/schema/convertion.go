package schema

import (
	"github.com/Mushus/app/api/value/model"
	"github.com/Mushus/app/api/value/types"
)

// ConvertionInfos
type ConvertionInfos []ConvertionInfo

func NewCovertionInfos(models model.Models) ConvertionInfos {
	ci := ConvertionInfos{}
	for _, model := range models {
		tables := NewTablesFromModel(model)
		for _, table := range tables {
			ci = append(ci, ConvertionInfo{
				table:    table,
				modelKey: model.Key(),
			})
		}
	}
	return ci
}

func NewConvertionInfo(model model.Model) ConvertionInfos {
	modelKey := model.Key()
	table := NewMainTable(model)
	tableName := table.Name()
	ci := ConvertionInfos{
		{
			modelKey: modelKey,
			table:    table,
		},
	}

	props := model.Props()
	hierarchy := []string{string(modelKey)}
	return append(ci, NewConvertionInfoFromProps(modelKey, props, tableName, hierarchy)...)
}

func NewConvertionInfoFromProps(modelKey model.Key, props types.Props, parentName TableName, hierarchy []string) ConvertionInfos {
	ci := ConvertionInfos{}
	for _, prop := range props {
		ci = append(ci, NewConvertionInfoFromProp(modelKey, prop, parentName, hierarchy)...)
	}
	return ci
}

func NewConvertionInfoFromProp(modelKey model.Key, prop types.Prop, parentName TableName, hierarchy []string) ConvertionInfos {
	typ := prop.Type()
	return NewConvertionInfoFromType(modelKey, typ, parentName, hierarchy, prop.Key())
}

func NewConvertionInfoFromType(modelKey model.Key, typ types.Type, parentName TableName, hierarchy []string, propKey types.PropKey) ConvertionInfos {
	switch v := typ.(type) {
	case types.List:
		listTable := NewListTable(v, parentName, hierarchy, propKey)
		ci := ConvertionInfos{
			{
				modelKey: modelKey,
				table:    listTable,
			},
		}
		myHierarchy := append([]string{}, hierarchy...)
		myHierarchy = append(myHierarchy, string(propKey))
		return append(ci, NewConvertionInfoFromType(modelKey, v.ItemsType(), listTable.Name(), myHierarchy, "")...)
	case types.Variable:
		ci := ConvertionInfos{}
		props := v.Props()
		variableTables := NewVariableTables(v, parentName, hierarchy, propKey)
		for i, prop := range props {
			table := variableTables[i]
			ci = append(ci, ConvertionInfo{
				modelKey: modelKey,
				table:    table,
			})
			tableName := table.Name()
			myHierarchy := append([]string{}, hierarchy...)
			myHierarchy = append(myHierarchy, string(propKey))
			ci = append(ci, NewConvertionInfoFromType(modelKey, prop.Type(), tableName, myHierarchy, "")...)
		}
		return ci
	case types.Group:
		return NewConvertionInfoFromProps(modelKey, v.Props(), parentName, hierarchy)
	default:
	}
	return ConvertionInfos{}
}

type ConvertionInfo struct {
	modelKey model.Key
	table    Table
}

package schema

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Mushus/app/api/value/model"
	"github.com/Mushus/app/api/value/types"
)

// Tables テーブル情報です
type Tables []Table

// NewTablesFromModels テーブル一覧を生成します
func NewTablesFromModels(models model.Models) Tables {
	tables := Tables{}
	for _, model := range models {
		tables = append(tables, NewTablesFromModel(model)...)
	}
	return tables
}

// NewTablesFromModel テーブルのスキーマをモデルから作成します
func NewTablesFromModel(m model.Model) Tables {
	main := NewMainTable(m)
	mainName := main.Name()
	tables := Tables{main}
	props := m.Props()
	return append(tables, NewTablesFromProps(props, mainName, []string{string(m.Key())})...)
}

// NewTablesFromProps スキーマを取得します
// 再帰的にテーブルにできそうなスキーマを探索します
func NewTablesFromProps(props types.Props, parentName TableName, hierarchy []string) Tables {
	tables := Tables{}
	for _, prop := range props {
		tables = append(tables, NewTableFromProp(prop, parentName, hierarchy)...)
	}
	return tables
}

// NewTableFromProp プロパティからテーブル情報を生成します
func NewTableFromProp(prop types.Prop, parentName TableName, hierarchy []string) Tables {
	typ := prop.Type()
	return NewTablesFromType(typ, parentName, hierarchy, prop.Key())
}

func NewTablesFromType(typ types.Type, parentName TableName, hierarchy []string, propKey types.PropKey) Tables {
	switch v := typ.(type) {
	case types.List:
		listTable := NewListTable(v, parentName, hierarchy, propKey)
		tables := Tables{listTable}
		myHierarchy := append([]string{}, hierarchy...)
		myHierarchy = append(myHierarchy, string(propKey))
		return append(tables, NewTablesFromType(v.ItemsType(), listTable.Name(), myHierarchy, "")...)
	case types.Variable:
		props := v.Props()
		variableTables := NewVariableTables(v, parentName, hierarchy, propKey)
		tables := append(Tables{}, variableTables...)
		for i, prop := range props {
			tableName := variableTables[i].Name()
			myHierarchy := append([]string{}, hierarchy...)
			myHierarchy = append(myHierarchy, string(propKey))
			tables = append(tables, NewTablesFromType(prop.Type(), tableName, myHierarchy, "")...)
		}
		return tables
	case types.Group:
		return NewTablesFromProps(v.Props(), parentName, hierarchy)
	default:
	}
	return Tables{}
}

func NewVariableTables(typ types.Variable, parentName TableName, hierarchy []string, propKey types.PropKey) Tables {
	tables := Tables{}
	props := typ.Props()
	for _, prop := range props {
		columns := NewVariableColumns(prop, parentName)
		namePrefix := strings.Join(hierarchy, "_")
		name := TableName(fmt.Sprintf("%s_%s_list", namePrefix, prop.Key()))
		if propKey != "" {
			name = TableName(fmt.Sprintf("%s_%s_%s_var", namePrefix, propKey, prop.Key()))
		}
		tables = append(tables, Table{
			name:    name,
			columns: columns,
		})
	}
	return tables
}

// ToCreateSQL テーブルを生成する SQL を構築します
func (t Tables) ToCreateSQL() []string {
	sqls := make([]string, len(t))
	for i, v := range t {
		sqls[i] = v.ToCreateSQL()
	}
	return sqls
}

// Find スキーマを探す
func (t Tables) find(name TableName) (Table, bool) {
	for _, schema := range t {
		if name == schema.name {
			return schema, true
		}
	}
	return Table{}, false
}

// Table テーブル情報です
type Table struct {
	name    TableName
	columns Columns
}

// NewMainTable 主テーブルを生成します
func NewMainTable(model model.Model) Table {
	props := model.Props()
	columns := NewPropsColumns(props, "")
	name := TableName(model.Key())
	return Table{
		name:    name,
		columns: columns,
	}
}

// NewListTable リスト用のテーブルを生成します
func NewListTable(listType types.List, parentName TableName, hierarchy []string, propName types.PropKey) Table {
	columns := NewListColumns(listType, parentName)
	namePrefix := strings.Join(hierarchy, "_")
	name := TableName(fmt.Sprintf("%s_list", namePrefix))
	if propName != "" {
		name = TableName(fmt.Sprintf("%s_%s_list", namePrefix, propName))
	}
	return Table{
		name:    name,
		columns: columns,
	}
}

// Name テーブル名を取得します
func (t Table) Name() TableName {
	return t.name
}

func (t Table) Columns() Columns {
	return t.columns
}

// ColumnNames カラム名を取得します
func (t Table) ColumnNames() ColumnNames {
	return t.Columns().Names()
}

// ToCreateSQL create table 文を生成します
func (t Table) ToCreateSQL() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("CREATE TABLE `%s` (\n", t.name))
	columnDefs := make([]string, len(t.columns))
	// カラム作成
	for i, column := range t.columns {
		size := string(column.typ)
		if column.size > 0 {
			size = fmt.Sprintf("%s(%d)", column.typ, column.size)
		}
		columnDef := []string{fmt.Sprintf("\t`%s` %s", column.name, size)}
		if !column.nullable {
			columnDef = append(columnDef, "NOT NULL")
		}
		if column.primarykey {
			columnDef = append(columnDef, "PRIMARY KEY")
		}
		columnDefs[i] = strings.Join(columnDef, " ")
	}
	// インデックス生成
	index := []string{}
	unique := []string{}
	for _, column := range t.columns {
		if column.primarykey {
			continue
		}
		if column.index {
			index = append(index, fmt.Sprintf("`%s`", column.name))
		}
		if column.uniquekey {
			unique = append(index, fmt.Sprintf("`%s`", column.name))
		}
	}
	if len(index) > 0 {
		columnDefs = append(columnDefs, strings.Join(index, ", "))
	}
	if len(unique) > 0 {
		columnDefs = append(columnDefs, strings.Join(unique, ", "))
	}
	buf.WriteString(strings.Join(columnDefs, ",\n"))
	buf.WriteString("\n")
	buf.WriteString(") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")
	return buf.String()
}

// TableName テーブル名
type TableName string

// Esc エスケープされたテーブル名を取得します
func (t TableName) Esc() string {
	return fmt.Sprintf("`%s`", t)
}

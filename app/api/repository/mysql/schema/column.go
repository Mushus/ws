package schema

import (
	"fmt"
	"log"

	"github.com/Mushus/app/api/value/types"
)

// Columns テーブルのカラム
type Columns []Column

// NewListColumns リストテーブル用のカラムを作成します
func NewListColumns(listType types.List, parentName TableName) Columns {
	columns := Columns{
		{
			name:       "id",
			typ:        TypeID,
			size:       SizeID,
			nullable:   false,
			primarykey: true,
		},
		// 親テーブルへの紐づけカラム
		{
			name:     ColumnName(parentName + "_id"),
			typ:      TypeID,
			size:     SizeID,
			nullable: false,
			foeignKey: FoeignKey{
				table:  parentName,
				column: "id",
			},
		},
		// ソート用のカラム
		{
			name:     "order",
			typ:      TypePropOrder,
			size:     SizeUnset,
			nullable: false,
		},
	}
	return append(columns, NewTypeColumns(listType.ItemsType(), "", 0, false)...)
}

// NewVariableColumns いろいろテーブル用のカラムを作成します
func NewVariableColumns(prop types.Prop, parentName TableName) Columns {
	columns := Columns{
		{
			name:       "id",
			typ:        TypeID,
			size:       SizeID,
			nullable:   false,
			primarykey: true,
		},
		// 親テーブルへの紐づけカラム
		{
			name:     ColumnName(parentName + "_id"),
			typ:      TypeID,
			size:     SizeID,
			nullable: false,
			foeignKey: FoeignKey{
				table:  parentName,
				column: "id",
			},
		},
		// ソート用のカラム
		{
			name:     "order",
			typ:      TypePropOrder,
			size:     SizeUnset,
			nullable: false,
		},
	}
	return append(columns, NewTypeColumns(prop.Type(), "", 0, false)...)
}

// NewPropsColumns はプロパティのリストをスキーマに変更します
func NewPropsColumns(props types.Props, prefix string) Columns {
	columns := Columns{}
	for _, prop := range props {
		columns = append(columns, NewPropColumns(prop, prefix)...)
	}
	return columns
}

// NewPropColumns プロパティからスキーマを定義します
func NewPropColumns(prop types.Prop, prefix string) Columns {
	name := prefix + string(prop.Key())
	return NewTypeColumns(prop.Type(), name, size(prop.MaxSize()), prop.Nullable())
}

// NewTypeColumns 型情報からスキーマを定義します
// 名前 name が空文字 "" だった場合、デフォルト値として "value" が採用されます
func NewTypeColumns(typ types.Type, name string, size size, nullable bool) Columns {
	log.Println(typ.Key(), name)
	switch v := typ.(type) {
	case types.ID:
		return Columns{
			{
				name:       ColumnName(name),
				typ:        getSchemaType(v, size),
				size:       SizeID,
				nullable:   false,
				primarykey: true,
			},
		}
	case types.List:
		return Columns{}
	case types.Variable:
		// NOTE: 名前がなくなってしまうので value がデフォルト値として使用される
		if name == "" {
			name = "value"
		}

		return Columns{
			{
				name:     ColumnName(name),
				typ:      TypeVarchar,
				size:     255,
				nullable: nullable,
			},
		}
	case types.Group:
		// NOTE: Groupは名前が空でもプロパティ名が存在するので空になこことはない
		// なのでprefixをつけないことで対応
		prefix := ""
		if name != "" {
			prefix = name + "_"
		}
		return NewPropsColumns(v.Props(), prefix)
	default:
		// NOTE: 名前がなくなってしまうので value がデフォルト値として使用される
		if name == "" {
			name = "value"
		}

		return Columns{
			{
				name:     ColumnName(name),
				typ:      getSchemaType(v, size),
				size:     size,
				nullable: nullable,
			},
		}
	}
}

func (c Columns) Names() ColumnNames {
	names := make(ColumnNames, len(c))
	for i, column := range c {
		names[i] = column.Name()
	}
	return names
}

func (c Columns) IndexOf(name ColumnName) (int, error) {
	for i, column := range c {
		if column.Name() == name {
			return i, nil
		}
	}
	return -1, fmt.Errorf("column %q not found", name)
}

// スキーマのタイプを選ぶ
func getSchemaType(typ types.Type, size size) typ {
	switch typ.(type) {
	case types.String:
		switch {
		case size <= SizeVarcharMax:
			return TypeVarchar
		case size <= SizeMediumtextMax:
			return TypeMediumtext
		default:
			return TypeLongtext
		}
	case types.Integer:
		return TypeBigint
	default:
		return TypeDefault
	}
}

// Column カラムのスキーマ情報です
type Column struct {
	name       ColumnName
	typ        typ
	size       size
	nullable   bool
	primarykey bool
	uniquekey  bool
	index      bool
	foeignKey  FoeignKey
}

// Name カラム名を取得します
func (c Column) Name() ColumnName {
	return c.name
}

// ColumnName カラム名
type ColumnName string

// Esc カラム名をエスケープします
func (c ColumnName) Esc() string {
	return fmt.Sprintf("`%s`", c)
}

// ColumnNames カラム名リスト
type ColumnNames []ColumnName

// Esc カラム名リストをエスケープします
func (c ColumnNames) Esc() []string {
	list := make([]string, len(c))
	for i, name := range c {
		list[i] = name.Esc()
	}
	return list
}

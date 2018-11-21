package record

import "github.com/Mushus/app/api/repository/mysql/builder/schema"

// Values エンティティの値を表します
type Values map[schema.ColumnName]Value

// SortedValue ソートされた値です
type SortedValue []Value

// List リストを返します
func (s SortedValue) List() []interface{} {
	list := []interface{}{}
	for i, v := range s {
		list[i] = v
	}
	return list
}

// Sorted 明示的にソートしたことを表す
func Sorted(list []interface{}) SortedValue {
	values := make(SortedValue, len(list))
	for i, v := range list {
		values[i] = v
	}
	return values
}

// Value データベースのレコードのとある値を表します
type Value interface{}

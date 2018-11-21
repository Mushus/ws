package schema

// FoeignKey 外部キー
type FoeignKey struct {
	use    bool
	table  TableName
	column ColumnName
}

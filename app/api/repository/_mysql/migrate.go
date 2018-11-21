package mysql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Mushus/app/api/repository/mysql/builder/schema"
)

const (
	// SQLShowTables テーブル一覧を取得
	SQLShowTables = `SHOW TABLES`
	// SQLShowColumns カラム情報を取得
	// パラメータ1 はデータベース名です
	SQLShowColumns = `SELECT
	COLUMN_NAME AS ` + "`column`" + `,
	DATA_TYPE AS type,
	IS_NULLABLE = 'YES' AS nullable,
	COLUMN_DEFAULT AS ` + "`default`" + `,
	IFNULL(CHARACTER_MAXIMUM_LENGTH, 0) AS size
	FROM COLUMNS
	WHERE TABLE_SCHEMA = ?`
	// SQLShowPrimaryKeys 主キー一覧を取得します
	// パラメータ1 はデータベース名です
	SQLShowPrimaryKeys = `SELECT
	TABLE_NAME` + "`table`" + `,
	COLUMN_NAME` + "`column`" + `
	FROM information_schema.TABLE_CONSTRAINTS t
	JOIN information_schema.KEY_COLUMN_USAGE k
	USING(CONSTRAINT_NAME, TABLE_SCHEMA, TABLE_NAME)
	WHERE CONSTRAINT_TYPE = 'PRIMARY KEY'
	AND t.table_schema = ?`
	// SQLShowKeys キー一覧を取得します
	// パラメータ1 はデータベース名です
	SQLShowKeys = `SELECT
	TABLE_NAME AS ` + "`table`" + `,
	COLUMN_NAME AS ` + "`column`" + `,
	!NON_UNIQUE AS ` + "`unique`" + `
	FROM INFORMATION_SCHEMA.STATISTICS
	WHERE TABLE_SCHEMA = ?`
	// SQLShowForeignKeys 外部キー一覧を取得します
	// パラメータ1 はデータベース名です
	SQLShowForeignKeys = `SELECT
	TABLE_NAME AS ` + "`table`" + `,
	COLUMN_NAME as ` + "`column`" + `,
	REFERENCED_TABLE_NAME AS ref_table,
	REFERENCED_COLUMN_NAME AS ref_column
	FROM information_schema.TABLE_CONSTRAINTS t
	JOIN information_schema.KEY_COLUMN_USAGE k
	USING(CONSTRAINT_NAME, TABLE_SCHEMA, TABLE_NAME)
	WHERE CONSTRAINT_TYPE = 'FOREIGN KEY'
	AND t.table_schema = ?`
)

// ColumnSchemaRecords DBから取得したカラムのスキーマ一覧です
type ColumnSchemaRecords []ColumnSchemaRecord

// Norm スキーマを正規化して内部で用いる形式に変換します
func (c ColumnSchemaRecords) Norm() schema.Columns {
	schemas := make(schema.Columns, len(s))
	for i, v := range s {
		schemas[i] = v.Norm()
	}
	return schemas
}

// ColumnSchemaRecord DBから取得したカラムのスキーマです
type ColumnSchemaRecord struct {
	// フィールド名
	Field string `db:"Field"`
	// BIGINT, DOUBLE, VARCHAR, MEDIUMTEXT, LONGTEXT
	Type string `db:"Type"`
	// YES | NO
	Null string `db:"Null"`
	// PRI: プライマリーキー UNI: ユニークキー MUL: インデックス
	Key string `db:"Key"`
	// NULL
	Default interface{} `db:"Default"`
	// ?
	Extra string `db:"Extra"`
}

// TypeName 型名を取得します
func (s schemaRecord) TypeName() string {
	typ := s.Type
	// NOTE: 型名は `int(11)` みたいな形で入っている
	pos := strings.Index(typ, "(")
	if pos < 0 {
		return s.Type
	}
	return typ[:pos]
}

// Size 型のサイズを取得します
func (s schemaRecord) Size() int {
	typ := s.Type
	// NOTE: 型名は `int(11)` みたいな形で入っている
	startPos := strings.Index(typ, "(")
	if startPos < 0 {
		return 0
	}
	// NOTE: index + 1 がカッコのあとの位置
	startPos++
	endPos := strings.LastIndex(typ, ")")
	if endPos < 0 || !(startPos < endPos) {
		return 0
	}
	sizeStr := typ[startPos:endPos]
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return 0
	}
	return size
}

func (s schemaRecord) IsNull() bool {
	// NOTE: 取りうる値は YES or NO
	return s.Null == "YES"
}

func (s schemaRecord) Norm() Schema.Column {
	return Schema.Column{
		name:     s.Field,
		typ:      s.TypeName(),
		size:     s.Size(),
		nullable: s.IsNull(),
	}
}

// GatherAllSchemas 古いスキーマを集めます
func (d DataStore) GatherAllSchemas() (schemas.Tables, error) {
	tableNames, err := d.GetAllTableName()
	if err != nil {
		return nil, fmt.Errorf("cannot get table names: %v", err)
	}

	tables := make(schema.Tables, len(tableNames))
	for i, name := range tableNames {
		s, err := d.getSchema(name)
		if err != nil {
			return nil, fmt.Errorf("cannot get table schema: %v", err)
		}
		tables[i] = &schema.Table{
			name:    name,
			columns: s,
		}
	}
	return tables, nil
}

// GetColumnSchemas はスキーマをデータベースから取得します
func (d DataStore) GetColumnSchemas(tableName tableName) (schemas.Columns, error) {
	schemas := schemaRecords{}
	if err := d.db.Select(&schemas, ``, tableName); err != nil {
		return nil, fmt.Errorf("failed to get table shema: %v", err)
	}
	return schemas.Norm(), nil
}

// GetAllTableNames はテーブル名を取得します
// エラー err が nil のときは必ず list は nil ではありません。
func (d DataStore) GetAllTableNames() (list []schema.TableName, err error) {
	rows, err := d.db.Query(SQLShowTables)
	if err != nil {
		return nil, err
	}
	list = []tableName{}
	for rows.Next() {
		var tableName tableName
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		list = append(list, tableName)
	}
	return list, nil
}

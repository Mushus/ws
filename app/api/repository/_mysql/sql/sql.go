package sql

import (
	"fmt"

	"github.com/Mushus/app/api/repository/mysql/builder/record"
	"github.com/Mushus/app/api/repository/mysql/builder/schema"
	"github.com/jmoiron/sqlx"
	sq "gopkg.in/Masterminds/squirrel.v1"
)

type TableType int

const (
	TypeUnknown TableType = iota
	TypeMain
	TypeList
)

type TableInfo struct {
	table     *schema.Table
	tableType TableType
	Columns   map[schema.ColumnName]*ColumnInfo
	Children  []*TableInfo
}

func (t TableInfo) Table() *schema.Table {
	return t.table
}

type ColumnInfo struct {
	hierarchy []string
}

type ModelKey string

type ConversionTable map[ModelKey]*TableInfo

func (c ConversionTable) GetTableInfo(key ModelKey) *TableInfo {
	if v, ok := c[key]; ok {
		return v
	}
	return nil
}

type Converter struct {
	table ConversionTable
	db    *sqlx.DB
}

func NewConverter(db *sqlx.DB) Converter {
	return Converter{
		table: ConversionTable{},
		db:    nil, // TODO:
	}
}

func (c Converter) Restore(key ModelKey, ids []string) {
	tableInfo := c.table.GetTableInfo(key)
	c.RestoreRecords(tableInfo, ids)
}

func (c Converter) RestoreRecords(tableInfo *TableInfo, ids []string) (*HierarchyTemporary, error) {
	table := tableInfo.Table()
	tableName := table.Name()
	sql, args, err := sq.
		Select(table.ColumnNames().Esc()...).
		From(tableName.Esc()).
		Where(sq.Eq{"id": ids}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to create restore sql %q: %v", tableName, err)
	}

	rows, err := c.db.Queryx(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to create restore %q: %v", tableName, err)
	}

	records := record.Records{}
	for rows.Next() {
		cols, err := rows.SliceScan()
		if err != nil {
			fmt.Errorf("failed to scan record %q: %v", tableName, err)
		}
		sorted := record.Sorted(cols)
		records = append(records, sorted)
	}
	// TODO:
	tmp := record.NewTemporary(table, records)

	tmpIds, err := tmp.IDs()
	if err != nil {
		return nil, fmt.Errorf("failed to get IDs: %v", err)
	}
	hchildren := make([]*HierarchyTemporary, len(tableInfo.Children))
	for i, child := range tableInfo.Children {
		hchild, err := c.RestoreRecords(child, tmpIds)
		if err != nil {
			return nil, err
		}
		hchildren[i] = hchild
	}
	return &HierarchyTemporary{
		temporary: tmp,
		children:  hchildren,
	}, nil
}

type HierarchyTemporary struct {
	temporary record.Temporary
	children  []*HierarchyTemporary
}

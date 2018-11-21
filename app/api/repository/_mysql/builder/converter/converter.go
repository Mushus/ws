package converter

import (
	"github.com/Mushus/app/api/repository/mysql/builder/record"
	"github.com/Mushus/app/api/repository/mysql/builder/schema"
	"github.com/Mushus/app/api/value/model"
	sq "gopkg.in/Masterminds/squirrel.v1"
)

// 分割数
const concurrencyNum = 100

type Converter struct {
	models  model.Models
	schemas schema.Tables
}

func NewConverter(models model.Models) Converter {
	return Converter{
		models:  models,
		schemas: schema.NewTablesFromModels(models),
	}
}

// ToRecord レコード一覧を作成します
// func (c Converter) ToRecord(values ...model.Value) record.LabeledTemporary {
// 	tmp := []record.Temporary{}
// 	for _, v := range values {
// 		tmp = append(tmp, c.toRecord(v)...)
// 	}
// 	return record.LabelTemporary(tmp)
// }

// func (c Converter) toRecord(typed model.Value) []record.Temporary {
// 	props := typed.Value()
// 	model := typed.Model()
// 	return []record.Temporary{}
// }

// ToInsertSQLs Insert文を作成する
func (c Converter) ToInsertSQLs(labeld record.LabeledTemporary) ([]Query, error) {
	prepare := []Query{}
	for _, table := range c.schemas {
		tableName := table.Name()
		tableRecords, found := labeld[tableName]
		if !found {
			continue
		}

		columnNames := table.ColumnNames()
		stmt := sq.Insert(tableName.Esc()).Columns(columnNames.Esc()...)

		records := tableRecords.Records()
		valueLength := len(records)
		for i := 0; i < valueLength; i += concurrencyNum {
			last := i + concurrencyNum
			if last > valueLength {
				last = valueLength
			}
			// 指定件数ごとに区切る
			for _, record := range records[i:last] {
				stmt = stmt.Values(record.List()...)
			}

			// SQLに変換する
			sql, args, err := stmt.ToSql()
			if err != nil {
				return nil, err
			}
			prepare = append(prepare, Query{
				sql:  sql,
				args: args,
			})
		}
	}
	return prepare, nil
}

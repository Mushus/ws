package record

import (
	"errors"

	"github.com/Mushus/app/api/repository/mysql/builder/schema"
)

// 一時テーブルです
// 挿入前のレコードをプールします
type Temporary struct {
	table   *schema.Table
	records Records
}

func NewTemporary(table *schema.Table, records Records) Temporary {
	return Temporary{
		table:   table,
		records: records,
	}
}

// Add 一時テーブルを統合します
func (t *Temporary) Add(dist Temporary) error {
	// NOTE: 同一のテーブルのレコードしか持てません
	if dist.table != t.table {
		return errors.New("cannot combine temporaries")
	}
	t.records = append(t.records, dist.records...)
	return nil
}

// Table テーブルを取得します
func (t Temporary) Table() *schema.Table {
	return t.table
}

// Records レコード一覧を取得します
func (t Temporary) Records() Records {
	return t.records
}

// IDs ID列を取得します
func (t Temporary) IDs() ([]string, error) {
	idIndex, err := t.Table().Columns().IndexOf("id")
	if err != nil {
		return nil, err
	}
	records := t.Records()
	values := make([]string, len(records))
	for i, value := range records {
		if len(value) <= idIndex {
			return nil, errors.New("index out of range")
		}
		id, ok := value[idIndex].(string)
		if !ok {
			return nil, errors.New("cast error")
		}
		values[i] = id
	}
	return values, nil
}

// LabeledTemporary 分類された一時テーブル
type LabeledTemporary map[schema.TableName]Temporary

// LabelTemporary 一時テーブルリストを分類します
func LabelTemporary(tmps []Temporary) LabeledTemporary {
	lt := LabeledTemporary{}
	for _, tmp := range tmps {
		table := tmp.Table()
		key := table.Name()
		orig, exists := lt[key]
		if !exists {
			orig = NewTemporary(table, []SortedValue{})
		}
		// NOTE: 事前条件に合致しているのでエラー無視
		orig.Add(tmp)
		lt[key] = orig
	}
	return lt
}

package converter_test

import (
	"reflect"
	"testing"

	"github.com/Mushus/app/api/repository/mysql/builder/schema"
)

func TestDividedValue(t *testing.T) {
	ts := schema.Tables{
		schema.NewMainTable{
			name: "table1",
			columns: []schema{
				{name: "hoge"},
				{name: "fuga"},
			},
		},
		{
			name: "table2",
			columns: []schema{
				{name: "piyo"},
			},
		},
	}

	dividedValues := LabeledTemporary{
		"table1": []entityValue{
			{"hoge": "hoge", "fuga": "fuga"},
			{"hoge": "hogehoge", "fuga": "fugafuga"},
		},
		"table2": []entityValue{
			{"piyo": "piyo"},
		},
	}

	want := []Query{
		{
			sql:  "INSERT INTO `table1` (`hoge`,`fuga`) VALUES (?,?),(?,?)",
			args: []interface{}{"hoge", "fuga", "hogehoge", "fugafuga"},
		},
		{
			sql:  "INSERT INTO `table2` (`piyo`) VALUES (?)",
			args: []interface{}{"piyo"},
		},
	}

	got, err := dividedValues.toInsertSQLs(ts)
	if err != nil {
		t.Fatalf("dividedValues.toInsertSQLs() must not occer error: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("dividedValues.toInsertSQLs() = %#v, want %#v", got, want)
	}

}

package converter

import (
	"reflect"
	"testing"

	"github.com/Mushus/app/api/value/model"
	"github.com/Mushus/app/api/value/types"
	"github.com/Mushus/app/api/value/validation"
)

var nop = validation.Option{}

func TestToModelValue(t *testing.T) {
	model := model.New("top", types.Props{
		types.NewProp("title", types.String{}, false, nop),
		types.NewProp("sentence", types.String{}, false, nop),
		types.NewProp("main", types.NewList(types.String{}, nop), false, nop),
	})

	value, err := model.Cast(map[string]interface{}{
		"id":       "hogehoge",
		"title":    "hoge",
		"sentence": "hogehoge",
		"main":     []interface{}{"hoge", "fuga", "piyo"},
	})
	if err != nil {
		t.Fatalf("invalid test code: failed to cast: %v", err)
	}

	got := toEntities(value)
	want := []entity{
		entity{
			tableName: "top_main_prop",
			value: entityValue{
				"id":     got[0].value["id"],
				"top_id": "hogehoge",
				"order":  0,
				"value":  "hoge",
			},
		},
		entity{
			tableName: "top_main_prop",
			value: entityValue{
				"id":     got[1].value["id"],
				"top_id": "hogehoge",
				"order":  1,
				"value":  "fuga",
			},
		},
		entity{
			tableName: "top_main_prop",
			value: entityValue{
				"id":     got[2].value["id"],
				"top_id": "hogehoge",
				"order":  2,
				"value":  "piyo",
			},
		},
		entity{
			tableName: "top",
			value: entityValue{
				"id":       "hogehoge",
				"title":    "hoge",
				"sentence": "hogehoge",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("toEntities(value) = %#v, want %#v", got, want)
	}
}

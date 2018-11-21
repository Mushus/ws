package schema_test

import (
	"testing"
	"github.com/Mushus/app/api/value/model"
	"github.com/Mushus/app/api/value/types"
	"github.com/Mushus/app/api/value/validation"
	"github.com/Mushus/app/api/repository/mysql/schema"
	//"github.com/davecgh/go-spew/spew"
)

var nop = validation.Option{}
var models = model.Models{
	model.New("top", types.Props{
		types.NewProp("title", types.String{}, false, nop),
		types.NewProp("sentence", types.String{}, false, nop),
		types.NewProp("main", types.NewList(types.String{}, nop), false, nop),
		types.NewProp("main2",
			types.NewList(
				types.NewList(
					types.NewGroup(types.Props{
						types.NewProp("hoge", types.String{}, false, nop),
						types.NewProp("fuga", types.String{}, false, nop),
					}, nop), nop,
				), nop,
			), false, nop,
		),
		types.NewProp("main3",
			types.NewList(
				types.NewGroup(types.Props{
					types.NewProp("piyo", types.Integer{}, false, nop),
					types.NewProp("hoge",
						types.NewList(
							types.NewGroup(types.Props{
								types.NewProp("hoge", types.String{}, false, nop),
								types.NewProp("fuga", types.String{}, false, nop),
							}, nop), nop,
						), false, nop,
					),
				}, nop), nop,
			),false, nop,
		),
		types.NewProp("main4",
			types.NewList(
				types.NewVariable(types.Props{
					types.NewProp("piyo", types.Integer{}, false, nop),
					types.NewProp("hoge", types.String{}, false, nop),
					types.NewProp("hoge",
						types.NewGroup(types.Props{
							types.NewProp("hoge", types.String{}, false, nop),
							types.NewProp("fuga", types.String{}, false, nop),
						}, nop), false, nop,
					),
				}, nop), nop,
			),false, nop,
		),
	}),
}

func TestCreatTables(t *testing.T) {
	s := schema.NewTablesFromModels(models)
	t.Logf("%#v", s)
}

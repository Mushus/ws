package value

import (
	"github.com/Mushus/app/api/config"
	"github.com/Mushus/app/api/value/model"
	"github.com/Mushus/app/api/value/types"
	"github.com/Mushus/app/api/value/validation"
)

// CreateEntities はエンティティの内部的な定義を設定ファイルから作ります
func CreateEntities(cfg config.Config) (model.Models, error) {
	models := model.Models{}
	typs, err := types.CreateList(cfg.Types)
	if err != nil {
		return models, err
	}

	for _, ec := range cfg.Models {
		props := types.Props{}
		for pkey, p := range ec.Props {
			typ, err := typs.Get(p.Type)
			if err != nil {
				return models, err
			}
			// TODO: nullable, optionを設定する
			prop := types.NewProp(pkey, typ, false, validation.Option{})
			props = append(props, prop)
		}
		models.Add(model.New(model.Key(ec.Key), props))
	}
	return models, nil
}

package config

// Find は Models 一覧の中から key に対応する model を探します
// もし model が見つからなかった場合は ok が false を返します
func (e Models) Find(key string) (model Model, ok bool) {
	for _, v := range e {
		if v.Key == key {
			return v, true
		}
	}
	return model, false
}

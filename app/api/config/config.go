package config

type (
	// Config は設定ファイルです
	Config struct {
		// アプリケーションに関する設定
		App App

		// 型の定義
		Types Types
		// エンティティの定義
		Models Models
	}

	// App はアプリケーションに関する設定です
	App struct {
		Port int
	}
	// Types 型定義一覧です
	Types []Type
	// Type 型定義です
	Type struct {
		Key        string
		Base       string
		Validation Validation
	}
	// Models エンティティ一覧
	Models []Model
	// Model は実体を定義する設定です
	Model struct {
		Key   string
		Props map[string]Property
	}

	// Property は Model が持っているプロパティを定義する設定です
	Property struct {
		Type       string
		Validation Validation
	}

	// Validation は Property に対する制約の設定です
	Validation struct {
	}
)

// LoadConfig は設定ファイルを読み込みます
func LoadConfig() Config {
	// TODO: ファイルから読み込む実装をする
	return Config{
		App: App{
			Port: 8080,
		},
		Models: Models{
			{
				Key: "article",
				Props: map[string]Property{
					"title": {
						Type: "text",
					},
				},
			},
		},
	}
}

// Check は設定ファイルがただしいかバリデーションを行います
// 問題がないとき、返り値は `nil` になります
// 問題があるとき、戻り値は対応したエラーになります
func Check(cfg Config) error {
	if cfg.Models == nil {
		return ErrEntitiesIsRequired
	}

	return nil
}

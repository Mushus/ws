package config

type (
	// Config はアプリケーションの設定です。
	Config struct {
		Server       Server
		Applications Applications
	}

	// Server はアプリケーションのサーバー設定です
	Server struct {
		Port int
	}

	// Applications アプリケーションのキーとその一覧を表します
	Applications map[string]Application
	// Application ゲートウェイにぶら下がるアプリケーションを表します
	Application struct {
		URL string
	}
)

// LoadConfig はアプリケーションの設定を読み込みます。
func LoadConfig() Config {
	return Config{
		Server: Server{
			Port: 8080,
		},
		Applications: Applications{
			"master": {
				URL: "http://localhost:8090",
			},
		},
	}
}

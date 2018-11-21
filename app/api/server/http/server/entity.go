package server

type (
	// Status サービスが正常に稼働しているかの状態です
	Status string
	// Response はヘルスチェックの結果です
	Response struct {
		// すべての結果
		Status Status `json:"status"`
		// 各サービスの結果
		Detail map[string]Status `json:"detail"`
	}
)

const (
	// StatusOK はサービスが正常なことを表します
	StatusOK = "ok"
	// StatusNG はサービスが異常なことを表します
	StatusNG = "ng"
)

package mysql

// Option データストアの設定を行います
type Option func(ds *MySQL) error

// Host は MySQL の接続先のホストを設定します
func Host(host string) Option {
	return func(ds *MySQL) error {
		ds.host = host
		return nil
	}
}

// Port は MySQL の接続先のポート番号を設定します
func Port(port int) Option {
	return func(ds *MySQL) error {
		ds.port = port
		return nil
	}
}

// User は MySQL に接続するユーザーを設定します
func User(user string) Option {
	return func(ds *MySQL) error {
		ds.user = user
		return nil
	}
}

// Password は MySQL に接続するパスワードを設定します
func Password(password string) Option {
	return func(ds *MySQL) error {
		ds.password = password
		return nil
	}
}

// Database は MySQL のデータベースを設定します
func Database(database string) Option {
	return func(ds *MySQL) error {
		ds.database = database
		return nil
	}
}

// ForceMigration は確認なしに強制的にマイグレーションを行います
func ForceMigration(ok bool) Option {
	return func(ds *MySQL) error {
		ds.forceMigration = ok
		return nil
	}
}

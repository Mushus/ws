package service

// NewCfg is
func NewCfg() CfgService {
	if cfgInstance == nil {
		cfgInstance = &cfgService{
			config: Config{
				DB: DBConfig{
					DSN: ":memory:",
				},
				Port: 8080,
			},
		}
	}
	return cfgInstance
}

var cfgInstance CfgService

// CfgService is
type CfgService interface {
	Get() Config
}

type cfgService struct {
	config Config
}

func (c *cfgService) Get() Config {
	return c.config
}

// Config is
type Config struct {
	DB   DBConfig
	Port int
}

// DBConfig is
type DBConfig struct {
	DSN string
}

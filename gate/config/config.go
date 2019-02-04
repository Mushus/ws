package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

// Load filepath から設定ファイルを読み込みます
func Load(filepath string) (*Config, error) {

	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read yaml file")
	}

	yml := yaml.MapSlice{}
	err = yaml.Unmarshal(b, &yml)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal yaml")
	}

	cfg := Config{}

	app, err := createApp(yml)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read yaml parameter")
	}
	cfg.App = *app

	services, err := createServices(yml)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read yaml parameter")
	}
	cfg.Services = services

	return &cfg, nil
}

var defaultApp = App{
	Port: 80,
}

func createApp(cfg yaml.MapSlice) (*App, error) {
	app := defaultApp
	for _, appCfg := range cfg {
		if appCfg.Key != "app" {
			continue
		}

		appCfgVal, ok := appCfg.Value.(yaml.MapSlice)
		if !ok {
			return nil, errors.Errorf("invalid parameter `app`: expect object, but found %T", appCfg.Value)
		}

		for _, appCfgCol := range appCfgVal {
			if appCfgCol.Key == "port" {
				port, ok := appCfgCol.Value.(int)
				if !ok {
					return nil, errors.Errorf("invalid parameter `app.port`: expect int, but found %T", appCfgCol.Value)
				}
				app.Port = Port(port)
			}
		}
	}
	return &app, nil
}

// createServices フォーマット済みの設定ファイルを書き出します
func createServices(cfg yaml.MapSlice) ([]Service, error) {
	services := []Service{}
	for _, serviceCfg := range cfg {
		if serviceCfg.Key != "services" {
			continue
		}

		servicesCfgVal, ok := serviceCfg.Value.(yaml.MapSlice)
		if !ok {
			return nil, errors.Errorf("invalid parameter `services`: expect object, but found %T", serviceCfg.Value)
		}

		for _, servicesCfgCol := range servicesCfgVal {
			path, ok := servicesCfgCol.Key.(string)
			if !ok {
				return nil, errors.Errorf("invalid parameter `services[key]`: expect string, but found %T", servicesCfgCol.Key)
			}

			optionsCfg, ok := servicesCfgCol.Value.(yaml.MapSlice)
			if !ok {
				return nil, errors.Errorf("invalid parameter `services.path`: expect []object, but found %T", servicesCfgCol.Key)
			}

			options := []Option{}
			for _, option := range optionsCfg {
				options = append(options, option)
			}
			services = append(services, Service{
				Path:    path,
				Options: options,
			})
		}
	}

	return services, nil
}

type Config struct {
	App         App
	Services    []Service
	Middlewares []Middrlware
}

// App アプリケーションに関する設定
type App struct {
	Port Port `yaml:"app"`
}

// Port ポート番号
type Port int

// HTTPString ポート番号を http パッケージの形式で表す
func (p Port) HTTPString() string {
	return fmt.Sprintf(":%d", p)
}

// Service セキュリティに関する設定
type Service struct {
	Path    string
	Options []Option
}

// Option 設定項目
type Option interface{}

// Middrlware セキュリティに関する設定
type Middrlware struct {
}

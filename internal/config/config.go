package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		RMQ  `yaml:"rabbitmq"`
		DB   `yaml:"db"`
	}

	App struct {
		Name     string `yaml:"name"    env-required:"true" env:"APP_NAME" env-default:"cdg"`
		Version  string `yaml:"version" env-required:"true" env:"APP_VERSION" env-default:"0.1.0"`
		DataPath string `yaml:"data-path" env-required:"true" env:"APP_DATA_PATH" env-default:"/tmp/data"`
	}

	HTTP struct {
		Port int `yaml:"port" env-required:"true" env:"PORT" env-default:"3001"`
	}

	RMQ struct {
		Uri string `yaml:"uri" env-required:"true" env:"RMQ_URI" env-default:"amqp://rabbitmq:5672"`
	}

	DB struct {
		CONNECTION string `yaml:"uri" env:"DB_CONNECTION" env-required:"true"`
	}
)

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		err = cleanenv.ReadEnv(cfg)
		if err != nil {
			panic(err)
		}
	}

	return cfg, err
}

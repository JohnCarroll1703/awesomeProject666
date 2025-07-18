package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type Config struct {
	AppPort   string `envconfig:"APP_PORT" default:":8080"`
	LogLevel  string `envconfig:"LOG_LEVEL" default:"info"`
	JWTSecret string `envconfig:"JWT_SECRET" default:"secret"`
	Databases struct {
		PostgresDSN string `envconfig:"POSTGRES_DSN"`
	}
	GrpcPort string `envconfig:"GRPC_PORT" default:"4040"`
}

var (
	cfg  *Config
	err  error
	once sync.Once
)

func GetConfig(envfiles ...string) (*Config, error) {
	once.Do(func() {
		_ = godotenv.Load(envfiles...)
		var c Config
		if err = envconfig.Process("", &c); err != nil {
			err = fmt.Errorf("error parse config from env")
			return
		}

		cfg = &c
	})

	if cfg == nil {
		return nil, errors.New("config not initialized")
	}
	return cfg, err
}

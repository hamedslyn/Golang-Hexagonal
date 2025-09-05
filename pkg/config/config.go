package config

import (
	"log"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type PostgresConfig struct {
	ConnectionString string `json:"connection_string"`
}

type AppConfig struct {
	Environment string         `json:"environment"`
	Port        string         `json:"port"`
	Postgres    PostgresConfig `json:"postgres"`
}

func Load(path string) (*AppConfig, error) {
	k := koanf.New(".")

	if err := k.Load(file.Provider(path), json.Parser()); err != nil {
		return nil, err
	}

	envPath := os.Getenv("ENV_FILE")
	if envPath == "" {
		envPath = ".env"
	}
	if _, err := os.Stat(envPath); err == nil {
		if err := k.Load(file.Provider(envPath), dotenv.Parser()); err != nil {
			return nil, err
		}
	}

	if err := k.Load(env.Provider("APP_", ".", func(s string) string {
		key := strings.TrimPrefix(s, "APP_")
		key = strings.ToLower(key)
		key = strings.ReplaceAll(key, "__", ".")
		return key
	}), nil); err != nil {
		return nil, err
	}

	cfg := &AppConfig{
		Environment: k.String("environment"),
		Port:        k.String("port"),
		Postgres: PostgresConfig{
			ConnectionString: k.String("postgres.connection_string"),
		},
	}

	return cfg, nil
}

func MustLoad() *AppConfig {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "configs/config.json"
	}

	cfg, err := Load(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}

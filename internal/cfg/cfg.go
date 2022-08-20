package cfg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

// Config contains the current selected environment's configuration.
type Config struct {
	Env              string
	DatabaseHost     string `yaml:"database_host"`
	DatabasePort     string `yaml:"database_port"`
	DatabaseName     string `yaml:"database_name"`
	DatabaseUsername string `yaml:"database_username"`
	DatabasePassword string `yaml:"database_password"`
	DatabaseURL      string `yaml:"database_url"`
	ServerPort       string `yaml:"server_port"`
	ServerCache      bool   `yaml:"server_cache"`
}

type environments struct {
	Development Config `yaml:"development"`
	Production  Config `yaml:"production"`
	Test        Config `yaml:"test"`
}

const (
	envFile     = "./config/environment.yml"
	development = "development"
	production  = "production"
	test        = "test"
)

// LoadConfig returns a new environment configuration from environment.yml
func LoadConfig(env string) (*Config, error) {
	var conf Config

	err := godotenv.Load(".env")
	if err != nil {
		return &conf, err
	}

	path, err := filepath.Abs(envFile)
	if err != nil {
		return &conf, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return &conf, err
	}

	extData := []byte(os.ExpandEnv(string(data)))

	var envs environments
	err = yaml.Unmarshal(extData, &envs)
	if err != nil {
		return &conf, err
	}

	switch env {
	case development:
		conf = envs.Development
		conf.Env = development
		conf.DatabaseURL = buildDatabaseURL(envs.Development)
		return &conf, nil
	case production:
		conf = envs.Production
		conf.Env = production
		conf.DatabaseURL = buildDatabaseURL(envs.Production)
		return &conf, nil
	case test:
		conf = envs.Test
		conf.Env = test
		conf.DatabaseURL = buildDatabaseURL(envs.Test)
		return &conf, nil
	default:
		return &conf, fmt.Errorf("%s is not a recognized environment variable", env)
	}
}

func buildDatabaseURL(c Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUsername,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
	)
}

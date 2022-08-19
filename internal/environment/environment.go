package environment

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type envFields struct {
	DatabaseHost     string `yaml:"database_host"`
	DatabasePort     string `yaml:"database_port"`
	DatabaseName     string `yaml:"database_name"`
	DatabaseUsername string `yaml:"database_username"`
	DatabasePassword string `yaml:"database_password"`
	DatabaseURL      string `yaml:"database_url"`
	ServerPort       string `yaml:"server_port"`
	ServerCache      bool   `yaml:"server_cache"`
}

type Config struct {
	CurrentEnv  string
	Development envFields `yaml:"development"`
	Production  envFields `yaml:"production"`
	Test        envFields `yaml:"test"`
}

const (
	envFile     = "./config/environment.yml"
	development = "development"
	production  = "production"
	test        = "test"
)

var conf = Config{}

// New returns a new environment configuration from environment.yml
func New(env string) (*Config, error) {
	cEnv, err := currentEnv(env)
	if err != nil {
		return &conf, err
	}
	conf.CurrentEnv = cEnv

	err = godotenv.Load(".env")
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

	err = yaml.Unmarshal(extData, &conf)
	if err != nil {
		return &conf, err
	}

	conf.Development.DatabaseURL = buildDBURL(conf.Development)
	conf.Test.DatabaseURL = buildDBURL(conf.Test)

	return &conf, nil
}

func currentEnv(env string) (string, error) {
	switch env {
	case development:
		return development, nil
	case production:
		return production, nil
	case test:
		return production, nil
	default:
		return "", fmt.Errorf("%s is not a recognized environment variable", env)
	}
}

func buildDBURL(ef envFields) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		ef.DatabaseUsername,
		ef.DatabasePassword,
		ef.DatabaseHost,
		ef.DatabasePort,
		ef.DatabaseName,
	)
}

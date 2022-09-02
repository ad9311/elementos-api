package cfg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/pelletier/go-toml"
)

// Config ...
type Config struct {
	DatabaseURL string `toml:"database_url"`
	ServerPort  string `toml:"server_port"`
	SeverSecure bool   `toml:"server_secure"`
	ServerCache bool   `toml:"server_cache"`
}

type environments struct {
	Development Config `toml:"development"`
	Production  Config `toml:"production"`
	Test        Config `toml:"test"`
}

const (
	envFile     = "./environments.toml"
	development = "development"
	production  = "production"
	test        = "test"
)

// LoadConfig ...
func LoadConfig(env string) (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, err
	}

	path, err := filepath.Abs(envFile)
	if err != nil {
		return Config{}, err
	}

	rawData, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	data := []byte(os.ExpandEnv(string(rawData)))

	var envs environments
	toml.Unmarshal(data, &envs)

	switch env {
	case development:
		return envs.Development, nil
	case production:
		return envs.Production, nil
	case test:
		return envs.Test, nil
	default:
		return Config{}, fmt.Errorf("%s is not a recognized environment variable", env)
	}
}

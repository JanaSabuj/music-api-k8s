package config

import "os"

type Config struct {
	DbHost string `yaml:"DbHost"`
	DbName string `yaml:"DbName"`
	DbPass string `yaml:"DbPass"`
}

var (
	dbHost = "localhost:3306"
	dbName = "evergreen_music_db"
	dbPass = "green"
)

func NewConfig() (*Config, error) {
	var err error
	cfg := &Config{
		DbHost: dbHost,
		DbName: dbName,
		DbPass: dbPass,
	}

	// update config values from env, if any
	cfg.GETENVs()

	return cfg, err
}

func (c *Config) GETENVs() {
	if val, found := os.LookupEnv("CONFIG_DBHOST"); found {
		c.DbHost = val
	}

	if val, found := os.LookupEnv("CONFIG_DBNAME"); found {
		c.DbName = val
	}

	if val, found := os.LookupEnv("CONFIG_DBPASS"); found {
		c.DbPass = val
	}
}

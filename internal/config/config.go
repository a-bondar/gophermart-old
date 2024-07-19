package config

import (
	"flag"
	"os"
)

type Config struct {
	RunAddr     string
	DatabaseURI string
}

func NewConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.RunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&config.DatabaseURI, "d", "", "database URI")
	flag.Parse()

	if envRunAddr, ok := os.LookupEnv("RUN_ADDRESS"); ok {
		config.RunAddr = envRunAddr
	}

	if envDatabaseURI, ok := os.LookupEnv("DATABASE_URI"); ok {
		config.DatabaseURI = envDatabaseURI
	}

	return config
}

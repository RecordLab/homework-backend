package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Mongo  MongoConfig
}

var DefaultConfig = Config{
	Server: DefaultServerConfig,
	Mongo:  DefaultMongoConfig,
}

type ServerConfig struct {
	BindAddr string `mapstructure:"bind_addr"`
	Secret   string
}

var DefaultServerConfig = ServerConfig{
	BindAddr: ":8080",
}

type MongoConfig struct {
	URL      string
	Database string
}

var DefaultMongoConfig = MongoConfig{
	URL:      "mongodb://localhost",
	Database: "dailyscoop",
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("dailyscoop")
	viper.AddConfigPath(".")
	cfg := DefaultConfig
	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return cfg, nil
		}
		return Config{}, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

package config

import (
	"nguyendangviet12022004/url-short/user/pkg/logging"
	"os"

	"github.com/spf13/viper"
)

type Db struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string `mapstructure:"db_name"`
}

type Jwt struct {
	PrivateKeyPath string `mapstructure:"private_key_path"`
}

type Server struct {
	Host string
	Port string
}

type Config struct {
	Db     Db
	Jwt    Jwt
	Server Server
}

func LoadConfig() (*Config, error) {

	// config file type and name
	viper.SetConfigType("yaml")
	viper.SetConfigName("user-config")

	// config file path
	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		viper.AddConfigPath(configPath)
	}
	viper.AddConfigPath("/etc/url-short/user/config")
	viper.AddConfigPath("./../config")

	// read config file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// unmarshal config file
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	// logging
	logger := logging.GetLogger()
	logger.Info("Using config file: " + viper.ConfigFileUsed())
	return &cfg, nil
}

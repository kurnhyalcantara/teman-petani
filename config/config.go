package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`

	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	DbSslMode  string `mapstructure:"DB_SSL_MODE"`
	DbTimeZone string `mapstructure:"DB_TIMEZONE"`
	DbRetry    string `mapstructure:"DB_RETRY"`
	DbTimeout  string `mapstructure:"DB_TIMEOUT"`

	DbMaxIdle string `mapstructure:"DB_MAX_IDLE_CONNS"`
	DbMaxOpen string `mapstructure:"DB_MAX_OPEN_CONNS"`

	LoggerOutput  string `mapstructure:"LOGGER_OUTPUT"`
	FluentBitHost string `mapstructure:"FLUENTBIT_HOST"`
	FluentBitPort string `mapstructure:"FLUENTBIT_PORT"`
}

func InitConfig() *Config {
	var config *Config

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("unable to load config: ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("unable to unmarshal config: ", err)
	}

	return config
}

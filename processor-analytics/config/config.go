package config

import (
	"github.com/spf13/viper"
	"github.com/stakkato95/lambda-architecture/processor-analytics/logger"
)

type Config struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	KafkaService string `mapstructure:"KAFKA_SERVICE"`
}

var AppConfig Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("config not found")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		logger.Panic("config can not be read")
	}

	if AppConfig == (Config{}) {
		logger.Panic("config is emtpy")
	}
}

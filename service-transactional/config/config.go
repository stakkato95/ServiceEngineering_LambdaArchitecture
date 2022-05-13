package config

import (
	"github.com/spf13/viper"
	"github.com/stakkato95/lambda-architecture/service-transactional/logger"
)

type Config struct {
	KafkaService string `mapstructure:"KAFKA_SERVICE"`
	KafkaTopic   string `mapstructure:"KAFKA_TOPIC"`
	DbDriver     string `mapstructure:"DB_DRIVER"`
	DbSource     string `mapstructure:"DB_SOURCE"`
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

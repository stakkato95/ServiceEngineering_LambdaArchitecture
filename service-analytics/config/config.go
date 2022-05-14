package config

import (
	"github.com/spf13/viper"
	"github.com/stakkato95/lambda-architecture/service-analytics/logger"
)

type Config struct {
	ServerPort        string `mapstructure:"SERVER_PORT"`
	CassandraCluster  string `mapstructure:"CASSANDRA_CLUSTER"`
	CassandraUser     string `mapstructure:"CASSANDRA_USER"`
	CassandraPassword string `mapstructure:"CASSANDRA_PASSWORD"`
	CassandraKeyspace string `mapstructure:"CASSANDRA_KEYSPACE"`
	CassandraTable    string `mapstructure:"CASSANDRA_TABLE"`
	DbDriver          string `mapstructure:"DB_DRIVER"`
	DbSource          string `mapstructure:"DB_SOURCE"`
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

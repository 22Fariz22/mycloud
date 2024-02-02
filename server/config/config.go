package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath("./server/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

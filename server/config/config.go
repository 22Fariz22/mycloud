package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() error {
	fmt.Println("in init")
	viper.AddConfigPath("./server/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

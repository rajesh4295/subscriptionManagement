package util

import (
	"errors"

	"github.com/spf13/viper"
)

func InitEnv(env string) (*viper.Viper, error) {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.SetConfigName(env)
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("config file not found")
	}
	return viper, nil
}

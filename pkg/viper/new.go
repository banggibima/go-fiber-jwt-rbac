package viper

import (
	"github.com/spf13/viper"
)

func New() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigType("env")
	v.SetConfigName(".env")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

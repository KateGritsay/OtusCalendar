package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(path string, config interface{}) error {
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
		err = viper.Unmarshal(&config)

		return err
}
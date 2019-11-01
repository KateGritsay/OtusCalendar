package config

import (
	"github.com/spf13/viper"
)

func Config(path string, config interface{}) error {
	viper.AddConfigPath(path)      // path to look for the config file in

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
		err = viper.Unmarshal(&config)

		return err
}
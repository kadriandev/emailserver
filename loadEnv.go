package main

import (
	"time"

	"github.com/spf13/viper"
)

type ConfigVars struct {
	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`
	Email        string        `mapstructure:"EMAIL"`
	AppPassword  string        `mapstructure:"APP_PASSWORD"`
}

func LoadConfig(path string) (config ConfigVars, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Profile string
	Application   AppConfiguration
}

type AppConfiguration struct {
	Port string
	Database MongoConfiguration
}

type MongoConfiguration struct {
	DB_ROOT   string
	DB_NAME   string
	MONGO_URI string
}

func LoadConfig() (Configuration, error) {
	viper.Set("PROFILE", "dev")
	conf := Configuration{}
	profile := viper.GetString("PROFILE")

	var APP_NAME string
	if profile != "default" {
		APP_NAME = "application-" + profile
	} else {
		APP_NAME = "application"
	}
	viper.SetConfigName(APP_NAME)
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return Configuration{}, err
	}

	return conf, err
}

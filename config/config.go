package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Port int `mapstructure:"PORT"`
	}
)

func LoadConfig() (*Config, error) {
	var config Config

	// Set Viper to use environment variables
	viper.AutomaticEnv()
	viper.BindEnv("PORT")

	// Unmarshal the environment variables into the config struct
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Print the loaded config for verification
	fmt.Printf("Loaded Config: %+v\n", config)

	return &config, nil
}

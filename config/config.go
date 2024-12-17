package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Port                   int    `mapstructure:"PORT"`
		MinioEndpoint          string `mapstructure:"MINIO_ENDPOINT"`
		MinioAccessKey         string `mapstructure:"MINIO_ACCESS_KEY"`
		MinioSecretKey         string `mapstructure:"MINIO_SECRET_KEY"`
		MinioUseSSL            string `mapstructure:"MINIO_USE_SSL"`
		MinioBucketName        string `mapstructure:"MINIO_BUCKET_NAME"`
		MinioPresignedDuration int    `mapstructure:"MINIO_PRESIGNED_DURATION"`
	}
)

func LoadConfig() (*Config, error) {
	var config Config

	// Set Viper to use environment variables
	viper.AutomaticEnv()
	viper.BindEnv("PORT")
	viper.BindEnv("MINIO_ENDPOINT")
	viper.BindEnv("MINIO_ACCESS_KEY")
	viper.BindEnv("MINIO_SECRET_KEY")
	viper.BindEnv("MINIO_USE_SSL")
	viper.BindEnv("MINIO_BUCKET_NAME")
	viper.BindEnv("MINIO_PRESIGNED_DURATION")

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

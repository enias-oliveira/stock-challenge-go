package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost      string `mapstructure:"DB_HOST"`
	DBName      string `mapstructure:"DB_NAME"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	APIHost     string `mapstructure:"API_HOST"`
	APIPort     string `mapstructure:"API_PORT"`
	StooqAPIKey string `mapstructure:"STOOQ_API_KEY"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD",
	"API_HOST", "API_PORT", "STOOQ_API_KEY", "JWT_SECRET",
}

// TODO: Read from enviroment variables. (Currently only reading from .env file)
func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}

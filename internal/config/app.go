package config

import "github.com/spf13/viper"

type AppConfig struct {
	Port     string `mapstructure:"port"`
	DB       DB
	Messages Messages
}

type DB struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	AddedSuccessfully   string `mapstructure:"added_successfully"`
	RemovedSuccessfully string `mapstructure:"removed_successfully"`
}

type Errors struct {
	Default       string `mapstructure:"default"`
	InvalidEmail  string `mapstructure:"invalid_email"`
	InvalidAge    string `mapstructure:"invalid_age"`
	AlreadyExists string `mapstructure:"already_exists"`
}

func InitApp(configPath string) (*AppConfig, error) {
	viper.SetConfigName("main")
	viper.AddConfigPath(".")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

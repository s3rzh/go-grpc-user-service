package config

import "github.com/spf13/viper"

type AppConfig struct {
	Port       string `mapstructure:"port"`
	DB         DB
	Cache      Cache
	Queue      Queue
	ClickHouse ClickHouse
	Messages   Messages
}

type DB struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Cache struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Queue struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	QueueName string `mapstructure:"queuename"`
}

type ClickHouse struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
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
	Default          string `mapstructure:"default"`
	InvalidEmail     string `mapstructure:"invalid_email"`
	InvalidInputData string `mapstructure:"invalid_input_data"`
	AlreadyExists    string `mapstructure:"already_exists"`
	NotExists        string `mapstructure:"not_exists"`
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

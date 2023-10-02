package config

import (
    "database/sql"
    "fmt"
    "log"
    "github.com/spf13/viper"
    _ "github.com/lib/pq"
)

type Config struct {
	DB		PSQL 		`mapstructure:"database"`
	Server 	Server 		`mapstructure:"server"`
}

type PSQL struct {
	Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
}

type Server struct {
	UserAddress     string `mapstructure:"user_address"`
	UserAddressHttp string `mapstructure:"user_address_http"`
}

func InitConfig() (*Config, error) {
	configFilePath := "config.local.yaml"
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func InitDB(config *Config) (*sql.DB, error) {
	DBConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.DBName)
	
	db, err := sql.Open("postgres", DBConnectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Successfully connected to database!")
	
	return db, nil
}
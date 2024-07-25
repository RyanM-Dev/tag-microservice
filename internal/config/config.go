package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
}

var MysqlConfig Config

func LoadMysqlConfig() {
	viper.SetConfigName("mysql")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&MysqlConfig); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}
}

func ConfigToDsn(config Config) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	return dsn
}
